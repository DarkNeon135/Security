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
	//content ,_ := file.ReadFromFile("internal/"+"chipher.txt")
	//hx, _ := hex.DecodeString(content[15])
	//hx2, _ := hex.DecodeString(content[1])
	////hx := []byte("3a0adee4783a538403b9c29eaac958550242d3778ed9a61918959bf4ca849afa68450f5edc6e311a7f7ed1d7ec")
	////hx2 := []byte(content[4])
	//test := decrypt.XorDecrypt(hx,hx2)
	//hxKey := []byte("Is sicklied o'er with the pale cast of thought,")
	//brbrb := decrypt.XorDecrypt(test.Str,hxKey)
	//fmt.Println(string(brbrb.Str))

	//br := decrypt.XorDecrypt([]byte("For who would bear the whips and scorns of time,"),[]byte("Th'oppressor's wrong, the proud man's contumely,"))

	//fmt.Println(string(brbrb.Str))
}
