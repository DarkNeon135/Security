package main

import (
	"Lab1/worker"
	"log"
	"time"
)

func main() {
	start := time.Now()
	decodedResult, err := worker.DecryptThatFileSuka("chipher.txt")
	if err != nil {
		log.Fatal("Some error occurred", err)
	}
	for _, stm := range decodedResult {
		log.Print(stm)
	}
	log.Print(time.Since(start))
}
