// main.go

package main

import (
	"fmt"
)

func main() {
	person := NewPerson("Alice", 25, 1.65)
	fmt.Println(person.getName())
}
