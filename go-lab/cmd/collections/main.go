// DOCX掲載コード：2.1 実行して共有を観察する
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	array := [3]int{1, 2, 3}
	arrayCopy := array
	arrayCopy[0] = 99
	fmt.Println("arrays:", array, arrayCopy)

	xs := []int{1, 2, 3}
	window := xs[:2]
	window[0] = 99
	fmt.Println("shared:", xs)

	cloned := make([]int, len(xs))
	copy(cloned, xs)
	cloned[0] = 7
	fmt.Println("copied:", xs, cloned)

	var selected []int
	for _, n := range xs {
		if n%2 != 0 {
			selected = append(selected, n)
		}
	}
	fmt.Println("odd:", selected)

	scores := map[string]int{"Go": 80}
	fmt.Println("present:", scores["Go"])
	value, ok := scores["C#"]
	fmt.Println("missing:", value, ok)
	scores["C#"] = 90
	delete(scores, "Go")
	fmt.Println("entries:", len(scores))

	s := "Go言語"
	fmt.Println("text:", len(s), utf8.RuneCountInString(s))
	for byteIndex, r := range s {
		fmt.Printf("%d:%c ", byteIndex, r)
	}
	fmt.Println()
}

// 2.5 必須演習の解答例
func filterAtLeast(values []int, min int) []int {
	result := make([]int, 0, len(values))
	for _, value := range values {
		if value >= min {
			result = append(result, value)
		}
	}
	return result
}
