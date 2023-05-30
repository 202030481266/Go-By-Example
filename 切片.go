package main

import "fmt"

func main() {
	s := make([]string, 3) // 声明一个长度为3的string切片
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) // 直接打印整个slice
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s)) // 打印slice的长度

	fmt.Println("--------------------------")

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended:", s)

	fmt.Println("--------------------------")

	c := make([]string, len(s))
	copy(c, s) // 直接拷贝
	fmt.Println("copy:", c)
	l := s[2:5]
	fmt.Println("slice from 2 to 5:", l)

	fmt.Println("--------------------------")
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	fmt.Println("--------------------------")
	twoD := make([][]int, 3) // 固定第一维的slice
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d array:", twoD) // 直接打印整个二维slice真的方便
}
