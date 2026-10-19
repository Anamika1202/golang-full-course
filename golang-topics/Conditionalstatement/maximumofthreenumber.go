package conditionalstatement

import "fmt"

func Maxofthreenumber() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	max := 0
	if a > b && a > c {
		max = a
	}
	if b > a && b > c {
		max = b
	} else {
		max = c
	}

	fmt.Println("Maximum =", max)
}
