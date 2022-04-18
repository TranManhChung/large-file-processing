package util

import (
	"fmt"
	"log"
)

const CSVStoragePath = "/Users/lap-00935/go/src/github.com/TranManhChung/large-file-processing/service/storage/data/"

func RecoverFunc(funcName string) {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("Recovered in %s", funcName), r)
	}
}
