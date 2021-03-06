package decrypt

import "math"

func MakeBrutForce(encryptedStr []byte, keyLength int, possibleChars string) Decode {

	keyIndexes := []int{0}
	newDigitLoop := len(possibleChars)
	iterationCount := 0
	var decodedArray []Decode

	for len(keyIndexes) <= keyLength {
		iterationCount++
		possibleKey := ""
		for _, charIdx := range keyIndexes {
			possibleKey += string(possibleChars[charIdx])
		}

		decodedStruct := XorDecrypt(encryptedStr, []byte(possibleKey))
		decodedArray = append(decodedArray, decodedStruct)

		if iterationCount == newDigitLoop {
			for j := 0; j < len(keyIndexes); j++ {
				keyIndexes[j] = 0
			}
			keyIndexes = append([]int{0}, keyIndexes...)
			newDigitLoop = int(math.Pow(float64(len(possibleChars)), float64(len(keyIndexes))))
			iterationCount = 0
			continue
		}

		for i := 0; i < len(keyIndexes); i++ {
			keyIndexes[i] = iterationCount / int(math.Pow(float64(len(possibleChars)), float64(len(keyIndexes)-1-i))) % len(possibleChars)
		}
	}
	result := findHighestScore(decodedArray)

	return result
}
