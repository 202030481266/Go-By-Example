package main

import (
	"fmt"
	"time"
)

func main() {
	// 在条件语句之前可以有一个声明语句；在这里声明的变量可以在这个语句所有的条件分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num)
	}
	for i := 1; i < 10; i = i + 1 {
		fmt.Println(i)
	}
	t := time.Now()
	fmt.Println(t.Hour(), t.Minute(), t.Second())
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend!")
	default:
		fmt.Println("Today is weekdays!")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("This is a bool variable!")
		case int:
			fmt.Println("This is a int variable!")
		default:
			fmt.Println("Unknown Type")
		}
	}
	whatAmI(true)
	whatAmI(23)
	whatAmI("xiaoshulin")
}
