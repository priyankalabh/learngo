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

func reverseNumber(num int) int {
	var revnum int
	n := calculateDigits(num)
	prod := 1
	for i := 1; i < n; i++ {
		prod = prod * 10
	}

	for num > 0 {
		digit := num % 10
		revnum = revnum + (digit * prod)
		num = num / 10
		prod = prod / 10
	}
	return revnum
}

func factorial(num int) int {
	var fact = 1
	for num > 0 {
		fact = fact * num
		num = num - 1
	}
	return fact
}

func powerNumber(num, power int) int {
	n := 1
	for i := 0; i < power; i++ {
		n = n * num
	}
	return n
}

func prodnum(num int) int {
	if num == 0 {
		return 0
	}
	prod := 1
	for num > 0 {
		digit := num % 10
		prod = prod * digit
		num = num / 10
	}
	return prod
}

func sumOfAllDigit(num int) int {
	sum := 0
	for num > 0 {

		digit := num % 10
		sum = sum + digit
		num = num / 10
	}
	return sum
}

func main() {
	fmt.Printf("this number is prime:%t\n", isPrime(4))
	fmt.Printf("the sum of odd numbers:%d\n", sumOddNumbers(12))
	fmt.Printf("the digits in a number are:%d", calculateDigits(1))
	fmt.Printf("the reverse of digits in a number are:%d", reverseNumber(1345))
	fmt.Printf("\nthe factorial of a number are:%d\n", factorial(5))
	fmt.Printf("\nthe power in a number are:%d", powerNumber(3, 2))
	fmt.Printf("\nthe product of a number are:%d", prodnum(0))
	fmt.Printf("\nthe sum of the digits in a number are:%d", sumOfAllDigit(123))

}
