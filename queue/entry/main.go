package main

import (
	"fmt"

	"github.com/sniancmk/learnGo/queue"
)

func main() {
	q := queue.Queue{}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
}
