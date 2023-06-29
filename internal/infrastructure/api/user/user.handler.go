package api

import (
	"api-book/internal/domain/dtos"
	"api-book/internal/domain/repository"
	"encoding/json"
	"fmt"

	"net/http"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.UserRepo.FindAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener los usuarios: %v", err), http.StatusInternalServerError)
		return
	}

	data := dtos.ResponseDTO{
		Status:  "success",
		Message: "data found",
		Data:    users,
	}

	jData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error procesar los datos usuarios: %v", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
