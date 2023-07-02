package api

import (
	"api-book/internal/domain/models"
	"api-book/internal/domain/repository"
	"api-book/internal/infrastructure/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"net/http"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

/** LISTA LOS USUARIOS CON EL HISTORIAL DE LOS PRESTAMOS */
func (uh *UserHandler) GetHistoryLendUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			utils.BadResponse("Method not allowed"),
			http.StatusMethodNotAllowed)
		return
	}
	users, err := uh.UserRepo.GetUsersHistory()
	if err != nil {
		http.Error(w, utils.BadResponse(err.Error()), http.StatusUnprocessableEntity)
		return
	}

	jData, _ := utils.CreateResponseApi(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)

}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error al obtener el id")),
			http.StatusBadRequest)
		return
	}

	_, err = uh.UserRepo.GetUserId(idNum)
	if err != nil {
		http.Error(w, utils.BadResponse("no existe el usuario"), http.StatusNotFound)
		return
	}

	current_time := time.Now()
	var updateUser models.UpdateUser

	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, utils.BadResponse("Error al decodificar los datos del usuario"), http.StatusBadRequest)
		return
	}

	updateUser.UpdatedAt = current_time
	// updateUser.Password = utils.HashPassword(updateUser.Password, utils.GenerateSalt(11))

	err = uh.UserRepo.UpdateUser(&updateUser, idNum)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error actualizar el usuario: %v", err)),
			http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(updateUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)

}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	current_time := time.Now()
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, utils.BadResponse("Error al decodificar los datos del usuario"), http.StatusBadRequest)
		return
	}

	newUser.Password = utils.HashPassword(newUser.Password, utils.GenerateSalt(11))
	newUser.CreatedAt = current_time

	err = uh.UserRepo.CreateUser(&newUser)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error al crear el usuario: %v", err)),
			http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)

}

func (uh *UserHandler) GetUserId(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error al obtener el id: %v", err)),
			http.StatusBadRequest)
		return
	}
	user, err := uh.UserRepo.GetUserId(idNum)
	if err != nil {
		http.Error(w, utils.BadResponse(err.Error()), http.StatusNoContent)
		return
	}
	jData, err := utils.CreateResponseApi(user)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error procesar los datos del usuario: %v", err)),
			http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	users, err := uh.UserRepo.FindAllUsers()
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error al obtener los usuarios: %v", err)),
			http.StatusInternalServerError)
		return
	}

	jData, err := utils.CreateResponseApi(users)

	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error procesar los datos usuarios: %v", err)),
			http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
