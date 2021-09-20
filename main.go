package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ex1()
	// ex2()
	// ex3()
}

func ex1() {
	var x int
	fmt.Print("Pleas input number : ")
	fmt.Scanln(&x)
	fmt.Println("------------------------")

	for i := 1; i <= x; i++ {
		fmt.Print("* ")
	}
	fmt.Println()

	for i := 0; i < x-2; i++ {
		fmt.Print("* ")

		for j := x; j > 2; j-- {
			fmt.Print("  ")
		}
		fmt.Println("*")
	}

	for i := 1; i <= x; i++ {
		fmt.Print("* ")
	}
}

func ex2() {
	var number int
	var remainder int
	var temp int
	var reverse int = 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Print(temp, " is a palindrome.")
		return
	} else {
		fmt.Print(temp, " is not a palindrome.")
		return
	}
}

func ex3() {
	fmt.Print("Enter a number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Text()
	temp := strings.Split(in, " ")

	ary := make([]int, len(temp))
	for i := range ary {
		ary[i], _ = strconv.Atoi(temp[i])
	}

	smallestNumber := ary[0]
	for _, element := range ary {
		if element < smallestNumber {
			smallestNumber = element
		}
	}
	fmt.Println("Smallest number is  ", smallestNumber)
}