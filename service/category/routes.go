package category

import (
	"fmt"
	"net/http"

	"github.com/EaleksS/green-shop-api/types"
	"github.com/EaleksS/green-shop-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CategoryStore
	ProductStore types.ProductStore
}

func NewHandler(store types.CategoryStore, ProductStore types.ProductStore) *Handler {
	return &Handler{store: store, ProductStore: ProductStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/category", h.handleGetCategory).Methods(http.MethodGet)
	router.HandleFunc("/category", h.handleCreateCategory).Methods(http.MethodPost)
}

func (h *Handler) handleGetCategory(w http.ResponseWriter, r *http.Request) {
	cs, err := h.store.GetCategory()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	ps, errPs := h.ProductStore.GetProducts()
	if errPs != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	maxPrice := findMaxPrice(ps)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": cs,
		"maxPrice": maxPrice,
	})
}

func (h *Handler) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateCategoryPayload

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

	err := h.store.CreateCategory(types.Category{
		Name: payload.Name,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"description": "The creation of a new category has been successfully completed"})
}