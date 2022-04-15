package storage

import "net/http"

func SetupRoutes() {
	http.HandleFunc("/upload", upload)
}
