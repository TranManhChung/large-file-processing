package storage

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/TranManhChung/large-file-processing/pkg/util"
	"github.com/TranManhChung/large-file-processing/queue"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
)

const (
	FormKey = "myFile"
	MaxMem  = 10 << 20
	MaxLine = 50
)

type UploadResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func (s Service) upload(w http.ResponseWriter, r *http.Request) {
	defer util.RecoverFunc("upload")
	log.Printf("[Info][Upload] Start")
	defer log.Printf("[Info][Upload] End")

	w.Header().Set("Content-Type", "application/json")

	r.ParseMultipartForm(MaxMem)

	file, handler, err := r.FormFile(FormKey)
	if err != nil {
		log.Printf("[Error][Upload] Retrieve file failed, detail: %v", err)
		json.NewEncoder(w).Encode(UploadResponse{
			Status:  "failed",
			Message: "handle file failed",
		})
		return
	}
	if handler.Filename[len(handler.Filename)-4:] != ".csv" {
		json.NewEncoder(w).Encode(UploadResponse{
			Status:  "failed",
			Message: "invalid format",
		})
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

	s.WorkerPool.AddTask(func() error {
		return SplitFile(newLocation, MaxLine)
	})

	json.NewEncoder(w).Encode(UploadResponse{
		Status: "success",
	})
}

func SplitFile(filePath string, maxLines int) error {
	defer util.RecoverFunc("SplitFile")

	log.Println("[SplitFile] Start")
	defer log.Println("[SplitFile] End")

	var dest *os.File
	var data []string

	source, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer source.Close()

	csvReader := csv.NewReader(source)

	for err == nil {
		destPath := util.CSVStoragePath + fmt.Sprintf("%v.csv", time.Now().UnixNano())
		if dest, err = os.Create(destPath); err != nil {
			break
		}
		writer := csv.NewWriter(dest)

		for counter := 0; counter < maxLines; counter++ {
			data, err = csvReader.Read()
			if err == io.EOF {
				if err = writer.Write(data); err != nil {
					break
				}
				err = io.EOF
				break
			}
			if err != nil {
				break
			}
			if err = writer.Write(data); err != nil {
				break
			}

		}
		writer.Flush()
		dest.Close()
		queue.GetQueue().Publish(destPath)
	}

	// delete origin file because it isn't necessary anymore
	if err == io.EOF {
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
		log.Printf("[SplitFile] Clean up file %v", filePath)
	}

	return err
}
