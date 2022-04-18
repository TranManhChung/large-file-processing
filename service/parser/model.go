package parser

import (
	"strconv"
)

type Price struct {
	Unix   float64 `bun:",pk," json:"unix"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
}

func sliceStrToPrice(in []string) (Price, error) {
	id, err := strconv.ParseFloat(in[0], 64)
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
		Unix:   id,
		Symbol: in[1],
		Open:   open,
		High:   high,
		Low:    low,
		Close:  cl,
	}, nil
}
