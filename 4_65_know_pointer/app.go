package main

import "fmt"

func main() {
	age := 32
	agePointer := &age // address in memory
	
	fmt.Println("Age address:", agePointer)
	fmt.Println("Age value:", *agePointer)

	generalFunc(age)// copy of age
	editAdultYears(agePointer)// original age

	fmt.Println(age)
}

func generalFunc(age int) {
	point := &age
	fmt.Println("Age in general function:", point)
	age = age - 20

}

func editAdultYears(age *int) {
	fmt.Println("Age in pointer function:", age)
	*age = *age - 18
}