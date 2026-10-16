package conditionalstatement

import "fmt"

func findPositive_negativenumber() {

	var n int
	fmt.Scan(&n)
	if n >= 0 {
		fmt.Printf("Number Is Positive %d", n)
	} else {
		fmt.Printf("Number is Negative%d", n)
	}

}
