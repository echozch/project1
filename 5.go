package main

import (
	"fmt"
	"math"
)

func main() {
	var r, area float64
	fmt.Scanln(&r)
	area = r * r * math.Pi
	fmt.Println(area)
}
