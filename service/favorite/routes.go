package favorite

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
	store types.FavoriteStore
	userStore types.UserStore
	productStore types.ProductStore
}

func NewHandler(store types.FavoriteStore, userStore types.UserStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, userStore: userStore, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/favorite", auth.WithJWTAuth(h.handleGetFavorite, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/favorite", auth.WithJWTAuth(h.handleDeleteFavorite, h.userStore)).Methods(http.MethodDelete)
	router.HandleFunc("/favorite", auth.WithJWTAuth(h.handleAddFavorite, h.userStore)).Methods(http.MethodPut)
}

func (h *Handler) handleGetFavorite(w http.ResponseWriter, r *http.Request) {
	

	userID := auth.GetUserIDFromContext(r.Context())

	fmt.Println(userID)

	fs, ids, err := h.store.GetFavorite(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	ps, errPs := h.productStore.GetProductsByIDs(ids)
	if errPs != nil {
		utils.WriteError(w, http.StatusInternalServerError, errPs)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": fs,
		"total": len(fs),
		"ids": ids,
		"products": ps,
	})
}

func (h *Handler) handleDeleteFavorite(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateFavoritePayload

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

	err := h.store.DeleteFavorite(userID, payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Успешно удален из избранных")
}

func (h *Handler) handleAddFavorite(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateFavoritePayload

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

	err := h.store.AddFavorite(userID, payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Добавлен в избранные")
}
