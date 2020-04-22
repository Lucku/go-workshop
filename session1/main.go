package main

import "fmt"

type Direction struct {
	address string
	number int
}

type Person struct {
	firstName string
	lastName string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func NewPerson(firstName, lastName string, age int, direction Direction) *Person {
	return &Person{
		firstName:      firstName,
		lastName: lastName,
		age:       age,
	}
}

func myFunction(input1 string, input2 string) (string, int) {

	concat := input1 + input2

	return concat, len(concat)
}


func main() {

	/*
	var c = 1.5
	var d = "Test"
	var e float64 = "2.22"
	var f int64 = 16
	var g complex64 = 1.5+37i
	 */

	var fu func(string, string) (string, error)

	fu = func(var1, var2 string) (string, error) {
		return var1 + var2, nil
	}

	fmt.Println(fu)

	var p *Person
	
	p = &Person{
		firstName: "Gerrit",
		lastName: "Kasprik",
		age:       30,
	}

	h := []string{"one", "two", "three"}

	fmt.Println(h)

	a := 12
	b := &a

	fmt.Println("Before: ", a, *b)

	*b = 20

	fmt.Println("After new assignment of *b: ", a, *b)

	fmt.Printf("%v\n", p)
}
