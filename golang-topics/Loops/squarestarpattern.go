package loops

import "fmt"

func SquareStartPattern() {
	n := 4

	for i := 1; i <= n; i++ {

		for j := 1; j <= n; j++ {

			fmt.Print("*")

		}

		fmt.Println()

	}

}
