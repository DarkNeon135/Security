package worker

import (
	"Lab4/internal/file"
	"Lab4/internal/generator"
	"Lab4/internal/hash"
	"log"
)

func StartHashing() {
	passwordArr1 := generator.GeneratePassword()

	file.WriteToTxt(passwordArr1, "strong")

	bcrypt, err := hash.MakeBCryptHash(passwordArr1)
	if err != nil {
		log.Fatal("Error when hashing password via bcrypt", err)
	}
	file.WriteToCsv(bcrypt, "strong")

	passwordArr2 := generator.GeneratePassword()

	file.WriteToTxt(passwordArr2, "weak")

	md5 := hash.MakeMd5Hash(passwordArr2)
	file.WriteToCsv(md5, "weak")
}
