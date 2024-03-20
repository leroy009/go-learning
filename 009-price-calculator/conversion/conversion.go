package conversion

import (
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, value := range strings {
		floatPrice, err :=strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, err
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}