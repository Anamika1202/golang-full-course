package VariabledataType

import "fmt"

func Generate_bill() {

	var pencil, pen, eraser float64
	fmt.Scan(&pencil, &pen, &eraser)
	subtotal := pencil + pen + eraser
	gst := subtotal + 0.18
	fmt.Printf("Subtotal: %.2f\\nGST (18%): %2f\\nTotal bill : %.2f\\n", subtotal, gst, subtotal+gst)

}
