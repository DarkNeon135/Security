package hash

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

const saltSize = 8

func genSalt() []byte {

	abc := "aeioubcdfghjklmnpqrstvwxyz"
	var salt = make([]byte, saltSize)
	for i := 0; i <= saltSize-1; i++ {

		salt[i] = abc[rand.Intn(len(abc)-1)]

	}
	return salt
}

func MakeBCryptHash(passwordArr []string) ([][]string, error) {
	var bcryptHashes [][]string

	for i := 0; i < len(passwordArr); i++ {

		salt := genSalt()

		saltPass := string(salt) + passwordArr[i]
		saltHex := hex.EncodeToString([]byte(saltPass))
		hash, err := bcrypt.GenerateFromPassword([]byte(saltHex), 1)
		if err != nil {
			return nil, err
		}

		aray4array := []string{string(hash)}
		bcryptHashes = append(bcryptHashes, aray4array)

	}
	return bcryptHashes, nil
}

func MakeMd5Hash(passwordArr []string) [][]string {
	var md5Hash [][]string

	for i := 0; i < len(passwordArr); i++ {
		hash := md5.Sum([]byte(passwordArr[i]))
		array4array := []string{hex.EncodeToString(hash[:])}
		md5Hash = append(md5Hash, array4array)
	}
	return md5Hash
}
