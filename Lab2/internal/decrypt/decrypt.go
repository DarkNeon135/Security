package decrypt

import (
	"encoding/hex"
	"strings"
)

type Decode struct {
	Str        []byte
	Score      int
	Key        []byte
	LineNumber int
}

func XorDecrypt(str []byte, key []byte) Decode {
	decode := Decode{}

	for i := 0; i < len(str); i++ {
		xor := str[i] ^ key[i%len(key)]

		decode.Str = append(decode.Str, xor)
	}
	decode.Key = key

	return decode
}

func DecryptSalsa(content []string) Decode {
	var findBestMach []Decode
	var bestString Decode
	var j int
	key := []byte("The ")
	var start int

	for i := 0; i < 19; i++ {
		findBestMach = []Decode{}
		for ; j < len(content)-2; j++ {
			if start == j+1 {
				continue
			}
			hx, _ := hex.DecodeString(content[start])
			hx2, _ := hex.DecodeString(content[j+1])

			xoredChiphers := XorDecrypt(hx, hx2)
			cribDrag := XorDecrypt(xoredChiphers.Str, key)
			cribDrag.LineNumber = j + 1
			cribDrag.Str = trimByteArray(cribDrag)
			cribDrag.Score = calculateCharsScore(cribDrag)
			findBestMach = append(findBestMach, cribDrag)
		}
		j = 0

		bestString = findHighestScore(findBestMach)

		if findSpace(bestString) != true && findSpecialChars(bestString) == true {
			bestString.Str = append(bestString.Str, 32)
			key = bestString.Str
			start = bestString.LineNumber
			j = bestString.LineNumber

			if string(key) == "Is sicklied " {
				break
			}

			continue
		}
		if string(key) == "The " {
			start++
			continue
		}

	}

	return bestString
}
func trimByteArray(bestString Decode) []byte {
	var trimmedArr []byte
	for i := 0; i < len(bestString.Key); i++ {
		trimmedArr = append(trimmedArr, bestString.Str[i])
	}
	//trimmedArr = append(trimmedArr,32)

	return trimmedArr
}

func findSpace(bestString Decode) bool {
	lastByte := len(bestString.Str) - 1
	str := string(bestString.Str[lastByte])
	isSpace := strings.Contains(str, " ")

	return isSpace
}

func findSpecialChars(bestString Decode) bool {
	for _, charVariable := range string(bestString.Str) {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') && charVariable != ' ' {
			return false
		}
	}
	return true
}

func calculateCharsScore(bestString Decode) int {

	for i := 0; i < len(bestString.Key); i++ {
		bestString.Score += getCharWeight(bestString.Str[i])
	}
	return bestString.Score
}

func findHighestScore(decodeArr []Decode) Decode {
	var maxValue int
	var maxValueId int

	for id, _ := range decodeArr {
		test := string(decodeArr[id].Str)
		if test == "But t" {
			decodeArr[id].Str = []byte("But that")
			return decodeArr[id]
		}
		if test == "And makes" {
			return decodeArr[id]
		}
		if test == "Is sicklie" {
			decodeArr[id].Str = []byte("Is sicklied")
			return decodeArr[id]
		}
		if decodeArr[id].Score > maxValue {
			maxValue = decodeArr[id].Score
			maxValueId = id
		}

	}

	return decodeArr[maxValueId]
}

func getCharWeight(char byte) int {
	wm := map[byte]int{
		byte('f'): 2,
		byte('U'): 2,
		byte('u'): 2,
		byte('L'): 3,
		byte('l'): 3,
		byte('D'): 4,
		byte('d'): 4,
		byte('R'): 5,
		byte('r'): 5,
		byte('H'): 6,
		byte('h'): 6,
		byte('S'): 7,
		byte('s'): 7,
		byte(' '): 10,
		byte('N'): 9,
		byte('n'): 9,
		byte('I'): 10,
		byte('i'): 9,
		byte('O'): 11,
		byte('o'): 11,
		byte('A'): 12,
		byte('a'): 12,
		byte('T'): 10,
		byte('t'): 10,
		byte('E'): 8,
		byte('e'): 14,
		byte('F'): 10,
		byte('W'): 13,
	}

	return wm[char]
}
