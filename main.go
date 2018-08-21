package main

import "fmt"

/* FizzBuzz - prints every number 1 thru 100,
numbers divisible by 3 prints Fizz,
numbers divisible by 5 prints Buzz,
and numbers divisible by both prints FizzBuzz */

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, " FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, " Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
