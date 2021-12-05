package worker

import (
	"Lab2/internal/decrypt"
	"Lab2/internal/file"
	"encoding/hex"
	"errors"
	"fmt"
)

func MakeDecryption(fileName string) ([]string, error) {
	content, err := file.ReadFromFile("internal/" + fileName)
	if err != nil {
		return nil, errors.New("[worker] -> Error in reading file " + err.Error())
	}
	test := decrypt.DecryptSalsa(content)
	fmt.Printf("Try to Google it! -> %s\n", string(test.Str))
	//Google this lines
	var result []string
	key := []byte("Is sicklied o'er with the pale cast of thought, ")

	for i := 0; i < len(content)-1; i++ {
		hx, _ := hex.DecodeString(content[test.LineNumber])
		hx2, _ := hex.DecodeString(content[i])
		if len(hx) > len(hx2) {
			var newHx []byte
			for j := 0; j < len(hx2); j++ {
				newHx = append(newHx, hx[j])
			}
			hx = newHx
		}
		xoredHex := decrypt.XorDecrypt(hx, hx2)
		xored := decrypt.XorDecrypt(xoredHex.Str, key)
		result = append(result, string(xored.Str))
	}

	return result, nil
}
