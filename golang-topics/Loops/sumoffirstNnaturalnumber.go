package loops

import (
	"fmt"
)

func SumoffirstNnatural() {

	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println("Sum =", sum)

}
