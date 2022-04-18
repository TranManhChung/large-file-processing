package parser

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/common/util"
	"github.com/TranManhChung/large-file-processing/service/common/worker"
	"github.com/TranManhChung/large-file-processing/service/queue"
	"io"
	"log"
	"os"
	"reflect"
)

type Service struct {
	WorkerPool worker.Pool
	PriceRepo  IPriceRepo
}

func New() func() {
	cfg := NewDefaultConfig()
	ctx := context.Background()
	bunDB, err := NewBunDB(cfg)
	if err != nil {
		log.Fatalf("Connect to db failed, detail: %v", err)
	}
	oRepo := NewPriceRepo(bunDB)
	service := Service{
		WorkerPool: worker.New(cfg.Worker.MaxWorkerPoolTask, cfg.Worker.MaxWorkers, cfg.Worker.WorkerName),
		PriceRepo:  oRepo,
	}
	service.WorkerPool.Run()

	go func(ctx context.Context, sv Service) {
		sv.consumeMsg(ctx)
	}(ctx, service)
	return func() {
		err = bunDB.Close()
		fmt.Println("Clean up parser, detail: ", err)
	}
}

func (s *Service) consumeMsg(ctx context.Context) {
	util.RecoverFunc("consumeMsg")
	for {
		msg := queue.GetQueue().Subscribe()
		s.parse(ctx, msg)
	}
}

const MaxItem = 10

func (s *Service) parse(ctx context.Context, msg string) {
	numField := reflect.TypeOf(Price{}).NumField()
	source, err := os.Open(msg)
	if err != nil {
	}

	csvReader := csv.NewReader(source)
	numItems := 0
	prices := make(map[string][]Price)
	for {
		data, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("[Parser] Read file failed, file: %v, detail: %e", msg, err)
			break
		}
		if reflect.DeepEqual(data, []string{"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE"}) {
			fmt.Println("fasndfkasfdkasjfkasdjlfkdsjflaj")
			continue
		}
		if numField != len(data) {
			log.Printf("[Parser] Not enough fields, file: %v, data: %v, numfield: %v, actual: %v", msg, data, numField, len(data))
			break
		}
		o, err := sliceStrToPrice(data)
		if err != nil {
			log.Printf("[Parser] Parse data failed, file: %v, detail: %e", msg, err)
			break
		}
		prices[o.Symbol] = append(prices[o.Symbol], o)
		numItems++
		if numItems == MaxItem {
			for _, v := range prices {
				if err = s.PriceRepo.Adds(ctx, v...); err != nil {
					log.Fatalf("[Parser] Store data failed, file: %v, detail: %e", msg, err)
				}
			}
			prices = make(map[string][]Price)
		}
	}

	source.Close()
}
