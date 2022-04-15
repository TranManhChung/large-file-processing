package parser

type Parser struct {
	TaskQueue      chan string
	NumberOfWorker int
}

