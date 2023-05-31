package main

import "fmt"

func main() {

	//fmt.Println("Hello, ninjas!")

	var nameOne string = "hello"
	fmt.Println(nameOne)

	var nameTwo = "World" //go infers the type of this

	var nameThree string

	nameFour := "viasat"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//ints
	var ageOne int = 10
	var ageTwo = 20
	ageThree := 30

	//bit memory

	var numOne int8 = 25
	var numTwo int8 = -125
	var numThree uint = 25

	//float
	var score1 float32 = 25.88
	var score2 float64 = 35.888888888
	// for more info see https://golang.org/ref/spec#Numeric_types

}
