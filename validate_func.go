package validator

import (
	"regexp"

	"golang.org/x/exp/constraints"
)

// MaxLength returns an error if the length of value is greater than length.
func MaxLength[T LengthAble[any]](value T, length int, err error) error {
	if len(value) > length {
		return err
	}
	return nil
}

// MinLength returns an error if the length of value is less than length.
func MinLength[T LengthAble[any]](value T, length int, err error) error {
	if len(value) < length {
		return err
	}
	return nil
}

// StrMaxLength returns an error if the length of value is greater than length.
func StrMaxLength(value string, length int, err error) error {
	return MaxLength(value, length, err)
}

// StrMinLength returns an error if the length of value is less than length.
func StrMinLength(value string, length int, err error) error {
	return MinLength(value, length, err)
}

// Gt returns an error if b is greater than a.
func Gt[T constraints.Ordered](a, b T, err error) error {
	if a < b {
		return err
	}
	return nil
}

// Lt returns an error if b is less than a.
func Lt[T constraints.Ordered](a, b T, err error) error {
	if a > b {
		return err
	}
	return nil
}

// Gte returns an error if b is greater than or equal to a.
func Gte[T constraints.Ordered](a, b T, err error) error {
	if a <= b {
		return err
	}
	return nil
}

// Lte returns an error if b is less than or equal to a.
func Lte[T constraints.Ordered](a, b T, err error) error {
	if a >= b {
		return err
	}
	return nil
}

// Eq returns an error if a is not equal to b.
func Eq(a, b any, err error) error {
	if a != b {
		return err
	}
	return nil
}

// NotEq returns an error if a is equal to b.
func NotEq(a, b any, err error) error {
	if a == b {
		return err
	}
	return nil
}

// Contains returns an error if value is not in values.
func Contains[T comparable](value T, values []T, err error) error {
	for _, v := range values {
		if v == value {
			return nil
		}
	}
	return err
}

// NotContains returns an error if value is in values.
func NotContains[T comparable](value T, values []T, err error) error {
	for _, v := range values {
		if v == value {
			return err
		}
	}
	return nil
}

// Regex returns an error if value does not match the regular expression.
func Regex(value string, pattern string, err error) error {
	exp, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	return WithRegex(value, exp, err)
}

// WithRegex returns an error if value does not match the regular expression.
func WithRegex(value string, exp *regexp.Regexp, err error) error {
	if !exp.MatchString(value) {
		return err
	}
	return nil
}

// Email returns an error if value is not a valid email address.
func Email(value string, err error) error {
	return WithRegex(value, emailRegex, err)
}
