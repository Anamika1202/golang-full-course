package loops

import "fmt"

func FindFactorialNumber() {

	var num uint
	var fact uint

	fmt.Println("Enter The Number: ")
	fmt.Scan(&num)

	for i := uint(1); i <= num; i++ {
		fact = fact * i
	}

	fmt.Println("Factorial Of Number is %d", fact)

}
