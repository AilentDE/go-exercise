package main

import "fmt"

type courseTypes map[int64]string
func (c courseTypes) output() {
	fmt.Println(c)
}

func main() {
	// array
	userNames := make([]string, 2, 5)// 長度為 2、容量為 5 的切片
	fmt.Println(userNames, len(userNames), cap(userNames))

	userNames[0] = "Jay"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manul")
	userNames = append(userNames, "Jack")
	userNames = append(userNames, "Aria")

	fmt.Println(userNames, len(userNames), cap(userNames))

	// map
	courses := make(courseTypes, 3)

	courses[0] = "offline"
	courses[1] = "online"
	courses[2] = "disappear"
	courses.output()

	//apply loop
	for i,val := range userNames {
		fmt.Printf("%d user: %s\n", i, val)
	}
	for key, val := range courses {
		fmt.Printf("key: %d - value: %s\n", key, val)
	}
}