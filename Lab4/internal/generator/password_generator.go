package generator

import (
	"Lab4/internal/file"
	"math/rand"
	"strconv"
	"time"
)

func GeneratePassword() []string {
	names, _ := file.ReadFromFile("names.txt")
	nouns, _ := file.ReadFromFile("nouns.txt")
	var passwordArr []string

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100000; i++ {

		var password string
		algorithm := int(rand.Uint32())
		algorithmId := algorithm % 100

		if algorithmId <= 30 {
			password = nouns[algorithm%len(nouns)]
		} else if algorithmId > 30 && algorithmId <= 50 {
			password = generateNounsNumbers(nouns, algorithm)
		} else if algorithmId > 50 && algorithmId <= 70 {
			password = GenerateNameBirth(names, algorithm)
		} else if algorithmId > 70 && algorithmId <= 80 {
			password = GenerateRandomNumbers(6, algorithm)
		} else if algorithmId > 80 && algorithmId <= 95 {
			password = names[algorithm%len(names)]
		} else {
			password = GenerateRandomPass(algorithm%15, algorithm%10)
		}

		passwordArr = append(passwordArr, password)
	}

	return passwordArr
}

func GenerateBirthYear(random int) string {
	min := 1970
	max := 2005
	randomYear := strconv.Itoa(random%(max-min) + min)

	return randomYear
}
func GenerateRandomNumbers(numberSize int, random int) string {
	digits := "0123456789"

	var midPart = make([]byte, numberSize)

	if numberSize == 0 {
		return "12345678"
	}
	if random%1 == 0 {
		for k, _ := range midPart {
			midPart[k] = digits[random%len(digits)]
		}
	} else {
		for k, _ := range midPart {
			midPart[k] = digits[rand.Intn(len(digits))]
		}
	}

	return string(midPart)
}

func GenerateNameBirth(names []string, random int) string {
	password := names[random%len(names)] + GenerateBirthYear(random)

	return password
}

func generateNounsNumbers(nouns []string, random int) string {
	password := nouns[random%len(nouns)] + GenerateRandomNumbers(5, random)
	return password
}
func GenerateRandomPass(alphabetSize, numberSize int) string {
	vowels := "aeiou"
	consonants := "bcdfghjklmnpqrstvwxyz"
	digits := "0123456789"

	prefixSize := alphabetSize / 2
	if alphabetSize%2 != 0 {
		prefixSize = int(alphabetSize/2) + 1
	}
	suffixSize := alphabetSize - prefixSize

	var prefixPart = make([]byte, prefixSize)

	for i := 0; i <= prefixSize-1; i++ {
		if i%2 == 0 {
			// use consonants
			prefixPart[i] = consonants[rand.Intn(len(consonants)-1)]
		} else {
			// use vowels
			prefixPart[i] = vowels[rand.Intn(len(vowels)-1)]
		}
	}

	var midPart = make([]byte, numberSize)

	// use digits
	for k, _ := range midPart {
		midPart[k] = digits[rand.Intn(len(digits))]
	}

	var suffixPart = make([]byte, suffixSize)

	for i := 0; i <= suffixSize-1; i++ {
		if i%2 == 0 {
			// use consonants
			suffixPart[i] = consonants[rand.Intn(len(consonants)-1)]
		} else {
			// use vowels
			suffixPart[i] = vowels[rand.Intn(len(vowels)-1)]
		}
	}

	return string(prefixPart) + string(midPart) + string(suffixPart)
}
