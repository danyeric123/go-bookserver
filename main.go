package main

import (
	"log"
	"os"
	"net/http"
	"github.com/danyeric123/go-bookserver/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Println("Server running on 8080")
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe("localhost:8080", loggedRouter))
}
