package main

import "fmt"


func main() {
	sumMult3Or5()
}

func sumMult3Or5() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Result: %d\n", sum)
}