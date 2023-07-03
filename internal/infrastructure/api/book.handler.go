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

/** LISTA EL HISTORIAL DE PRESTAMO DE LOS LIBROS */
func (h *BookHandler) GetHistoryLendBookV2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			utils.BadResponse("Method not allowed"),
			http.StatusMethodNotAllowed,
		)
		return
	}
	books, err := h.BookRepo.GetBooksHistoryV2()
	if err != nil {
		http.Error(
			w,
			utils.BadResponse(err.Error()),
			http.StatusInternalServerError,
		)
		return
	}
	jData, _ := utils.CreateResponseApi(books)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

/** LISTA EL HISTORIAL DE PRESTAMO DE LOS LIBROS */
func (h *BookHandler) GetHistoryLendBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			utils.BadResponse("Method not allowed"),
			http.StatusMethodNotAllowed,
		)
		return
	}
	books, err := h.BookRepo.GetBooksHistory()
	if err != nil {
		http.Error(
			w,
			utils.BadResponse(err.Error()),
			http.StatusInternalServerError,
		)
		return
	}
	jData, _ := utils.CreateResponseApi(books)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	// utils.LogInfoData("newBook.log", "CreateBook", book)
	if err != nil {
		http.Error(w, utils.BadResponse("Error procesar los datos"), http.StatusBadRequest)
		return
	}

	err = h.BookRepo.CreateBook(&book)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("error al crear el libros: %v", err)),
			http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(book)
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, utils.BadResponse("error al obtener el id"), http.StatusBadRequest)
		return
	}

	bookIdNum, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("Error al obtener el id")),
			http.StatusBadRequest)
		return
	}

	book, err := h.BookRepo.GetBookByID(bookIdNum)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("No se pudo encontrat el libro: %v", err)),
			http.StatusInternalServerError)
		return
	}

	jData, _ := utils.CreateResponseApi(book)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, utils.BadResponse("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	page, limit := utils.GetPaginatorSql(r.Header)

	books, err := h.BookRepo.ListBooks(page, limit)
	if err != nil {
		http.Error(w,
			utils.BadResponse(fmt.Sprintf("error al listar libros : %v", err)),
			http.StatusInternalServerError)
		return
	}
	jData, _ := utils.CreateResponseApi(books)
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
