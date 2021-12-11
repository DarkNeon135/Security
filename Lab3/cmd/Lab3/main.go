package main

import (
	"Lab3/internal/casino"
	"fmt"
)

func main() {
	var cas casino.Account
	err := cas.MakeAccount(5555555510)
	if err != nil {
		fmt.Println(err)
	}
}
