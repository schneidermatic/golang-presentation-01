package main

import (
	"fmt"
	"strconv"
)

func say_hello(message string) {
	fmt.Println(message)

}

func main() {


Loop:
	for x := 0; x < 5; x++ {
		fmt.Println("x: " + strconv.Itoa(x))
		for y := 0; y < 5; y++ {
			if x >= 2 {
		                say_hello("Hello, 世界")
				break Loop
			}
		}
	}

}
