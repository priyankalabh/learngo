/*
Basic programs like Hello world, arithmetic operations, looping, switch
*/
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(divisor, dividend int) int {
	return divisor / dividend
}

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func half(x int) int {
	return x / 2
}

func double(x int) int {
	return 2 * x
}

func makeDoubleOrHalf(num int) int {
	if isEven(num) {
		return half(num)
	}
	return double(num)
}

func main() {
	fmt.Printf("The sum of number is: %d\n", add(10, 12))
	fmt.Printf("The subtraction of number is: %d\n", sub(20, 16))
	fmt.Printf("The product of number is: %d\n", mult(12, 8))
	fmt.Printf("The division of number is: %d\n", div(65, 5))
	fmt.Printf("The number is even: %t\n", isEven(10))
	fmt.Printf("The number is: %d\n", makeDoubleOrHalf(10))
}
