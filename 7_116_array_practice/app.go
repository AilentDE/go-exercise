package main

import "fmt"

type hobby struct {
	name      string
	frequency int
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]hobby{
		{
			"coding",
			6,
		},
		{
			"drawing",
			2,
		},
		{
			"music",
			7,
		},
	}
	fmt.Println("All hobbies:", hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("First hobby:", hobbies[0])
	fmt.Println("Other hobbies:", hobbies[1:])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	betterHobbiesOptions := hobbies[:2]
	fmt.Println("Better hobbies:", betterHobbiesOptions)
	betterHobbiesOptions = hobbies[0:2]
	fmt.Println("[Re]Better hobbies:", betterHobbiesOptions)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	betterHobbiesOptions = betterHobbiesOptions[1:3]
	fmt.Println("Change preferred hobbies:", betterHobbiesOptions)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Use Golang.", "Get faster backend."}
	fmt.Println("My goals:", goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Do better on backend."
	goals = append(goals, "Get better job.")
	fmt.Println("New goals:", goals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		title string
		id string
		price int
	}
	products := []Product{
		{"apple", "p1", 30}, {"pen", "p2", 100},
	}
	products = append(products, Product{"applepen", "p3", 3000})
	fmt.Println("Products:", products)
}
