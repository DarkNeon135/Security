package file_processing

import (
	"io/ioutil"
	"strings"
)

func ReadFromFile(fileName string) ([]string, error) {
	configFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	configLines := strings.Split(string(configFile), "\n")

	return configLines, nil
}
