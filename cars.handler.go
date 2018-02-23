package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CarsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;chartset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetCars()); err != nil {
		panic(err)
	}
}

func CarsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["id"]
	car, err := GetCar(carId)

	if err == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json;chartset=UTF-8")
		if err := json.NewEncoder(w).Encode(car); err != nil {
			panic(err)
		}

	}
}
