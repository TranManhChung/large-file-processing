package util

import (
	"fmt"
	"log"
)

//const CSVStoragePath = "/Users/lap-00935/go/src/github.com/TranManhChung/large-file-processing/storage/data/"
const CSVStoragePath = "/app/"

func RecoverFunc(funcName string) {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("Recovered in %s", funcName), r)
	}
}
