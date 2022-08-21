package errorexercise

import (
	"errors"
	"testing"
)

func TestValidateLength(t *testing.T) {
	given := "1234567"
	want := errors.New("too short string")

	err := validateLength(given)
	if errors.Is(err, want) {
		t.Errorf("validateLength(\"\", 0) = %v, want nil", err)
	}
}

func TestValidateAgeNegative(t *testing.T) {
	given := -1
	want := errors.New("your age is negative!")

	err := validateAge(given)
	if errors.Is(err, want) {
		t.Errorf("validateAge(%d) = %v, want nil", given, err)
	}
}

func TestValidateAgeTooYoung(t *testing.T) {
	given := 17
	want := ErrTooYoung(17)

	err := validateAge(given)

	if err != want {
		t.Errorf("validateAge(%d) = %s(%T), want %s(%T)", given, err, err, want, want)
	}
}
