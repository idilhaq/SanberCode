package main

import "fmt"

func yourHobbies(name string, hobbies ...string) {
	fmt.Println("Hello, my name is", name)
	fmt.Println("My hobbies are: ")
	for _, hobby := range hobbies {
		fmt.Println(hobby)
	}
}

func main() {
	yourHobbies("John", "Gaming", "Hiking", "Reading")
	fmt.Println()

	var hobbies = []string{"Sleeping", "Eating"}
	yourHobbies("Doe", hobbies...)
}
