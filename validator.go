package validator

import (
	"regexp"
)

// Validator defines a Validate interface.
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

type Validate struct {
	group []Validator
}

// Validate implements the Validator interface.
func (v *Validate) Validate() error {
	return Group(v.group).Validate()
}

// Add adds a validator to the group.
func (v *Validate) Add(validator Validator) *Validate {
	v.group = append(v.group, validator)
	return v
}

// ValidateFunc adds a function to the group.
func (v *Validate) ValidateFunc(f func() error) *Validate {
	return v.Add(Func(f))
}

// StrMaxLength returns an error if the length of text is greater than length.
func (v *Validate) StrMaxLength(text string, length int, err error) *Validate {
	return v.ValidateFunc(func() error { return StrMaxLength(text, length, err) })
}

// StrMinLength returns an error if the length of text is less than length.
func (v *Validate) StrMinLength(text string, length int, err error) *Validate {
	return v.ValidateFunc(func() error { return StrMinLength(text, length, err) })
}

// Eq returns an error if a is not equal to b.
func (v *Validate) Eq(a, b any, err error) *Validate {
	return v.ValidateFunc(func() error { return Eq(a, b, err) })
}

// NotEq returns an error if a is equal to b.
func (v *Validate) NotEq(a, b any, err error) *Validate {
	return v.ValidateFunc(func() error { return NotEq(a, b, err) })
}

// GtInt returns an error if b is greater than a.
func (v *Validate) GtInt(a, b int, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt returns an error if b is less than a.
func (v *Validate) LtInt(a, b int, err error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, err) })
}

// GtInt8 returns an error if a is greater than b.
func (v *Validate) GtInt8(a, b int8, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt8 returns an error if a is less than b.
func (v *Validate) LtInt8(a, b int8, err error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, err) })
}

// GtInt16 returns an error if a is greater than b.
func (v *Validate) GtInt16(a, b int16, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt16 returns an error if a is less than b.
func (v *Validate) LtInt16(a, b int16, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtInt32 returns an error if a is greater than b.
func (v *Validate) GtInt32(a, b int32, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt32 returns an error if a is less than b.
func (v *Validate) LtInt32(a, b int32, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtInt64 returns an error if a is greater than b.
func (v *Validate) GtInt64(a, b int64, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtInt64 returns an error if a is less than b.
func (v *Validate) LtInt64(a, b int64, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint returns an error if a is greater than b.
func (v *Validate) GtUint(a, b uint, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt8 returns an error if a is less than b.
func (v *Validate) LtUInt8(a, b uint8, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint8 returns an error if a is greater than b.
func (v *Validate) GtUint8(a, b uint8, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// GtUint16 returns an error if a is greater than b.
func (v *Validate) GtUint16(a, b uint16, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt16 returns an error if a is less than b.
func (v *Validate) LtUInt16(a, b uint16, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint32 returns an error if a is greater than b.
func (v *Validate) GtUint32(a, b uint32, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt32 returns an error if a is less than b.
func (v *Validate) LtUInt32(a, b uint32, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtUint64 returns an error if a is greater than b.
func (v *Validate) GtUint64(a, b uint64, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtUInt64 returns an error if a is less than b.
func (v *Validate) LtUInt64(a, b uint64, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtFloat32 returns an error if a is greater than b.
func (v *Validate) GtFloat32(a, b float32, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtFloat32 returns an error if a is less than b.
func (v *Validate) LtFloat32(a, b float32, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// GtFloat64 returns an error if a is greater than b.
func (v *Validate) GtFloat64(a, b float64, err error) *Validate {
	return v.ValidateFunc(func() error { return Gt(a, b, err) })
}

// LtFloat64 returns an error if a is less than b.
func (v *Validate) LtFloat64(a, b float32, er error) *Validate {
	return v.ValidateFunc(func() error { return Lt(a, b, er) })
}

// Regex returns an error if the value does not match the pattern.
func (v *Validate) Regex(value string, pattern string, err error) *Validate {
	return v.ValidateFunc(func() error { return Regex(value, pattern, err) })
}

// WithRegex returns an error if the value does not match the pattern.
func (v *Validate) WithRegex(value string, exp *regexp.Regexp, err error) *Validate {
	return v.ValidateFunc(func() error { return WithRegex(value, exp, err) })
}

// Email returns an error if the value is not a valid email address.
func (v *Validate) Email(value string, err error) *Validate {
	return v.ValidateFunc(func() error { return Email(value, err) })
}

// New returns a new Validate.
func New() *Validate {
	return &Validate{}
}
