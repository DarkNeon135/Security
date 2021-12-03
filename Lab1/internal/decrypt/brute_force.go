package decrypt

import "math"

func MakeBrutForce(encryptedStr []byte, keyLength int) Decode {

	possibleCharValues := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyIndexes := []int{0}
	newDigitLoop := len(possibleCharValues)
	iterationCount := 0
	var decodedArray []Decode

	for len(keyIndexes) <= keyLength {
		iterationCount++
		possibleKey := ""
		for _, charIdx := range keyIndexes {
			possibleKey += string(possibleCharValues[charIdx])
		}

		decodedStruct := XorDecrypt(encryptedStr, possibleKey)
		decodedArray = append(decodedArray, decodedStruct)
		//
		// checking for key suitability... blaaaa decoder... blaaaa xor... blaaaa...
		//

		if iterationCount == newDigitLoop {
			for j := 0; j < len(keyIndexes); j++ {
				keyIndexes[j] = 0
			}
			keyIndexes = append([]int{0}, keyIndexes...)
			newDigitLoop = int(math.Pow(float64(len(possibleCharValues)), float64(len(keyIndexes))))
			iterationCount = 0
			continue
		}

		for i := 0; i < len(keyIndexes); i++ {
			keyIndexes[i] = iterationCount / int(math.Pow(float64(len(possibleCharValues)), float64(len(keyIndexes)-1-i))) % len(possibleCharValues)
		}
	}
	result := findHighestScore(decodedArray)

	return result
}
