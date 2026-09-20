// DOCX掲載コード：4.3 ジェネリクスを1つ使う
package main

import "fmt"

func contains[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(contains([]int{1, 2, 3}, 2))
	fmt.Println(contains([]string{"Go", "C#"}, "Rust"))
}
