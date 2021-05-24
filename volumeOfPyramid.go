// Created by: Maryam Nona
// Created on: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main() {
	// This function does addition
	var length int
	var width int
	var height int

	// input
	fmt.Println("This program finds the volume of a pyramid")
	fmt.Println()
	fmt.Print("Enter the length (mm): ")
	fmt.Scanln(&length)
	fmt.Println("Enter the width (mm): ")
	fmt.Scanln(&width)
	fmt.Println("Enter the height (mm) ")
	fmt.Scanln(&height)

	// output
	fmt.Println("The volume is: ", (length*width*height)/3, "mm³")
	fmt.Println("Done.")
}
