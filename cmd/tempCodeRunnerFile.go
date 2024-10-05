package main

import (
	"fmt"

	queue "github.com/kupalovmuhammadjon/Queue"
)

func main() {
	q := queue.NewQueue()
	q.Push("qq")
	fmt.Println(q.Pop())
}
