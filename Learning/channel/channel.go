package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 2)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	s1 := []int{7, 2, 8, -9, 4, 0}

	go sum(s1[:len(s)/2], c)
	fmt.Println(s1[:len(s)/2])
	go sum(s1[len(s)/2:], c)
	fmt.Println(s1[len(s)/2:])
	x1, y1 := <-c, <-c // receive from c

	fmt.Println(x1, y1, x1+y1)
}
