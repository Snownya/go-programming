package main

import (
	"fmt"
	"github.com/snownya/go-programming/cmd"
	"log"
)

func main() {
	fmt.Println("working....")
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}

}
