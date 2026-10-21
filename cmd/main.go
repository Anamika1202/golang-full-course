package main

import (
	"fmt"
	loops "golang-full-course/golang-topics/Loops"
)

func main() {
	//introdution.Sumoftwonumber()
	//introdution.CalculateSimpleInterest()

	//introdution.Maxofthreenumber()
	//introdution.Checkprimenumber()

	//loops.PrintNumber()
	//.SquareStartPattern()

	fmt.Println("Enter the Number you want to Reverse")
	var num uint

	fmt.Scan(&num)

	revnum := loops.ReversetheGivenNumber(num)
	fmt.Printf("reverse Number Is %d", revnum)

}
