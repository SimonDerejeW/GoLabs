package main

import "fmt"

func sendData(channel chan int){
	for i:=0; i < 5; i++{
		channel <- i
	}
	close(channel)
}

func main() {
	//Create Channel
	ch:= make(chan int)

	go sendData(ch)

	for val:= range ch{
		fmt.Println("Received: ", val)
	}

	fmt.Println("Channel Closed, Program Finished!")
}