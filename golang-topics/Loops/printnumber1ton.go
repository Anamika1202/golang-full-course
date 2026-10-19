package loops

import "fmt"

func PrintNumber1toN() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Printf("%d\n")

	}
}
