package main

import "fmt"

func main() {
	fmt.Println("Hello Pranjal")
	var myAge = 22
	fmt.Println("My Age is",myAge)

	//data type can be ommited if initial value is provided
	var withoutDataType int
	withoutDataType = 10
	fmt.Println("No without data type is ", withoutDataType)
	sayHello("siddhant");

	fmt.Println("now lets check some multipleReturn")
	add, sub := multipleReturn(100,200)
	fmt.Println("Addition is:", add," Substraction is", sub)
}

func sayHello(name string) string{
	fmt.Println("My master asked me to say hello.", name)
}


func multipleReturn(first, second float64) (float64, float64) {
	add := first+second;
	sub := first-second;
	return add, sub;
}