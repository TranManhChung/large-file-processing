package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const CSVStoragePath = "/Users/lap-00935/go/src/github.com/TranManhChung/large-file-processing/service/storage/data/"

func RecoverFunc(funcName string) {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("Recovered in %s", funcName), r)
	}
}

func SplitFile(filePath string, maxLines int) error {
	defer RecoverFunc("SplitFile")

	fmt.Println("[SplitFile] Start")
	var f *os.File
	var data []string

	source, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer source.Close()

	csvReader := csv.NewReader(source)

	for err == nil {
		if f, err = os.Create(CSVStoragePath + fmt.Sprintf("%v.csv", time.Now().UnixNano())); err != nil {
			break
		}
		w := csv.NewWriter(f)

		for counter := 0; counter < maxLines; counter++ {
			data, err = csvReader.Read()
			if err == io.EOF {
				if err = w.Write(data); err != nil {
					break
				}
				err = io.EOF
				break
			}
			if err != nil {
				break
			}
			if err = w.Write(data); err != nil {
				break
			}
		}
		w.Flush()
		f.Close()
	}

	// delete origin file because it isn't necessary anymore
	if err == io.EOF {
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
		log.Printf("[SplitFile] Clean up file %v", filePath)
	}

	fmt.Println("[SplitFile] End")
	return err
}
