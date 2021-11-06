package main

import "fmt"

// 1 - 1
// 2 - 1
// 3 - 2
// 4 - 3
// and so on...
func fibonacci(num int64) int64 {
	if num < 3 {
		return 1
	}
	return 2*fibonacci(num-2) + fibonacci(num-3)
}

func main() {
	var res, num int64
	fmt.Scan(&num)
	res = fibonacci(num)
	fmt.Print("phi(", num, ") = ", res)
}
