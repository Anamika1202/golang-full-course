package VariabledataType

import (
	"fmt"
)

func Calculateaverage() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Println("Average = %.2f\\n", a+b+c/3) //%.2f 2 decimal places 12.57
}
