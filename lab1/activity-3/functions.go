package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main(){
	sum:= add(5 , 3);

	//If Statement
	if (sum > 5){
		fmt.Println("Sum is greater than 5")
	} else {
		fmt.Println("Sum is less than or equal to 5")
	}

	//For loop
	for i:= 0; i <= 5; i++{
		fmt.Println("Loop Iteration: ", i)
	}

}