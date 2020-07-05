package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	financeRegistry "../models"
	registryService "../services"
)

func RegistryController() {
	log.Print("Into Controller...")
	router := mux.NewRouter()
	router.HandleFunc("/registries", PostRegistry).Methods("POST")
	router.HandleFunc("/create", CreateCategoryForRegistry).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))

}

func PostRegistry(w http.ResponseWriter, r *http.Request) {
	var registry financeRegistry.FinanceRegistry
	json.NewDecoder(r.Body).Decode(&registry)
	createdRegistry := registryService.CreateRegistry(registry)
	json.NewEncoder(w).Encode(createdRegistry)
}

func CreateCategoryForRegistry(w http.ResponseWriter, r *http.Request) {
	var registry financeRegistry.FinanceRegistry
	json.NewDecoder(r.Body).Decode(&registry)
	registry = registryService.CreateCategoryByBaseRegistries(registry)
	json.NewEncoder(w).Encode(registry)
}
