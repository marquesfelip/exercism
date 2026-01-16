package lsproduct

import (
    "errors"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    // Validations
	if span < 0 {
        return 0, errors.New("span must not be negative")
    }

    if span > len(digits) {
        return 0, errors.New("span must be smaller than string length")
    }

    for i := 0; i < len(digits); i++ {
        _, err := strconv.Atoi(string(digits[i]))

        if err != nil {
            return 0, errors.New("digits input must only contain digits")
        }
    }

    var maxProduct int64 = 0

    for i := 0; i <= len(digits)-span; i++ {
        var product int64 = 1
        
        for j := 0; j < span; j++ {
            digit, _ := strconv.Atoi(string(digits[i+j]))
            product *= int64(digit)
        }

        if product > maxProduct {
            maxProduct = product
        }
    }

    return maxProduct, nil
}
