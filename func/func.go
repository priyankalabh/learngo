package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}
func sumEvenMultOdd(num int) (int, int) {
	sum := 0
	prod := 1
	for num > 0 {
		digit := num % 10
		if digit%2 == 0 {
			sum = sum + digit
		} else {
			prod = prod * digit
		}
		num = num / 10
	}
	return sum, prod
}

func main() {
	a := 10
	b := 20
	a, b = swap(a, b)
	sum, prod := sumEvenMultOdd(23456789)
	fmt.Printf("After swaping the number a:%d\tb:%d", a, b)
	fmt.Printf("\n The sum of even digits are:%d\n The product of odd digits:%d", sum, prod)

}
