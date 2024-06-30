package product

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/EaleksS/green-shop-api/types"
	"github.com/EaleksS/green-shop-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", h.handleGetByIDProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/game/{name}", h.handleGetByNameProduct).Methods(http.MethodGet)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sortBy := query.Get("sort_by")
	search := query.Get("search")
	category := query.Get("category")
	lowPrice, _ := strconv.ParseFloat(query.Get("low_price"), 64)
	highPrice, _ := strconv.ParseFloat(query.Get("high_price"), 64)

	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	products := priceSort(reverseArray(ps), highPrice, lowPrice)
	productsSorting := sortBySort(products, sortBy)
	categorySorting := categorySort(productsSorting, category)
	searchSorting := searchSort(categorySorting, search)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": searchSorting,
		"total": len(searchSorting),
	})
}



func (h *Handler) handleGetByIDProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  id := vars["id"]

	ps, err := h.store.GetByIDProducts(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleGetByNameProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  name := vars["name"]
	decodedName, err := url.QueryUnescape(name)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ps, err := h.store.GetByNameProducts(decodedName)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateProducts(types.Product{
		Name: payload.Name,
		Category: payload.Category,
		Description: payload.Description,
		Image: payload.Image,
		Price: payload.Price,
		Quantity: payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"description": "The creation of a new product has been successfully completed"})
}


