package errorexercise

import (
	"errors"
	"fmt"
)

// validateLength return false when string length less than 8
// Please change return type to error with message "too short string"
func validateLength(s string) error {
	if len([]rune(s)) < 8 {
		return errors.New("too short string")
	}
	return nil
}

var ageError = errors.New("your age is negative!")

type ErrTooYoung int

func (e ErrTooYoung) Error() string {
	return fmt.Sprintf("%d is too young", e)
}

// in case of too young age please create a new type ErrTooYoung int` with method `Error() string`
// example error message: "17 is too young"
func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return ErrTooYoung(n)
	}
	return nil
}
