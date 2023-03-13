package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Hello %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Bye %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

//using strings package
func getInitials(n string) (string, string) {
	//splitting the strings into a slice
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

//using math package
func areaOfCircle(r float64) float64 {
	return math.Pi * r * r
}

//slice
var slice12 []string = []string{"tee", "taye", "mike"}

//map with key(string) and value(float64)
var menu = map[string]float64{
	"chicken": 150.00,
	"beef":    200.50,
	"fish":    100.70,
}

//structs. A struct is a collection of fields & represents a custom data type
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"chilli": 103.99, "rice": 499.00, "noddles": 100.00},
		tip:   0,
	}
	return b
}

//format the bill. This is a method of the bill struct. It is a function that is attached to a struct.
func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0
	//iterate over items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	/*total
	the -25 in %-25v, adds 25 characters to the string*/
	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "Total:", total)

	return fs
}
