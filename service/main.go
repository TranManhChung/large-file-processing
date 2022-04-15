package main

import (
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/common/util"
	"github.com/TranManhChung/large-file-processing/service/storage"
	"net/http"
)

func main() {
	defer util.RecoverFunc("save")

	fmt.Println("Server is running ...")
	storage.SetupRoutes()
	http.ListenAndServe(":8080", nil)
}
