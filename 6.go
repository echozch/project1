package main

import "fmt"

func main() {
	var n, j, a int
	a = 1
	fmt.Scanln(&n)
	for j = 2; j < n; j++ {
		if n%j == 0 {
			a = 0
			fmt.Printf("%d不是素数", n)
			break
		}
	}
	if a == 1 {
		fmt.Printf("%d是素数", n)
	}
}
