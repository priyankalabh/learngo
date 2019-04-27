package main

import "fmt"

func squareArray(input []int) []int {
	newArray := []int{}
	for i := 0; i < len(input); i++ {
		newArray = append(newArray, input[i]*input[i])
	}
	return newArray
}

func modifiedArray(input []int) ([]int, []int) {
	evenArray := []int{}
	oddArray := []int{}
	for i := 0; i < len(input); i++ {
		if input[i]%2 == 0 {
			evenArray = append(evenArray, input[i]+1)
		} else {
			oddArray = append(oddArray, input[i])
		}
	}
	return evenArray, oddArray
}

func missingInArray(input []int, number int) int {
	sum := number * (number + 1) / 2
	arraySum := 0
	for i := 0; i < len(input); i++ {
		arraySum = arraySum + input[i]
	}
	return sum - arraySum
}

func main() {
	x := squareArray([]int{1, 2, 3, 4})
	x = append(x, 20, 25, 30)
	fmt.Printf("x : %+v\n", x)

	a, b := modifiedArray([]int{2, 3, 4, 5, 7, 9, 8})
	fmt.Printf("even numbers incremented bu 1:%v\nall odd values in array:%v", a, b)

	fmt.Printf("\nMissing number= %d", missingInArray([]int{1, 2, 3, 5, 6}, 6))
}

/*
1. Write a function which takes an array as an input and return a new array with all even numbers present in input array
2. Write a function which takes an array as an input and return two new output arrays. One output array should have all odd numbers
and second output array should have each even number incremented by 1.
*/
