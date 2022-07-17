package main

var value []string

func Fizz(n int) []string {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			value = append(value, "fizzbuzz")
		} else if i%5 == 0 {
			value = append(value, "fizz")
		} else if i%3 == 0 {
			value = append(value, "buzz")
		} else {

			value = append(value, "1")
		}
	}
	return value
}
