package worker

import (
	"Lab2/internal/file"
	"errors"
	"fmt"
)

func DecryptSalsa(fileName string) (string, error) {
	content, err := file.ReadFromFile("internal/" + fileName)
	if err != nil {
		return "", errors.New("[worker] -> Error in reading file " + err.Error())
	}

	fmt.Println(content)

	return "", nil
}
