package main

import (
	"fmt"
)

func main() {

	fmt.Println("test")

	for a := 3; a <= 100; a++ {
		flag := 0
		for i := 2; i < a; i++ {
			if a != i {
				if a%i == 0 {
					flag = 1

				}
			}
		}
		if flag == 1 {
			//fmt.Println("Not a prime number")
		} else {
			fmt.Println(a)
		}
	}
}
