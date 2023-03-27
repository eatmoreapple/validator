package validator

import (
	"errors"
	"testing"
)

func TestValidator(t *testing.T) {
	type Item struct {
		Username  string
		Password  string
		Rpassword string
	}
	item := Item{
		Username:  "eatmoreapple",
		Password:  "password",
		Rpassword: "password",
	}
	v := New()
	if err := v.ValidateFunc(func() error {
		return MinLength[string](item.Username, 6, errors.New("username too short"))
	}).Eq(item.Password, item.Rpassword, errors.New("password not match")).Validate(); err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
