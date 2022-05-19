package main

import (
	"fmt"
)

func main() {
	s := makeRange(0, 10)
	// fmt.Printf("var1 = %T\n", s)               // check the type of variable
	// fmt.Printf("var2 = %T\n", float64(s[1]/2)) // check the type of variable
	for i := range s {
		//a := float64((s[i] % 2)) // transfor to float64
		if s[i]%2 == 0 {
			fmt.Println(s[i], "is even")
		} else {
			fmt.Println(s[i], "is odd")
		}
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
