package restbooksserver

import (
	"books/internal/pkg/model"
	"books/internal/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const SuccessResponseFieldKey = "status"
const ErrorResponseFieldKey = "error"

type BooksHandler struct {
	bookService service.BooksService
}

func GetNewBooksHandler(booksService service.BooksService) *BooksHandler {
	return &BooksHandler{bookService: booksService}
}

func (bh *BooksHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetAllBooks()
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

func (bh *BooksHandler) GetOrRemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	muxVar := mux.Vars(r)
	isbnStr := muxVar["isbn"]
	isbn, err := strconv.Atoi(isbnStr)
	if err != nil || isbn == 0 {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if r.Method == http.MethodGet {
		book, err := bh.bookService.GetBook(isbn)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, book)
	} else if r.Method == http.MethodDelete {
		bh.bookService.RemoveBook(isbn)
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	}
}

func (bh *BooksHandler) UpsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if r.Method == http.MethodPost {
		bh.bookService.AddBook(DBBook(book))
	} else if r.Method == http.MethodPut {
		bh.bookService.UpdateBook(DBBook(book))
	} else {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "ok"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
