package main

/*Importing fmt packages
fmt is used to print to the console
strings is used to work with strings*/
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//initializing variables with var keyword and type
	var nameTwo string = "Tgod major"
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

	/*initializing array variables. Arrays are fixed length lists of items of the same type.
	array variable with type*/
	var arrOne [3]int = [3]int{1, 2, 3}
	//array variable with inferred type
	var arrTwo = [3]int{2, 3, 4}
	//array variable with shorthand syntax
	arrThree := [3]int{3, 4, 5}

	//update a value from the array
	arrThree[1] = 10

	//slices are like arrays but they are dynamic in size
	var sliceOne = []int{9, 8, 4, 3}
	//update a value from the slice
	sliceOne[0] = 20
	/*append a value to the slice
	append returns a new slice with the appended value
	It does not update the slice, unless you assign the new returned
	slice to be the updated slice as below*/
	sliceOne = append(sliceOne, 5)

	//slice ranges, sliceOne[startIndexIncluding:upToNotIncluding]
	rangeOne := sliceOne[0:2]
	//sliceOne[1:] means from index 1 to the end of the slice
	rangeTwo := sliceOne[1:]
	//sliceOne[:2] means from the start of the slice to index 2 not including index 2
	rangeThree := sliceOne[:2]

	//updating nameTwo variable with the new string returned from the ReplaceAll method
	nameTwo = strings.ReplaceAll(nameTwo, "major", "minor")

	//using sort package to sort the slice. Sort is in-place, meaning it updates the original slice
	sort.Ints(sliceOne)
	fmt.Println("sliceOne:", sliceOne)

	//using strings package to print to the console
	fmt.Println(strings.Contains(nameTwo, "major"))
	/*using strings package Replace method to replace a string with another string
	It returns a new string, Not updating the original*/
	fmt.Println(strings.ReplaceAll(nameTwo, "major", "minor"))
	//using strings.toUpper method to convert a string to uppercase
	fmt.Println(strings.ToUpper(nameTwo))
	//using the strings.Split method to split a string into a slice
	fmt.Println(strings.Split(nameTwo, " "))

	//printing out the ranges of the slices
	fmt.Println(rangeOne, rangeTwo, rangeThree)
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
