package http

import (
	"errors"
	"net/http"
	"github.com/gorilla/mux"
)

type HTTPServer struct{
	httpHandlers *HTTPHandler
}

func NewHTTPServer(HTTPHandler *HTTPHandler) *HTTPServer{
	return &HTTPServer{
		httpHandlers: HTTPHandler,
	}
}

func (s *HTTPServer)StartServer() error{
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateBook)
	router.Path("/books/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetBook)
	router.Path("/books").Methods("GET").Queries("readed", "false").HandlerFunc(s.httpHandlers.HandleGetAllUnreadedBooks)
	router.Path("/books").Methods("GET").Queries("readed", "true").HandlerFunc(s.httpHandlers.HandleGetAllReadedBooks)
	router.Path("/books").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllBooks)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleReadBook)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.DeleteBook)

	if err := http.ListenAndServe(":9091", router); err != nil{
		if errors.Is(err, http.ErrServerClosed){
			return nil
		}

		return err
	}

	return nil
}