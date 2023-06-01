package _1_basics

import "fmt"

func FizzBuzz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(number)
}
