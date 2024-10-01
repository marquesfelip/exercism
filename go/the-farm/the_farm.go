package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	totalFodder := fodderAmount * fatteningFactor
	foodPerCow := totalFodder / float64(cows)

	return foodPerCow, nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}

	return nil
}

type InvalidCowsError struct {
	cows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}
