package parser

import (
	"strconv"
)

type Price struct { // remember to add tag for field
	Unix   int64 `bun:",pk,"`
	Symbol string
	Open   float64
	High   float64
	Low    float64
	Close  float64
}

func sliceStrToPrice(in []string) (Price, error) {
	id, err := strconv.ParseInt(in[0], 10, 64)
	if err != nil {
		return Price{}, err
	}
	open, err := strconv.ParseFloat(in[2], 64)
	if err != nil {
		return Price{}, err
	}
	high, err := strconv.ParseFloat(in[3], 64)
	if err != nil {
		return Price{}, err
	}
	low, err := strconv.ParseFloat(in[4], 64)
	if err != nil {
		return Price{}, err
	}
	cl, err := strconv.ParseFloat(in[5], 64)
	if err != nil {
		return Price{}, err
	}

	return Price{
		Unix:     id,
		Symbol: in[1],
		Open:   open,
		High:   high,
		Low:    low,
		Close:  cl,
	}, nil
}
