package main

import (
	"bufio"
	"fmt"
	"os"
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

	reader := bufio.NewReader(os.Stdin)
 fmt.Print("Enter first number: ")

 number, error := reader.ReadString('\n')

 if error != nil {
  fmt.Println("Error reading input:", error)
  return
 }
 fmt.Println("You entered:", number)





}
