package worker

import (
	"Lab1/internal/decrypt"
	"encoding/hex"
	"fmt"
	"os"
)

func MakeBrutForce(fileName string) {
	content, err := os.ReadFile("internal/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	hx, err := hex.DecodeString(string(content))
	if err != nil {

	}
	result := decrypt.MakeDecryption(hx)
	fmt.Println(result)

}
