package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	grades := make(map[string]int)

	var name string

	// take name
	for {

		fmt.Println("Enter your name: ")
		fmt.Scanln(&name)

		if strings.TrimSpace(name) == ""{
			fmt.Println("Error: Name can't be empty. Please enter a valid name")
		} else if _, err := strconv.Atoi(name) ; err == nil{
			fmt.Println("Error: Name can't be a number. Please enter a valid name")
		} else{
			break
		}
	}

	//take number of subjects

	var numSubjects int

	for {
		fmt.Println("How many grades are you taking: ")
		_, err := fmt.Scanln(&numSubjects)

		//validating input
		if (err != nil){
			fmt.Println("Please enter a valid number")
		}else if numSubjects <= 0 {
			fmt.Println("Invalid number of subjects.")
		} else{

			break
		}

	}
	
	//take subject name and grade
	var subject string
	var grade int

	fmt.Println("Enter the subject and the respective grade (Space Separated)")

	for i := 0 ; i < numSubjects; i++{
		_, err := fmt.Scan(&subject, &grade)
		
		if (err != nil){
			fmt.Println("Error: Please enter valid input")
			return
		} else if (0 > grade || grade > 100){
			fmt.Println("Grade should be between 0 and 100")
			return
		}
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