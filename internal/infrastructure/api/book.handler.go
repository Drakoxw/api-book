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

type BookHandler struct {
	BookRepo *repository.BookRepository
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Error procesar los datos", http.StatusBadRequest)
		return
	}

	err = h.BookRepo.CreateBook(&book)
	if err != nil {
		http.Error(w, fmt.Sprintf("error al crear el libros: %v", err), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(book)
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "error al obtener el id", http.StatusBadRequest)
		return
	}

	bookIdNum, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el id"), http.StatusBadRequest)
		return
	}

	book, err := h.BookRepo.GetBookByID(bookIdNum)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se pudo encontrat el libro: %v", err), http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(book)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	page, limit := utils.GetPaginatorSql(r.Header)

	books, err := h.BookRepo.ListBooks(page, limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("error al listar libros : %v", err), http.StatusInternalServerError)
		return
	}
	jData, _ := utils.CreateResponseApi(books)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
