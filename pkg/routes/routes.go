package routes

import (
	"github.com/danyeric123/go-bookserver/pkg/controllers"
	"github.com/gorilla/mux"
)

 var RegisterBookRoutes = func(router *mux.Router){
	 router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	 router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	 router.HandleFunc("/books/{bookId}", controllers.GetBook).Methods("GET")
	 router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	 router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
 }