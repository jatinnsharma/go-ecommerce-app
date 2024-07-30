package main

import (
	"fmt"
	_ "go-ecommerce-app/config" // underscore(_) used for not in use still package needed indirectly

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	// declare a variable and initialize
	app := fiber.New()

	// config.LoadAppSettings()
	HelperFunction()

	// Basic Types: int , float64 , string ,bool
	// Composite Types : array,slice,map,struct
	// Pointer types: *

	// var age int
	// var height float64
	// var firstName string
	// var isEmployed bool

	// age = 29
	// height = 144.4
	// firstName = "Jatin"
	// isEmployed = true
	// fmt.Println(age,height,firstName,isEmployed)

	// fmt.Printf("Age: %v\n",age)
	// fmt.Printf("Height: %v\n",height)
	// fmt.Printf("FirstName: %v\n",firstName)
	// fmt.Printf("Employed?:: %v\n",isEmployed)

	// correct format specifier
	// %d -> int , %f -> float , %s -> string , %t -> boolean

	// automatically assign data (Type of data)
	// At the time of declare go is going to be assigned
	age := 29
	height := 144.4
	firstName := "Jatin"
	isEmployed := true

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %f\n", height)
	fmt.Printf("FirstName: %s\n", firstName)
	fmt.Printf("Employed?:: %t\n", isEmployed)

	if age > 65 {
		fmt.Println("Senior citizen")
	} else if age > 16 {
		fmt.Println("Adult")
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("Hello World!")
	}

	app.Listen("localhost:9000")
}
