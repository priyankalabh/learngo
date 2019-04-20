package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i <= num-1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sumOddNumbers(num int) int {
	var sum int
	for i := 1; i <= num; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	return sum
}

func calculateDigits(num int) int {
	if num == 0 {
		return 1
	}
	var count int
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func main() {
	fmt.Printf("this number is prime:%t\n", isPrime(4))
	fmt.Printf("the sum of odd numbers:%d\n", sumOddNumbers(12))
	fmt.Printf("the digits in a number are:%d", calculateDigits(1))

}
