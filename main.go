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

	// ARRAY
	// var myFamily[3]string
	// myFamily[0]="Jay"
	// myFamily[1]="Jatin"
	// myFamily[2]="Justin"

	myFamily := [3]string{"Jay", "Jane", "Jatin"}
	myFamily[1] = "kate"

	fmt.Println("My family : %v", myFamily)

	// Multi-dimensional Array
	myCourses := [3][2]string{
		{"Go", "NodeJS"},  // 1st array
		{"AWS", "GCP"},    // 2nd array
		{"CDK", "Pulumi"}, // 3rd array
	}

	fmt.Println("My courses : %v", myCourses)
	// Result : My friends : %v [Mike Jatin Adam Sam Jay]

	// Slice
	var myFriends []string
	myFriends = append(myFriends, "Mike", "Jatin", "Adam")
	fmt.Println("My friends : %v", myFriends)
	// Result : My friends : %v [Mike Jatin Adam]
	
	myFriends = append(myFriends, "Sam", "Jay")
	// Result : My friends : %v [Mike Jatin Adam Sam Jay]
	fmt.Println("My friends : %v", myFriends)

	mySliceCourses := [][]string{
		{"Go", "NodeJS"},  // 1st array
		{"AWS", "GCP"},    // 2nd array
		{"CDK", "Pulumi"}, // 3rd array
	}
	course := []string{"IAC", "Cloud Formation"}
	mySliceCourses = append(mySliceCourses, course)
	mySliceCourses = append(mySliceCourses, []string{"react", "react-native"})

	fmt.Println("My slice courses : %v", mySliceCourses)
	// Result : My slice courses : %v [[Go NodeJS] [AWS GCP] [CDK Pulumi] [IAC Cloud Formation] [react react-native]]

	// make() use for create slice
	myBeCourses := make([]int, 2, 10)

	fmt.Println("My be courses : %v", myBeCourses)
	// Result : My be courses : %v [0 0]
	app.Listen("localhost:9000")
}
