package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/mylibrary"
	"time"
	"errors"
	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	bookList *mylibrary.List
}

func NewHTTPHandlers(bookList *mylibrary.List) *HTTPHandler {
	return &HTTPHandler{
		bookList: bookList,
	}
}

func (h *HTTPHandler) HandleCreateBook(w http.ResponseWriter, r *http.Request){
	var bookDTO BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil{
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	libraryBook := mylibrary.AddBook(bookDTO.Title, bookDTO.Author, bookDTO.Pages, bookDTO.Text)
	if err := h.bookList.AddBook(libraryBook); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}


		if errors.Is(err, mylibrary.ErrBookAlreadyInLibrary) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(libraryBook, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}

func (h *HTTPHandler) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	book, err := h.bookList.GetBook(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}

		if errors.Is(err, mylibrary.ErrBookNotFound){
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		}else{
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(book, "", "    ")
	if err != nil{
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil{
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandler)HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.bookList.ListBooks()
	b, err := json.MarshalIndent(books, "", "    ")
	if err != nil{
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil{
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandler)HandleGetAllUnreadedBooks(w http.ResponseWriter, r *http.Request){
	unreadedBooks := h.bookList.ListUnreadedBooks()
	b, err := json.MarshalIndent(unreadedBooks, "", "    ")
	if err != nil{
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil{
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandler)HandleGetAllReadedBooks(w http.ResponseWriter, r *http.Request){
	unreadedBooks := h.bookList.ListReadedBooks()
	b, err := json.MarshalIndent(unreadedBooks, "", "    ")
	if err != nil{
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil{
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandler)HandleReadBook(w http.ResponseWriter, r *http.Request){
	var readDTO ReadBookDTO
	if err := json.NewDecoder(r.Body).Decode(&readDTO); err != nil{
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	title := mux.Vars(r)["title"]

	var changedBook mylibrary.Book
	var err error

	if readDTO.Read{
		changedBook, err = h.bookList.ReadBook(title)
	}

	if err != nil{
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}

		if errors.Is(err, mylibrary.ErrBookNotFound){
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		}else{
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(changedBook, "", "    ")
	if err != nil{
		panic(err)
	}

	if _, err := w.Write(b); err != nil{
		fmt.Println("failed to write http response:", err)
		return
	}
}

func (h *HTTPHandler)DeleteBook(w http.ResponseWriter, r *http.Request){
	title := mux.Vars(r)["title"]

	if err := h.bookList.DeleteBook(title); err != nil{
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}

		if errors.Is(err, mylibrary.ErrBookNotFound){
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		}else{
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}