package storage

import (
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/common/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

const (
	FormKey        = "myFile"
	MaxMem         = 10 << 20
)

func (s Service) upload(w http.ResponseWriter, r *http.Request) {
	defer util.RecoverFunc("upload")
	log.Printf("[Info][Upload] Start")

	r.ParseMultipartForm(MaxMem)

	file, handler, err := r.FormFile(FormKey)
	if err != nil {
		log.Printf("[Error][Upload] Retrieve file failed, detail: %v", err)
		return
	}
	defer file.Close()

	newLocation := ""

	if tempFile := reflect.ValueOf(handler).Elem().FieldByName("tmpfile"); tempFile.String() != "" {
		oldLocation := tempFile.String()
		newLocation = util.CSVStoragePath + handler.Filename

		if err := os.Rename(oldLocation, newLocation); err != nil {
			log.Printf("[Error][Upload] Move file to new location failed, detail: %v", err)
			return
		}
	} else {
		tempFile, err := ioutil.TempFile(util.CSVStoragePath, "origin-*.csv")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		newLocation = tempFile.Name()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
	}

	s.WorkerPool.AddTask(newLocation)

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	log.Printf("[Info][Upload] End")
}
