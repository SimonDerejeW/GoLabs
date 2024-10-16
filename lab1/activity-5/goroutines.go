package main

import (
	"fmt"
	"time"
)

func printNumbers(){
	for i:= 0; i <= 5; i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}


func printLetters(){
	for i:= 'A'; i <= 'E'; i++ {
		fmt.Printf("%c\n",i)
		time.Sleep(time.Second)
	}
}

func main(){

	go printNumbers()

	go printLetters()

	time.Sleep(6 * time.Second)
	fmt.Println("Main Function has finished!")
}