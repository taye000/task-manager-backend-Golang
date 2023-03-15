package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
		items: map[string]float64{},
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
	/*add tip
	the -25 in %-25v, adds 25 characters to the string*/
	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "Tip:", b.tip)

	/*total
	the -25 in %-25v, adds 25 characters to the string*/
	fs += fmt.Sprintf("%-25v ... $%0.2f \n", "Total:", total+b.tip)

	return fs
}

//update the tip. Another method of the bill struct
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//add items to the bill. Another method of the bill struct
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

//helper function to take any input from the user via terminal(stdin)
func takeInput(prompt string, r *bufio.Reader) (string, error) {
	//print the prompt
	fmt.Print(prompt)
	//read the input from the terminal
	input, err := r.ReadString('\n')
	//remove the new line character or any trailing white spaces and return the input
	return strings.TrimSpace(input), err
}

//function to add items to the bill
func promptOptions(b bill) {
	//create a new reader
	reader := bufio.NewReader(os.Stdin)

	//prompt the user to choose an option
	opt, _ := takeInput("Choose option (a - add item, s - Save bill, t - add tip): ", reader)

	//switch statement to handle the different options
	switch opt {
	case "a":
		name, _ := takeInput("Enter item name: ", reader)
		price, _ := takeInput("Enter item price: ", reader)

		//convert price string to float64
		p, err := strconv.ParseFloat(price, 64)
		//check if there is an error
		if err != nil {
			fmt.Println("Input has to be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Printf("item %v added, with price %v",name, p)
	case "s":
		fmt.Println("You chose s")
	case "t":
		tip, _ := takeInput("Enter tip amount: ", reader)
		fmt.Println(tip)
	default: //default option, when the user enters an invalid option
		fmt.Println("Not a valid option")
		promptOptions(b) //call the function again to prompt the user to choose an option
	}
}

//function to take input from the user via terminal(stdin)
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	//create a new bill using the newBill function and the input from the user
	name, _ := takeInput("Create your bill, name: ", reader)

	fmt.Printf("Created the bill - %v \n", name)

	b := newBill(name)

	return b
}
