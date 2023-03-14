package main

import "fmt"

//Go Structures
//A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
//struct (strukturaning qisqartmasi) bitta o'zgaruvchiga har xil turdagi ma'lumotlarning a'zolari to'plamini yaratish uchun ishlatiladi.

//arrays vs struct
//*** arrays - store multiple values of the same data type***
//*** structs - store multiple values of different data types ***

//To declare a structure in Go, use the type and struct keywords:
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

//To access any member of a structure, use the dot operator (.) between the structure variable name and the structure member:

//You can also pass a structure as a function argument, like this:
func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("***Person 1***")
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("\n***Person 2***")
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	//as func parametr
	// Print Pers1 info by calling a function
	fmt.Println("\n***Person 1 from func***")
	printPerson(pers1)

	// Print Pers2 info by calling a function
	fmt.Println("\n***Person 2 from func***")
	printPerson(pers2)
}
