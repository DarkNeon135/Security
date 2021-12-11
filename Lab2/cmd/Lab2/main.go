package main

import (
	"Lab2/worker"
	"fmt"
)

func main() {
	decyptedPoem, err := worker.MakeDecryption("chipher.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range decyptedPoem {
		fmt.Println(line)
	}

}
