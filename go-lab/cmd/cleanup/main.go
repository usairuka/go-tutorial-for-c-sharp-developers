// DOCX掲載コード：defer
package main

import "fmt"

func main() {
	n := 1
	defer fmt.Println("deferred:", n)
	defer fmt.Println("last registered")
	n = 2
	fmt.Println("now:", n)
}
