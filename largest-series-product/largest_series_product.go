package lsproduct

import (
	"errors"
	"sort"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	if span < 0 {
		return 0, errors.New("span can not be negtive")
	}

	products := make([]int64, 0)
	for i := 0; i < len(digits)-span+1; i++ {
		product, err := calcProduct(digits[i : i+span])
		if err != nil {
			return 0, err
		}
		products = append(products, product)
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i] > products[j]
	})

	return int64(products[0]), nil
}

func calcProduct(digits string) (int64, error) {
	var product int64 = 1

	for _, d := range []byte(digits) {
		if d < 48 || d > 57 {
			return 0, errors.New("digits input must only contain digits")
		}
		product *= int64(d - 48)
	}
	return product, nil
}
