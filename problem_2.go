package main

import "fmt"

func main() {
	sumFiboEven()
}

func sumFiboEven() {
	a := 0
	b := 1
	c := 0
	total := 0
	for b < 4000000 {
		c = b
		b += a
		a = c
		if b % 2 == 0 {
			total += b
		}
	}
	fmt.Println(total)
	
}