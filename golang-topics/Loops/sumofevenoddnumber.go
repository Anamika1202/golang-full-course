package loops

import (
	"fmt"
)

func SumofEvenOddNumber() {

	var evensum uint
	var oddsum uint
	var num uint
	var choice uint

	for {

		fmt.Println("Enter the number :")
		fmt.Scan(&num)

		if num%2 == 0 {
			evensum = evensum + num
		} else {
			oddsum = oddsum + num
		}

		fmt.Printf("Do you want to continue? Press 1 for yes or 0 for no")
		fmt.Scan(&choice)
		if choice == 0 {
			fmt.Printf("Even Number Sum is : %d\n ", evensum)
			fmt.Printf("ODD Number Sum is %d\n ", oddsum)
			break
		}
	}

}
