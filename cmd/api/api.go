package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/EaleksS/green-shop-api/service/cart"
	"github.com/EaleksS/green-shop-api/service/category"
	"github.com/EaleksS/green-shop-api/service/favorite"
	"github.com/EaleksS/green-shop-api/service/product"
	"github.com/EaleksS/green-shop-api/service/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer  {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8000"
		},
		AllowCredentials: true,
		Debug: true,
	 })

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	categoryStore := category.NewStore(s.db)
	categoryHandler := category.NewHandler(categoryStore, productStore)
	categoryHandler.RegisterRoutes(subrouter)

	favoriteStore := favorite.NewStore(s.db)
	favoriteHandler := favorite.NewHandler(favoriteStore, userStore, productStore)
	favoriteHandler.RegisterRoutes(subrouter)

	cartStore := cart.NewStore(s.db)
	cartHandler := cart.NewHandler(cartStore, userStore, productStore)
	cartHandler.RegisterRoutes(subrouter)

	handler := corsMiddleware.Handler(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, handler)
}