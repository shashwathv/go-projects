package validation

import "errors"

var (
	ErrMissingNumber1 = errors.New("number 1 is required")
	ErrMissingNumber2 = errors.New("number 2 is required")
	ErrDivisionByZero = errors.New("divisor cannot be zero")
	ErrEmptySlice     = errors.New("numbers array cannot be empty")
)

func ValidateTwoNumbers(a *int, b *int) (int, int, error) {
	if a == nil {
		return 0, 0, ErrMissingNumber1
	}
	if b == nil {
		return 0, 0, ErrMissingNumber2
	}
	return *a, *b, nil
}

func ValidateDivide(dividend *int, divisor *int) (int, int, error) {
	if dividend == nil {
		return 0, 0, errors.New("divident cannot be zero")
	}
	if divisor == nil {
		return 0, 0, errors.New("divisor cannot be zero")
	}
	if *divisor == 0 {
		return 0, 0, ErrDivisionByZero
	}
	return *dividend, *divisor, nil
}

func ValidateSum(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return nil, ErrEmptySlice
	}
	return numbers, nil
}
