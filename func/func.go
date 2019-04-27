package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func sumEvenMultOdd(num int) (int, int) {
	sum := 0
	prod := 1
	ncopy := num
	for ncopy > 0 {
		digit := ncopy % 10
		if digit%2 == 0 {
			sum = sum + digit
		} else {
			prod = prod * digit
		}
		ncopy = ncopy / 10
	}
	return sum, prod
}

func checkArmstrong(num int) bool {
	count := 0
	sum := 0
	ncopy := num
	for ncopy > 0 {
		ncopy = ncopy / 10
		count++
	}
	if count != 3 {
		return false
	}
	ncopy = num
	for ncopy > 0 {
		digit := ncopy % 10
		sum = sum + (digit * digit * digit)
		ncopy = ncopy / 10
	}
	if num == sum {
		return true
	}
	return false
}

func squareRoot(num int) int {
	notApplicable := -1
	for i := 1; i < num/2; i++ {
		if num == i*i {
			return i
		}
		if i*i > num {
			return notApplicable
		}
	}
	return notApplicable
}

func calculateMiddleDigit(num int) int {
	digit := -1
	count := 0
	ncopy := num
	for ncopy > 0 {
		ncopy = ncopy / 10
		count++
	}
	if count%2 == 0 {
		return -1
	}
	ncopy = num
	for i := 0; i < count/2+1; i++ {
		digit = ncopy % 10
		ncopy = ncopy / 10
	}
	return digit
}

func countZerosAndOnes(num int) (int, int) {
	countz := 0
	counto := 0
	for num > 0 {
		digit := num % 10
		if digit == 0 {
			countz++
		}
		if digit == 1 {
			counto++
		}
		num = num / 10
	}
	return countz, counto

}

func main() {
	a := 10
	b := 20
	a, b = swap(a, b)
	sum, prod := sumEvenMultOdd(23456789)
	fmt.Printf("After swaping the number a:%d\tb:%d", a, b)
	fmt.Printf("\nThe sum of even digits are:%d\n The product of odd digits:%d", sum, prod)
	fmt.Printf("\nis armstrong number\t%t", checkArmstrong(153))
	fmt.Printf("\nsquareroot:%d", squareRoot(20))
	x, y := countZerosAndOnes(21001)
	fmt.Printf("\nNumber of zeros= %d\nThe number of ones=%d", x, y)
	fmt.Printf("\nThe middle number= %d", calculateMiddleDigit(23456))
}
