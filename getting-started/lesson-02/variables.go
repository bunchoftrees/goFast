package main

import "fmt"

func main() {	
	// Variable declaration
	var message string
	message = "Hello, World!"
	fmt.Println(message)

	// Type inference
	var name = "John Bear"
	fmt.Println(name)

	// Explicit type declaration
	var age int = 38
	fmt.Println(age)

	// Multiple variable declaration
	var city, state, year = "Chicago", "Illinois", 2025
	fmt.Println(city, state, year)

	// Multi-line variable declaration with type inference
	var (
		planet = "Earth"
		distance float32 = 92.96
		unit = "million miles"
		life = true
	)
	fmt.Println(planet, distance, unit, life)

	// Short variable declaration
	animal := "Bear"
	fmt.Println(animal)

	// Short variable declaration with multiple variables
	food, drink, cute := "Honey", "Water", true
	fmt.Println(food, drink, cute)
}