package hdr

import (
	"bytes"

	"github.com/HdrHistogram/hdrhistogram-go"
)

// Plot creates a percentile distribution plot from a slice of int64 values.
// It returns the percentile distribution as a formatted string.
func Plot(data []int64) (string, error) {
	// Create a histogram with a resolution of 1 microsecond
	// The maximum value can be set according to your needs, here it's set to 60 million microseconds (60 seconds)
	histogram := hdrhistogram.New(1, 60*1000000, 5)

	// Add the data to the histogram
	for _, value := range data {
		if value < 0 {
			continue
		}

		err := histogram.RecordValue(value)
		if err != nil {
			return "", err
		}
	}

	// Create a buffer to capture the output of the PercentilesPrint function
	var buf bytes.Buffer

	// Calculate and print the percentiles to the buffer
	_, err := histogram.PercentilesPrint(&buf, 1, 1.0)
	if err != nil {
		return "", err
	}

	// Get the output as a string
	return buf.String(), nil
}
