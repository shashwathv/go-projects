package service

import "errors"

var ErrDivisionByZero = errors.New("division by zero")
var ErrEmptySlice = errors.New("slice cannot be empty")

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

func Subtract[T int | float32 | float64](a T, b T) T {
	return a - b
}

func Multiply[T int | float32 | float64](a T, b T) T {
	return a * b
}

func Divide[T int | float32 | float64](a T, b T) (T, error) {
	if b == 0 {
		var zero T
		return zero, ErrDivisionByZero
	}
	return a / b, nil
}

func Sum(slice []int) (int, error) {
	result := 0
	if len(slice) == 0 {
		return 0, ErrEmptySlice
	}
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result, nil
}
