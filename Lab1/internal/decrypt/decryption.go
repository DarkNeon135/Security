package decrypt

func MakeDecryption(str []byte) []string {

	var decryptedMsg []string

	var result string
	for j := 0; j < len(str); j++ {
		xor := rune(str[j] ^ byte(37))
		result += string(xor)
	}
	decryptedMsg = append(decryptedMsg, result)

	return decryptedMsg
}
