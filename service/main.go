package main

import (
	"fmt"
	"github.com/TranManhChung/large-file-processing/service/common/util"
	"github.com/TranManhChung/large-file-processing/service/parser"
	"github.com/TranManhChung/large-file-processing/service/storage"
	"os"
	"os/signal"
)

func main() {
	defer util.RecoverFunc("save")

	fmt.Println("Server is running ...")

	var cleanups []func()
	cleanups = append(cleanups, storage.New())
	cleanups = append(cleanups, parser.New())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		fmt.Println("terminate app", sig.String())
		for _, v := range cleanups {
			v()
		}
		return
	}
}
