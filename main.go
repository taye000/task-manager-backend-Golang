package main

// Importing fmt package to format & print to the console
import "fmt"

func main() {
	//initializing variables with var keyword and type
	var nameTwo string = "Tgod"
	var ageOne int = 23

	//initializing variables with var keyword and inferred type
	var nameThree = "Tee"
	var ageTwo = 23

	//initializing variables with shorthand syntax
	nameOne := "taye"
	ageThree := 23

	//initializing float variables
	var scoreOne float32 = 23.4
	var scoreTwo float64 = 23345346345.4
	scoreThree := 24353453.4

	//initializing array variables. Arrays are fixed length lists of items of the same type.
	//array variable with type
	var arrOne [3]int = [3]int{1, 2, 3}
	//array variable with inferred type
	var arrTwo = [3]int{2, 3, 4}
	//array variable with shorthand syntax
	arrThree := [3]int{3, 4, 5}

	fmt.Println(arrOne, arrTwo, arrThree, len(arrOne))

	//printing variables to the console using fmt.Println
	fmt.Println(nameOne, ageThree, nameTwo, ageOne, nameThree, ageTwo)
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//printing variables to the console using fmt.Printf(formatted strings)
	fmt.Printf("my name is %v, and I am %v years old \n", nameOne, ageThree)
	//using %q to print strings with quotes
	fmt.Printf("my name is %q, and I am %v years old \n", nameOne, ageThree)
	//using %T to print the type of the variable
	fmt.Printf("my name is type %T, and my age is type %T \n", nameOne, ageThree)
	//using %f to print float variables
	fmt.Printf("my age is %f \n", scoreOne)
	//using %f to print float variables with 2 decimal places
	fmt.Printf("my age is %0.2f \n", scoreOne)

	//using Sprintf to format and store the formatted string in a variable
	var str = fmt.Sprintf("my name is %v, and I am %v years old \n", nameOne, ageThree)
	fmt.Println("str:", str)
}
