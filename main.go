package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {

	m := map[string]string{"a": "1"}
	s2 := m["a"]
	fmt.Println(s2)
}
