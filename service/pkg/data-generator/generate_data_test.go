package data_generator

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateData(t *testing.T) {

	dest, err := os.Create("/Users/lap-00935/Desktop/large-gen.csv")
	assert.Nil(t, err)
	defer dest.Close()

	writer := csv.NewWriter(dest)

	err = writer.Write([]string{
		"UNIX", "SYMBOL", "OPEN", "HIGH", "LOW", "CLOSE",
	})
	assert.Nil(t, err)

	for i := 0; i < 100000000; i++ {
		err = writer.Write([]string{
			fmt.Sprintf("%v", i),
			"btcusdt", "100", "99", "98", "97",
		})
		assert.Nil(t, err)
	}
	
	writer.Flush()
}
