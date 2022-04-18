package quering

import (
	"fmt"
	"log"
	"net/http"
)

type Service struct {
	PriceRepo IPriceRepo
}

func New() func() {
	cfg := NewDefaultConfig()
	bunDB, err := NewBunDB(cfg)
	if err != nil {
		log.Fatalf("Connect to db failed, detail: %v", err)
	}
	oRepo := NewPriceRepo(bunDB)
	service := Service{
		PriceRepo: oRepo,
	}

	http.HandleFunc("/prices", service.get)
	go http.ListenAndServe(":8081", nil)

	return func() {
		err = bunDB.Close()
		fmt.Println("Clean up quering, detail: ", err)
	}
}
