package introdution

import "fmt"

func CalculateSimpleInterest() {

	var principal, rate, time float64

	fmt.Println("Enter principal , rate and time: ")
	fmt.Scan(&principal, &rate, &time)
	si := principal * rate * time / 100
	fmt.Println("Simple Interest =", si)
}
