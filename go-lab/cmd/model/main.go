// DOCX掲載コード：3.2 ポインタと暗黙のinterface実装
package main

import (
	"errors"
	"fmt"

	"example.com/go-lab/internal/task"
)

type summarizer interface {
	Summary() string
}

func printSummary(s summarizer) {
	fmt.Println(s.Summary())
}

func main() {
	t, err := task.New(1, " Learn Go ")
	if err != nil {
		fmt.Println(err)
		return
	}

	copied := t
	copied.Complete()
	printSummary(t)
	printSummary(copied)

	p := &t
	p.Complete()
	printSummary(t)

	_, err = task.New(2, " ")
	if errors.Is(err, task.ErrEmptyTitle) {
		fmt.Println("validation:", err)
	}
}
