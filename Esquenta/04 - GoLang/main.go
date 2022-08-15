package main

//import "fmt"
import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 200})
	cars = append(cars, Car{Name: "Mercedes", Price: 150})
	cars = append(cars, Car{Name: "Porsche", Price: 125})
}

func main() {
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.POST("/cars/del", deleteCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func deleteCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	removeCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close() // adia o close até que tudo seja executado

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}

func removeCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close() // adia o close até que tudo seja executado

	stmt, err := db.Prepare("DELETE FROM cars WHERE name=($1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name)
	if err != nil {
		return err
	}
	return nil
}
