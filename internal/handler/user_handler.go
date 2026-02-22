package handler

import (
	"SpendSmartAPI/internal/domain"
	"SpendSmartAPI/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	usecase *usecase.UserUseCase
}

func NewUserHandler(usecase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.usecase.Create(r.Context(), &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := handler.usecase.FindAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := handler.usecase.FindById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
