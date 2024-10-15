package main

import "fmt"

func main(){

	//Array
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	//Slice(Dynamic Array)
	slice := []int{4,5,6}
	slice = append(slice, 7)
	fmt.Println(slice)

	//Maps (Key-Value Pairs)
	myMap := make(map[string]int)
	myMap["Alice"] = 25
	myMap["Bob"] = 50
	fmt.Println("Map: ",myMap)
	fmt.Println("Alice's Age: ", myMap["Alice"])

	//Looping over a slice
	for index, val := range slice{
		fmt.Printf("Index: %d, Value: %d\n", index, val)
	}

	//Looping over a map
	for key, value := range myMap{
		fmt.Printf("%s is %d years old!\n", key, value)
	}
}