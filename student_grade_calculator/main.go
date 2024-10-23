package main

import (
	"fmt"
)

func main(){
	grades := make(map[string]int)

	var name string

	// take name
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	//take number of subjects

	var numSubjects int
	
	fmt.Println("How many grades are you taking: ")
	fmt.Scanln(&numSubjects)

	if numSubjects <= 0 {
		fmt.Println("Invalid number of subjects.")
		return
	}

	//take subject name and grade
	var subject string
	var grade int

	fmt.Println("Enter the subject and the respective grade (Space Separated)")

	for i := 0 ; i < numSubjects; i++{
		fmt.Scan(&subject, &grade)
		grades[subject] = grade
	}

	fmt.Printf("Hello, %s! Here are your grades: \n", name)

	var total = 0

	for subject, grade := range grades{
		fmt.Printf("Subject %s: %d \n", subject, grade)
		total += grade
	}

	average := float64(total) / float64(numSubjects)

	fmt.Printf("%s, your average is %.2f", name, average)

}