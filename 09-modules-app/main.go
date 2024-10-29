package main

import (
	"fmt"

	"github.com/tkmagesh/Nutanix-GoFoundation-Oct-2024/09-modules-app/calculator"
	// "github.com/tkmagesh/Nutanix-GoFoundation-Oct-2024/09-modules-app/calculator/utils"

	"github.com/fatih/color"
	ut "github.com/tkmagesh/Nutanix-GoFoundation-Oct-2024/09-modules-app/calculator/utils"
)

func main() {
	color.Red("modules app executed!")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.Multiply(100, 200))
	fmt.Println("Operation count :", calculator.OpCount())
	// fmt.Println("is 21 odd number ?", utils.IsOdd(21))
	fmt.Println("is 21 odd number ?", ut.IsOdd(21))
}
