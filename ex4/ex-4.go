package main4

import "fmt"

func add(a int, b int) int {
	return a + b
}

func swap(str1 string, str2 string) (string, string) {
	return str2, str1
}

func divide(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

func main() {
	num1 := 5
	num2 := 10
	sum := add(num1, num2)

	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)

	// ################
	first := "Hello"
	second := "World"
	swapped1, swapped2 := swap(first, second)

	fmt.Printf("Original: %s, %s\n", first, second)
	fmt.Printf("Swapped: %s, %s\n", swapped1, swapped2)

	// ################
	num3 := 17
	num4 := 5
	quotient, remainder := divide(num3, num4)

	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
