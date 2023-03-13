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

var slice12 []string = []string {"tee", "taye", "mike"}
