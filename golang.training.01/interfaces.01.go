package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	for _, animal := range []Animal{Dog{}, Cat{} /*Llama{}, JavaProgrammer{}*/} {
		fmt.Println(animal.Speak())
	}
}
