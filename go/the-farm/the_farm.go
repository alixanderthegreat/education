package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(f FodderCalculator, i int) (float64, error) {
	a, err := f.FodderAmount(i)
	if err != nil {
		return 0.0, err
	}
	r, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	p := (a / float64(i)) * float64(r)
	return p, nil
}
func ValidateInputAndDivideFood(f FodderCalculator, i int) (float64, error) {
	if i <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(f, i)
}
func ValidateNumberOfCows(i int) error {
	if i < 0 {
		return fmt.Errorf("%d cows are invalid: there are no negative cows", i)
	} else if i == 0 {
		return fmt.Errorf("%d cows are invalid: no cows don't need food", i)
	}
	return nil
}
