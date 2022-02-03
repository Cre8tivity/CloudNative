package main

/*
import "fmt"

func sum(s []int, c chan int) {
	fmt.Println("starting channel %i", c)
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sending sum %i to channel c", sum)
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
*/
