// main.go

package main

import (
	"fmt"

	"github.com/nbo2001/learnGO/person"
)

func main() {

	var n int
	var name string

	fmt.Scanf("%d", &n)

	fmt.Scan(&name)

	personIn := person.NewPerson(name, 25, 1.65)
	fmt.Println(personIn.GetName())
}
