package api

import (
	"api-book/internal/domain/models"
	"api-book/internal/domain/repository"
	"api-book/internal/infrastructure/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type LendBookHandler struct {
	LendBookRepo *repository.LendBookRepository
}

/** SE CREA EL REGISTRO DEL PRESTAMO */
func (h *LendBookHandler) CreateLendBook(w http.ResponseWriter, r *http.Request) {

	var lendBook models.LendBook
	err := json.NewDecoder(r.Body).Decode(&lendBook)
	if err != nil {
		http.Error(w,
			utils.BadResponse("No se pudo procesar los datos"),
			http.StatusUnprocessableEntity)
		return
	}

	err = h.LendBookRepo.CreateLendBook(&lendBook)
	if err != nil {
		http.Error(w, utils.BadResponse(err.Error()), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(lendBook)
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)
}

/** SE REGISTRA LA FECHA DE DEVOLUCION DEL LIBRO */
func (h *LendBookHandler) ReturnBookToLibrary(w http.ResponseWriter, r *http.Request) {

	bookID := r.URL.Query().Get("lend_id")
	if bookID == "" {
		http.Error(w, utils.BadResponse("Se requiere el id"), http.StatusBadRequest)
		return
	}

	bookIDNum, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, utils.BadResponse("Error al obtener el id"), http.StatusBadRequest)
		return
	}

	_, err = h.LendBookRepo.GetLendBookByID(bookIDNum)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintln("No existe el Un prestamo con el id :", bookIDNum)),
			http.StatusInternalServerError)
		return
	}

	err = h.LendBookRepo.ReturnBookToLibrary(bookIDNum)
	if err != nil {
		http.Error(w, utils.BadResponse(err.Error()), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(bookIDNum)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
