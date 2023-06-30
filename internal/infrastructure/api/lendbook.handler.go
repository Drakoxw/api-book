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

func (h *LendBookHandler) CreateLendBook(w http.ResponseWriter, r *http.Request) {

	var lendBook models.LendBook
	err := json.NewDecoder(r.Body).Decode(&lendBook)
	if err != nil {
		http.Error(w, "No se pudo procesar los datos", http.StatusUnprocessableEntity)
		return
	}

	err = h.LendBookRepo.CreateLendBook(&lendBook)
	if err != nil {
		http.Error(w, fmt.Sprintf(err.Error()), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(lendBook)
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)
}

func (h *LendBookHandler) GetLendBook(w http.ResponseWriter, r *http.Request) {

	lendBookID := r.URL.Query().Get("id")
	if lendBookID == "" {
		http.Error(w, "se requiere un id", http.StatusBadRequest)
		return
	}

	lendBookIDNum, err := strconv.Atoi(lendBookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el id"), http.StatusBadRequest)
		return
	}

	lendBook, err := h.LendBookRepo.GetLendBookByID(lendBookIDNum)
	if err != nil {
		http.Error(w, fmt.Sprintf(err.Error()), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(lendBook)
}

func (h *LendBookHandler) ListLendBooks(w http.ResponseWriter, r *http.Request) {

	lendBooks, err := h.LendBookRepo.GetAllBooksAndLends()
	if err != nil {
		http.Error(w, fmt.Sprintf(err.Error()), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(lendBooks)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (h *LendBookHandler) ReturnBookToLibrary(w http.ResponseWriter, r *http.Request) {

	bookID := r.URL.Query().Get("lend_id")
	if bookID == "" {
		http.Error(w, "Se requiere el id", http.StatusBadRequest)
		return
	}

	bookIDNum, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el id"), http.StatusBadRequest)
		return
	}

	_, err = h.LendBookRepo.GetLendBookByID(bookIDNum)
	if err != nil {
		http.Error(w, fmt.Sprintln("No existe el Un prestamo con el id :", bookIDNum), http.StatusInternalServerError)
		return
	}

	err = h.LendBookRepo.ReturnBookToLibrary(bookIDNum)
	if err != nil {
		http.Error(w, fmt.Sprintf(err.Error()), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(bookIDNum)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (h *LendBookHandler) GetAllUsersAndLends(w http.ResponseWriter, r *http.Request) {

	usersLends, err := h.LendBookRepo.GetAllUsersAndLends()
	if err != nil {
		http.Error(w, fmt.Sprintf(err.Error()), http.StatusBadRequest)
	}

	jData, _ := utils.CreateResponseApi(usersLends)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
