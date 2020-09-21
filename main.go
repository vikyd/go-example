package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("hi")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(uuid.New())
}
