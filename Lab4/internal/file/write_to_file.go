package file

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

func WriteToCsv(passwordArr [][]string, filename string) {
	csvFile, err := os.Create(filename + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range passwordArr {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()
	csvFile.Close()
}

func WriteToTxt(passwordArr []string, filename string) {
	file, err := os.OpenFile(filename+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range passwordArr {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
