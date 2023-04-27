package server

import (
	"fmt"
	"net/http"

	"assignment4/internal/handlers"

	"github.com/gorilla/mux"
)

type server struct {
	port int
}

func NewServer(port int) server {
	return server{port: port}
}
func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/contact", handlers.GetContacts).Methods("GET")
	router.HandleFunc("/contact", handlers.SetContact).Methods("POST")
	router.HandleFunc("/contact", handlers.UpdateContact).Methods("PUT")
	router.HandleFunc("/contact", handlers.DeleteContact).Methods("DELETE")

	return router
}

func (s *server) StartServer() error {

	var err error
	router := InitRouter()
	port := fmt.Sprintf(":%d", s.port)

	err = http.ListenAndServe(port, router)
	if err != nil {
		fmt.Println("Error increating Server")
		return err
	}
	fmt.Println("Server Started...")
	return err
}
