package main

import (
	"fmt"
	"strconv"
)

// creating a map to better scale the program for more comparisons
var values = map[int]string{
	3: "Fizz",
	5: "Buzz",
	7: "Bazz",
}

// The FizzBuzz challenge typically requires the participant to write a program that prints a sequence of numbers from 1 to a given number, but with some variations:

// If the number is divisible by 3, print "Fizz" instead of the number.
// If the number is divisible by 5, print "Buzz" instead of the number.
// If the number is divisible by both 3 and 5, print "FizzBuzz" instead of the number.
func fizzbuzz(n int) {
	for i := 0; i <= n; i++ {
		str := ""

		for k, v := range values {
			if i%k == 0 {
				str += v
			}
		}

		if str == "" {
			str += strconv.Itoa(i)
		}

		fmt.Println(str)

	}
}

func main() {
	fizzbuzz(10)
}
