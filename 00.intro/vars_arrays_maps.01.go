package main

import "fmt"

func main() {
	var a int
	a = 0
	//a := 1 Would fail

	b := 1

	c := make([]string, 2)
	d := []string{"foo", "bar"}

	e := make(map[string]string)

	e["foo"] = "bar"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
