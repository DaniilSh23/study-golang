package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"
	"github.com/fatih/color" // Импорт стороннего модуля
)

// Структура, описывающая аккаунт, пароль для которого храним
type AccountExample struct {
	Login string `json:"login" xml:"test"` // Логин к сервису
	Password string `json:"password"` // Пароль к сервису
	Url string `json:"url"` // Адрес сервиса
}

// Метод для структуры account
func (AccountExample) structMethod() {
	fmt.Println("Я метод структуры account")
}

/*
Если указать (acc account) без звездочки, т.е. вот так: (acc *account), то будет создана копия структуры и передана в функцию (метод структуры), а со звездочкой - будет передан указатель на структуру.
*/

// Еще один метод для структуры, в котором будут использоваться данные из структуры. 
func (acc *AccountExample) structMethod2() {
	fmt.Printf("Метод для структуры со следующими полями:\nlogin: %v\npassword: %v\nurl: %v\n", acc.Login, acc.Password, acc.Url)
}

// Показать креды аккаунта
func (acc *AccountExample) ShowAccCreds() {
	cyanColor := color.New(color.FgCyan, color.Bold)
	cyanColor.Printf("Сервис: %v\nЛогин: %v\nПароль: %v\n", acc.Url, acc.Login, acc.Password)
}

// Метод для создания пароля аккаунта
func (acc *AccountExample) GenerateAccPassword(passLen int) {
	randPass := make([]string, passLen)
	
	for indx:=0; indx < int(passLen); indx++ {
		randElem := rand.IntN(65535)
		randPass[indx] = string(randElem)
	}
	acc.Password = strings.Join(randPass, "")
}

// Проверка, что у аккаунта НЕ указан пароль
func (acc *AccountExample) CheckPassIsEmpty() bool {
	if acc.Password == "" {
		return true
	}
	return false
}

// Проверка, что у аккаунта НЕ указан логин
func (acc *AccountExample) CheckLoginIsEmpty() bool {
	if acc.Login == "" {
		return true
	}
	return false
}



/*
	Композиция структу. Это когда мы встраиваем одну структуру в другую. 
	Например добавим структуру accoutWithTimestamp, которая будет содержать все, что в структуре account и еще два поля, связанных со временем.
*/

// структура с композицией (аккаунт с полями времени)
type AccountWithTimestamp struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Acc AccountExample // Встраиваем структуру account

	/*
	Допустим и такой вариант встраивания (явно именнованным полем). Но в таком случае будет недоступен короткий синтаксис вызова методов встроенной структуры и надо будет делать только вот так:

	newacc.acc.checkPassIsEmpty()
	*/ 	
	// acc account 
}


// Конструктор структуры account
func initAccount(login, password, urlString string) (*AccountExample, error) {
	
	// Валидация url адреса
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL address.")
	}

	return &AccountExample {
		Login: login,
		Password: password,
		Url: urlString,
	}, nil
}

// Конструктор структуры accountWithTimestamp
func InitAccountWithTimestamp(login, password, urlString string) (*AccountWithTimestamp, error) {
	
	// Валидация url адреса
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL address.")
	}

	return &AccountWithTimestamp {
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Acc: AccountExample {
			Login: login,
			Password: password,
			Url: urlString,
		},
	}, nil
}

