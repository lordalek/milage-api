package main

import (
	"encoding/json"
	"net/http"
)

func CarsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;chartset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetCars()); err != nil {
		panic(err)
	}
}

func CarsGet(w http.ResponseWriter, r *http.Request) {

}
