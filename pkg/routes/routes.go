package routes

import (
	"github.com/danyeric123/go-bookserver/pkg/controllers"
	"github.com/gorilla/mux"
)

 var RegisterBookRoutes = func(router *mux.Router){
	 router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	 router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	 router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	 router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	 router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
 }