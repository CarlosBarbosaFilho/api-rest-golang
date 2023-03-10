package routes

import (
	"log"
	"net/http"

	"github.com/CarlosBarbosaFilho/api-rest-golan/controllers"
	"github.com/CarlosBarbosaFilho/api-rest-golan/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", middleware.ContentTypeMiddleware(controllers.Home))
	r.HandleFunc("/api/personas", middleware.ContentTypeMiddleware(controllers.AllPersonas)).Methods("GET")
	r.HandleFunc("/api/personas/{id}", middleware.ContentTypeMiddleware(controllers.GetById)).Methods("GET")
	r.HandleFunc("/api/personas", middleware.ContentTypeMiddleware(controllers.CreatePersona)).Methods("POST")
	r.HandleFunc("/api/personas/{id}", middleware.ContentTypeMiddleware(controllers.RemovePersona)).Methods("DELETE")
	r.HandleFunc("/api/personas/{id}", middleware.ContentTypeMiddleware(controllers.EditPersona)).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
