package storage

import (
	"fmt"
	"github.com/TranManhChung/csv-upload/common/util"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

const (
	CSVStoragePath  = "/Users/lap-00935/go/src/github.com/TranManhChung/csv-upload/service/storage/data"
	FileNamePattern = "upload-*.csv"
	FormKey         = "myFile"
	MaxMem          = 10 << 20
)

func upload(w http.ResponseWriter, r *http.Request) {
	defer util.RecoverFunc("upload")
	log.Printf("[Info][Upload] start")

	r.ParseMultipartForm(MaxMem)

	file, _, err := r.FormFile(FormKey)
	if err != nil {
		log.Printf("[Error][Upload] Retrieve data from file failed, detail: %e", err)
		return
	}

	go save(&file)

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	log.Printf("[Info][Upload] end")
}



func save(file *multipart.File) {
	defer util.RecoverFunc("save")

	defer (*file).Close()

	tempFile, err := ioutil.TempFile(CSVStoragePath, FileNamePattern)
	if err != nil {
		log.Printf("[Error][Upload] Create new file failed, detail: %e", err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(*file)
	if err != nil {
		log.Printf("[Error][Upload] Read data from file failed, detail: %e", err)
		return
	}

	if _, err = tempFile.Write(fileBytes); err != nil {
		log.Printf("[Error][Upload] Write data to file failed, detail: %e", err)
		return
	}

}
