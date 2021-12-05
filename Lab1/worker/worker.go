package worker

import (
	"Lab1/internal/decrypt"
	"Lab1/internal/file_processing"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"strings"
)

func DecryptThatFileSuka(fileName string) ([]decrypt.Decode, error) {
	content, err := file_processing.ReadFromFile("internal/" + fileName)
	var decodedResult []decrypt.Decode
	if err != nil {
		log.Fatal(err)
	}

	decryptedStruct, err := doSecondTask(content)
	if err != nil {
		return nil, err
	}

	decodedResult = append(decodedResult, decryptedStruct)

	decryptedStruct, err = doThirdTask(content)
	if err != nil {
		return nil, err
	}

	decodedResult = append(decodedResult, decryptedStruct)

	return decodedResult, nil
}

func doSecondTask(content []string) (decrypt.Decode, error) {
	hx, _ := hex.DecodeString(content[0])
	//if err != nil {
	//	return decrypt.Decode{},errors.New("[worker] -> Error in singleByte xor "+ err.Error())
	//}
	possibleChars := decrypt.ABC + strings.ToLower(decrypt.ABC) + decrypt.Numbers
	return decrypt.MakeBrutForce(hx, 1, possibleChars), nil
}

func doThirdTask(content []string) (decrypt.Decode, error) {
	decodedBase, err := base64.URLEncoding.DecodeString(content[1])
	if err != nil {
		return decrypt.Decode{}, errors.New("[worker] -> Error in thirdTask " + err.Error())
	}
	var wtd []byte
	for i := 0; i < 15; i++ {
		wtd = append(wtd, decodedBase[i])
	}
	possibleChars := decrypt.ABC + strings.ToLower(decrypt.ABC) + decrypt.Numbers

	decryptedPart := decrypt.MakeBrutForce(wtd, 3, possibleChars)
	return decrypt.XorDecrypt(decodedBase, decryptedPart.Key), nil
}
