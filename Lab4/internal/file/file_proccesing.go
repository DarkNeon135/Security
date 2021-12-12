package file

import (
	"io/ioutil"
	"strings"
)

func ReadFromFile(fileName string) ([]string, error) {
	configFile, err := ioutil.ReadFile("internal/resources/" + fileName)
	if err != nil {
		return nil, err
	}

	configLines := strings.Split(string(configFile), "\n")

	return configLines, nil
}
