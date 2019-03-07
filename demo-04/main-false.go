package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for v := range s {     // !!!
		sum += v
	}

	c <- sum 	       // send sum to c
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	c := chan int 	       // !!!
	go sum(s1, c)          // [1 2 3]
	go sum(s2, c)          // [4 5 6]
	x, y := <-c, <-c       // receive from c
	fmt.Println(x, y, x+y) //output: 15 6 21
}
