package main

import (
	"errors"

	"github.com/nu7hatch/gouuid"
)

var cars []Car

func init() {
	id, _ := uuid.NewV4()

	CreateCar(Car{Name: "Toyota Carina", ID: *id})
}

func CreateCar(car Car) {
	cars = append(cars, car)
}

func GetCars() []Car {
	return cars
}

func GetCar(ID [16]byte) (Car, error) {
	for _, car := range cars {
		if car.ID == ID {
			return car, nil
		}
	}

	return Car{}, errors.New("not found")
}
