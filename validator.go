package validator

import (
	"regexp"
)

// Validator defines a Decorator interface.
type Validator interface {
	Validate() error
}

// Group is a group of validators.
type Group []Validator

// Validate implements the Validator interface.
func (g Group) Validate() error {
	for _, v := range g {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Func is a function that implements the Validator interface.
type Func func() error

// Validate implements the Validator interface.
func (f Func) Validate() error {
	return f()
}

// Decorator wraps a group of validators.
// It provides a fluent interface for adding validators.
type Decorator struct {
	group []Validator
}

// Validate implements the Validator interface.
func (v *Decorator) Validate() error {
	return Group(v.group).Validate()
}

// Add adds a validator to the group.
func (v *Decorator) Add(validator Validator) *Decorator {
	v.group = append(v.group, validator)
	return v
}

// ValidateFunc adds a function to the group.
func (v *Decorator) ValidateFunc(f func() error) *Decorator {
	return v.Add(Func(f))
}

// StrMaxLength returns an error if the length of text is greater than length.
func (v *Decorator) StrMaxLength(text string, length int, err error) *Decorator {
	return v.ValidateFunc(func() error { return StrMaxLength(text, length, err) })
}

// StrMinLength returns an error if the length of text is less than length.
func (v *Decorator) StrMinLength(text string, length int, err error) *Decorator {
	return v.ValidateFunc(func() error { return StrMinLength(text, length, err) })
}

// Eq returns an error if a is not equal to b.
func (v *Decorator) Eq(a, b any, err error) *Decorator {
	return v.ValidateFunc(func() error { return Eq(a, b, err) })
}

// NotEq returns an error if a is equal to b.
func (v *Decorator) NotEq(a, b any, err error) *Decorator {
	return v.ValidateFunc(func() error { return NotEq(a, b, err) })
}

// GtInt returns an error if b is greater than a.
func (v *Decorator) GtInt(a, b int, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt returns an error if b is less than a.
func (v *Decorator) LtInt(a, b int, err error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, err) })
}

// GtInt8 returns an error if a is greater than b.
func (v *Decorator) GtInt8(a, b int8, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt8 returns an error if a is less than b.
func (v *Decorator) LtInt8(a, b int8, err error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, err) })
}

// GtInt16 returns an error if a is greater than b.
func (v *Decorator) GtInt16(a, b int16, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt16 returns an error if a is less than b.
func (v *Decorator) LtInt16(a, b int16, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtInt32 returns an error if a is greater than b.
func (v *Decorator) GtInt32(a, b int32, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt32 returns an error if a is less than b.
func (v *Decorator) LtInt32(a, b int32, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtInt64 returns an error if a is greater than b.
func (v *Decorator) GtInt64(a, b int64, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt64 returns an error if a is less than b.
func (v *Decorator) LtInt64(a, b int64, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint returns an error if a is greater than b.
func (v *Decorator) GtUint(a, b uint, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt8 returns an error if a is less than b.
func (v *Decorator) LtUInt8(a, b uint8, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint8 returns an error if a is greater than b.
func (v *Decorator) GtUint8(a, b uint8, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// GtUint16 returns an error if a is greater than b.
func (v *Decorator) GtUint16(a, b uint16, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt16 returns an error if a is less than b.
func (v *Decorator) LtUInt16(a, b uint16, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint32 returns an error if a is greater than b.
func (v *Decorator) GtUint32(a, b uint32, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt32 returns an error if a is less than b.
func (v *Decorator) LtUInt32(a, b uint32, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint64 returns an error if a is greater than b.
func (v *Decorator) GtUint64(a, b uint64, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt64 returns an error if a is less than b.
func (v *Decorator) LtUInt64(a, b uint64, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtFloat32 returns an error if a is greater than b.
func (v *Decorator) GtFloat32(a, b float32, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtFloat32 returns an error if a is less than b.
func (v *Decorator) LtFloat32(a, b float32, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtFloat64 returns an error if a is greater than b.
func (v *Decorator) GtFloat64(a, b float64, err error) *Decorator {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtFloat64 returns an error if a is less than b.
func (v *Decorator) LtFloat64(a, b float32, er error) *Decorator {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// Regex returns an error if the value does not match the pattern.
func (v *Decorator) Regex(value string, pattern string, err error) *Decorator {
	return v.ValidateFunc(func() error { return Regex(value, pattern, err) })
}

// WithRegex returns an error if the value does not match the pattern.
func (v *Decorator) WithRegex(value string, exp *regexp.Regexp, err error) *Decorator {
	return v.ValidateFunc(func() error { return WithRegex(value, exp, err) })
}

// Email returns an error if the value is not a valid email address.
func (v *Decorator) Email(value string, err error) *Decorator {
	return v.ValidateFunc(func() error { return Email(value, err) })
}

// New returns a new Decorator.
func New() *Decorator {
	return &Decorator{}
}
