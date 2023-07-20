package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Assign values by struc fields order.
	fmt.Println(person{"Bob", 20})
	// Assign specific field values.
	fmt.Println(person{name: "Alice", age: 30})
	// If not defines, assumes zero-value
	fmt.Println(person{name: "Fred"})
	// Pointer to a new instance of 'peson' struct
	fmt.Println(&person{name: "Ann", age: 40})
	// Using a factory function of structs 'person'
	fmt.Println(newPerson("Jon"))

	// Storing struct into a variable.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Automatic dereferencing
	// Getting struct values direct from it's pointer.
	sp := &s
	fmt.Println(sp.age)
	// Setting struct values direct from it's pointer.
	sp.age = 51
	fmt.Println(sp.age)

	// Defining struct and directly assign field values to a variable.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
