package cart

import (
	"fmt"
	"net/http"

	"github.com/EaleksS/green-shop-api/service/auth"
	"github.com/EaleksS/green-shop-api/types"
	"github.com/EaleksS/green-shop-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CartStore
	userStore types.UserStore
	productStore types.ProductStore
}

func NewHandler(store types.CartStore, userStore types.UserStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, userStore: userStore, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart", auth.WithJWTAuth(h.handleGetCart, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/cart", auth.WithJWTAuth(h.handleDeleteCart, h.userStore)).Methods(http.MethodDelete)
	router.HandleFunc("/cart", auth.WithJWTAuth(h.handleAddCart, h.userStore)).Methods(http.MethodPut)
}

func (h *Handler) handleGetCart(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	fs, ids, err := h.store.GetCart(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	ps, errPs := h.productStore.GetProductsByIDs(ids)
	if errPs != nil {
		utils.WriteError(w, http.StatusInternalServerError, errPs)
		return
	}

	totalPrice := findTotalPrice(ps)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": fs,
		"total": len(fs),
		"ids": ids,
		"products": ps,
		"totalPrice": totalPrice,
	})
}

func (h *Handler) handleDeleteCart(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateCartPayload

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

	userID := auth.GetUserIDFromContext(r.Context())

	err := h.store.DeleteCart(userID, payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Успешно удален из корзины")
}

func (h *Handler) handleAddCart(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateCartPayload

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

	userID := auth.GetUserIDFromContext(r.Context())

	err := h.store.AddCart(userID, payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Добавлен в корзину")
}
