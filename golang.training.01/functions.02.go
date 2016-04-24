package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Speak() (bool, error) {
	fmt.Println("My name is", p.Name, "and i'm", p.Age, "years old")
	return true, nil
}

func main() {
	jose := Person{
		Name: "Jose",
	}

	spoke, err := jose.Speak()
	if err == nil {
		fmt.Println(spoke)
	}
}
