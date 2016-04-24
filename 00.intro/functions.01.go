package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Speak(person Person) (bool, error) {
	fmt.Println("My name is", person.Name, "and i'm", person.Age, "years old")
	return true, nil
}

func main() {
	jose := Person{
		Name: "Jose",
		Age:  36,
	}

	spoke, _ := Speak(jose)

	fmt.Println(spoke)
}
