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

	// Basic Types: int,float64,string,bool
	// Composite Types : array,slice,map,struct
	// Pointer types: *

	// Map
	// Key value type
	// store unique values , cannot keep the duplicate key
	myWishList := make(map[string]string)
	myWishList["first"] = "MacPro"
	myWishList["second"] = "900 Billion Dollar"
	myWishList["third"] = "a beautiful car"

	fmt.Println("My wish list %v", myWishList)
	delete(myWishList, "third")
	fmt.Println("My wish list %v", myWishList)


	type Details struct {
		Description string 
		Images 		string
	}

	// struct : group of variables
	type Product struct {
		Name  string `json:"product_name"`
		Price float64 `json:"price"`
		Details  Details `json:"details"`
	}

	// var product Product
	product := Product{
		Name:  "MacPro", 
		Price: 9000, 
		Details:Details{
			Description: "An incredible machine",
			Images: "http:/macproimage.jpg",
		},
	}

	product.Name = "Macbook Pro"

	fmt.Println("My wish list %v", product)
	
	app.Get("/product", func(c *fiber.Ctx) error {
		return c.JSON(product)
	})

	app.Listen("localhost:9000")
}
