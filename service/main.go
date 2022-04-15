package main

import (
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/storage"
	"net/http"
)

func main() {

	fmt.Println("Server is running ...")
	storage.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}
