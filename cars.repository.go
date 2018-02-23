package main

import (
	"errors"

	"github.com/google/uuid"
)

var cars []Car

func init() {

	id, err := uuid.Parse("b9ec3ec6-18e2-11e8-9619-1867b03edad8")

	if err != nil {
		panic(err)
	}

	CreateCar(Car{Name: "Toyota Carina", ID: id})

}

func CreateCar(car Car) {
	cars = append(cars, car)
}

func GetCars() []Car {
	return cars
}

func GetCar(ID string) (Car, error) {
	parsedID, _ := uuid.Parse(ID)
	for _, car := range cars {
		if car.ID == parsedID {
			return car, nil
		}
	}

	return Car{}, errors.New("not found")
}
