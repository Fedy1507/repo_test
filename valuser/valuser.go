package valuser

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

type User struct {
	Name  string `validate:"required,min=6"`
	Age   int    `validate:"gte=0,lte=120"`
	Email string `validate:"required,email"`
}

func Validate(user User) []string {
	validate := validator.New()
	var errors []string

	err := validate.Struct(user)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Fatal("Xatoliklar validatsiya xatolarini o'zgartirishda")
		}

		for _, e := range validationErrors {
			switch e.Field() {
			case "Name":
				errors = append(errors, fmt.Sprintf("Ism: 6 belgidan kam bo'lishi mumkin emas"))
			case "Age":
				errors = append(errors, fmt.Sprintf("Yosh: 120 dan katta bo'lishi mumkin emas"))
			case "Email":
				if strings.TrimSpace(user.Email) == "" {
					errors = append(errors, fmt.Sprintf("Email: bo'sh bo'lishi mumkin emas"))
				} else {
					errors = append(errors, fmt.Sprintf("Email: example@domain.com formatida bo'lishi kerak"))
				}
			}
		}
	}

	return errors
}
