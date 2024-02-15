package routes

import "github.com/gorilla/mux"

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", Controller)
}
