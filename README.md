## Validate

validate is a simple validation library for golang.

### Installation

```bash
go get github.com/eatmoreapple/validator
```

### Usage

```go
package main

import (
	"errors"
	"log"

	"github.com/eatmoreapple/validator"
)

type User struct {
	Username  string
	Password  string
	Rpassword string
	Email     string
}

func main() {
	var user = User{
		Username:  "eatmoreapple",
		Password:  "password",
		Rpassword: "password",
		Email:     "eatmoreorange@gmail.com",
	}

	v := validator.New()

	v.StrMaxLength(user.Username, 15, errors.New("username too long"))
	v.StrMinLength(user.Username, 3, errors.New("username too short"))
	v.Eq(user.Password, user.Rpassword, errors.New("password not match"))
	v.Email(user.Email, errors.New("email not valid"))

	v.ValidateFunc(func() error {
		// custom validation
		return nil
	})

	if err := v.Validate(); err != nil {
		log.Fatal(err)
	}
	log.Println("success")
}

```