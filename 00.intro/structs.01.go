package main

import "fmt"

// type person struct { // What's the difference
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{Name: "Alice", Age: 30}) // You can name the fields.
	fmt.Println(Person{Name: "Fred"})           // You don't need to specify all the params.
	fmt.Println(&Person{Name: "Ann", Age: 40})  // You can also create a pointer.

	s := Person{Name: "Sean", Age: 50}
	fmt.Println(s.Name)

	sp := &s
	fmt.Println(sp.Age)
	sp.Age = 51
	fmt.Println(sp.Age)
}
