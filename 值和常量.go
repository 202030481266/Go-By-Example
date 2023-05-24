package main

import (
	"fmt"
	"math"
)

const s string = "xiaoshulin"

func main() {
	fmt.Println(s)
	const n = 50000000
	const d = 3e20 / n
	fmt.Println(math.Sin(n))
	fmt.Println(int64(d))
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(false)
	fmt.Println(true)
}
