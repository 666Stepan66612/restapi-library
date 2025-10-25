package main

import (
	"fmt"
	"restapi/http"
	"restapi/mylibrary"
)

func main(){
	bookList := mylibrary.NewBook()
	httpHandlers := http.NewHTTPHandlers(bookList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil{
		fmt.Println("failed to start http server:", err)
	}
}