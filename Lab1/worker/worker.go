package worker

import (
	"Lab1/internal/decrypt"
	"fmt"
	"os"
)

func MakeBrutForce(fileName string) {
	content, err := os.ReadFile("internal/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	result := decrypt.MakeDecryption(content)
	fmt.Println(result)
	//bt := []byte("Matvey_Sosi Penis")
	//var result string
	//for i:=0;i<len(bt);i++ {
	//	result += string(bt[i] ^ 254)
	//}
	//test := []byte(result)
	//
	//var encrypt []string
	//for i:=0;i<256;i++{
	//	var xuy string
	//	for j:=0;j<len(test);j++{
	//		xor := test[j] ^ byte(i)
	//		xuy += string(xor)
	//	}
	//	encrypt = append(encrypt,xuy)
	//}
	//fmt.Println(encrypt)
}
