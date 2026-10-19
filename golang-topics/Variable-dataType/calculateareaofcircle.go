package VariabledataType

import "fmt"

func Calculateareofcircle() {
	var radius float64
	fmt.Scan(&radius)
	area := 3.14 * radius * radius
	fmt.Printf("Area = %2.f\\n", area) //0 decimal places, minimum width 2 %2.f 13
}
