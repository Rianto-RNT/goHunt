package main

import "fmt"

func main() {
	//int age = 10 //c++ way
	var age int = 30
	fmt.Println("age:", age)

	var name = "Rian"
	// fmt.Println("my name is:", name)
	_ = name

	s := "Learning Golang"
	fmt.Println(s)

	name = "Al Fatih"
	name1 := "Khairi"
	_ = name1
}
