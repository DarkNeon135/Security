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
	//hx, _ := hex.DecodeString("3a0a9cab782b4f8603eac28aadde1151005fd46a859df21d12c38eaa858596bf2548000e883d72117466c5c3a580f66b")
	//hx2, _ := hex.DecodeString("3a0adee4783a538403b9c29eaac958550242d3778ed9a61918959bf4ca849afa68450f5edc6e311a7f7ed1d7ec")
	//test := decrypt.XorDecrypt(hx,hx2)
	// fmt.Println( []byte(test.Str))
	//brbrb := decrypt.XorDecrypt([]byte(test.Str),[]byte("The "))
	//fmt.Println(brbrb)
	//br := decrypt.XorDecrypt([]byte("For who would bear the whips and scorns of time,"),"Th'oppressor's wrong, the proud man's contumely,\n")
	//fmt.Println([]byte(br.Str))
}
