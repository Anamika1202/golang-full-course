package loops

import (
	"fmt"
)

func Checkprimenumber() {
	var n int
	fmt.Scan(&n)
	prime := n > 1

	for i := 2; i < n && prime; i++ {

		if n%i == 0 {
			prime = false
		}

	}
	if prime {
		fmt.Println("Prime nūmber")
	} else {
		fmt.Println("Not a prime number")
	}
}
