// 1) 12 chi uy vazifada ValidateUser funksiya yaratilgan edi.
// 2) Bugungi test_repo repository yaratilgan ichida o'zgarishlar kiritilishi kerak
// 3) ValidateUser dasturri test_repo ni ichiga joylang
// 4) user package yaratib uning ichida Validate digan funksiyani e'lon qiling
// 6) main package da user.Validate qilingan holda chaqirinsin
// 7) dastur ishlashini hosil qiling (go run yoki go build o'rqali)
// 8) barchasi ishlagan taqdirda, hamma o'zgarishlarni git add, git commit va git push commandar o'rqali bajaring
// 9) natijada local codelaringiz github ga joylangan bolishi kerak
// 10) Repository linkini jonating

//  1. Oldindan belgilangan qoidalarga muvofiq foydalanuvchi
//     kiritishini ValidateUser digan method yarating.
//
// Qoidalar:
// Name empty bo'lishi kerak emas
// Name uzunligi kamida 6 belgigi bo'lishi kerak
// Age 0 dan kotta va 120 dan kichik bo'lishi
// Email empty bo'lishi kerak emas
// Email formatiga mos bolishi kerak (masalan example@domain.com)
//
// 2. Error slice yaratilgan holda barcha paydo bo'lgan errorlarni qoshib yuvo
// 3. Foydalanuvchi ma'lumotlarni terminaldan oqib oling
// 4. Oqib oliniyatgan jarayonda errorlarni ohirida chiqarib berin
//
// Masalan:
// Name: asd
// Age: 123
// Email: ""

// Errors:
// Name: length cannot be less than a 6 characters
// Age: couldn't be more than 120
// Email: couldn't be empty

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Fedy1507/repo_test/valuser"
)

type User struct {
	Name  string `validate:"required,min=6"`
	Age   int    `validate:"gte=0,lte=120"`
	Email string `validate:"required,email"`
}

func main() {
	var name string
	var age int
	var email string

	fmt.Printf("Ism: ")
	fmt.Scan(&name)
	fmt.Printf("Yosh: ")
	fmt.Scan(&age)
	fmt.Printf("Email: ")
	fmt.Scan(&email)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println()

	user := User{
		Name:  name,
		Age:   age,
		Email: email,
	}

	// Foydalanuvchi kiritishini tekshirish
	errorMessages := valuser.Validate(valuser.User(user))

	// Foydalanuvchi ma'lumotlarini chiqaring
	fmt.Printf("Ism: %s\n", user.Name)
	fmt.Printf("Yosh: %d\n", user.Age)
	fmt.Printf("Email: %s\n", user.Email)

	// Validatsiya xatoliklarini chiqaring
	fmt.Println("\nXatoliklar:")
	for _, errMsg := range errorMessages {
		fmt.Println(errMsg)
	}
}
