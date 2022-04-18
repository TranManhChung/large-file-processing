package util

import (
	"fmt"
	"log"
)

const CSVStoragePath = "/app/"

func RecoverFunc(funcName string) {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("Recovered in %s", funcName), r)
	}
}
