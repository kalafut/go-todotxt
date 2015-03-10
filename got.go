package main

import (
	"fmt"
	"log"

	"github.com/kalafut/go-todotxt"
)

func main() {
	tasks, err := todotxt.LoadFromFilename("todo.txt")
	if err != nil {
		log.Fatal(err)
	}

	t := tasks[len(tasks)-1]

	fmt.Println(t)
	fmt.Println(t.Id)

}
