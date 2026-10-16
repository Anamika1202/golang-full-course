package conditionalstatement

import "fmt"

func FindLeapYear() {

	var year int
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("not a leap year")
	}

}
