package main

import (
	"fmt"
	_ "go-ecommerce-app/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Conditional Statements
	// Pointers

	/*
		 Conditional Statements
		- if else
		- switch case
		- select
		- loop
	*/

	age := 29
	if age > 65 {
		fmt.Println("Senior Citizen")
	} else if age > 17 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}

	seatClass := "firstClass"

	switch seatClass {
	case "firstClass":
		{
			fmt.Println("You will get free drinks")
		}
	case "businessClass":
		{
			fmt.Println("You will get more legrooms")
		}
	default:
		{
			fmt.Println("You need to pay for services")
		}
	}

	var myFriends[]string
	for i:=0;i<2;i++{
		myNewFriend := fmt.Sprintf("Friend %d",i)
		myFriends = append(myFriends, myNewFriend)
	}

	for index,value := range myFriends{
		fmt.Println(index,value)
	}

	// isOver :=0
	// for{
	// 	isOver++;
	// 	fmt.Println(isOver)
	// 	if isOver > 99{
	// 		fmt.Println("It's really over now")
	// 		return 
	// 	}
	// }

	// Infinite loop
	// for {
	// 	fmt.Println("I am listening till server alive!")
	// } 


	jay:= "laptop" // 0xc0001961e0 memory address
	fmt.Println(jay) 
	
	var guest *string
	fmt.Println(guest) // nil
	guest = &jay 
	fmt.Println(guest) // 0xc0001961e0 memory address
	fmt.Println(*guest) // print laptop


	app.Listen("localhost:9000")
}
