package main

import (
	"Lab1/worker"
	"fmt"
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
		fmt.Println(fmt.Sprintf("Decypted string -> %s\nKey -> %s", stm.Str, stm.Key))
		fmt.Println("-----------------------------------------------------------------")
	}
	log.Print(time.Since(start))

}
