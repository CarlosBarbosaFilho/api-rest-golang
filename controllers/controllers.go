package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosBarbosaFilho/api-rest-golan/database"
	"github.com/CarlosBarbosaFilho/api-rest-golan/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func Personas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona
	database.DB.Find(&personas)

	json.NewEncoder(w).Encode(personas)
}

func GetById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var persona models.Persona

	database.DB.Find(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Create(&persona)
	json.NewEncoder(w).Encode(persona)
}

func RemovePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona
	database.DB.Delete(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func EditPersona(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	var persona models.Persona

	database.DB.Find(&persona, id)
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Save(&persona)

	json.NewEncoder(w).Encode(persona)

}
