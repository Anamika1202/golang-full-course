package loops

import "fmt"

func PrintReverseNumber(x int) {

	for x > 0 {
		lastdigit := x % 10
		fmt.Print(lastdigit)
		x = x / 10
	}

}
