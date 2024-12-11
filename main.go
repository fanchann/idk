package main

import (
	"net/http"

	"fanchann/elegant_err/handler"
	"fanchann/elegant_err/services"
)

func main() {
	usersServices := services.NewUsersServices()
	usersHandler := handler.NewUsersHandlerImpl(usersServices)

	http.HandleFunc("GET /search", usersHandler.SearchUserById)

	http.ListenAndServe(":8080", nil)
}
