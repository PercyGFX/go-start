package main

import (
	"fmt"
	"go-start/testpackage"
)


func getQuotation (x float64, y float64) (ans float64, err error) {
	if x < 0 || y < 0 {
		err = fmt.Errorf("negative values are not allowed")
		return
	}
	ans = x * y
	return
}

func main() {
	fmt.Println("Hello, World!")

	ans, err := getQuotation(10, -1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotation:", ans)
	}

	var name string = "Jone"

	pakaya := 8
	
	if (pakaya > 5) {
		fmt.Println("pakaya is greater than 5")
	} else {
		fmt.Println("pakaya is less than or equal to 5")
	}

	for x:= 1 ; x<= 10; x++ {
		fmt.Println(x)
	}



	testpackage.Test(name)
}
