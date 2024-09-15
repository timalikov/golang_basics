package main3

import "fmt"

func main() {
	var number int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	// #################
	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println("The sum of the first 10 natural numbers is:", sum)

	// #################
	var day int

	fmt.Print("Enter a number (1 for Monday, 2 for Tuesday, etc.): ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input. Please enter a number between 1 and 7.")
	}

}
