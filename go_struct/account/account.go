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
type Account struct {
	login string // Логин к сервису
	password string // Пароль к сервису
	url string // Адрес сервиса
}

// Метод для структуры account
func (Account) structMethod() {
	fmt.Println("Я метод структуры account")
}

/*
Если указать (acc account) без звездочки, т.е. вот так: (acc *account), то будет создана копия структуры и передана в функцию (метод структуры), а со звездочкой - будет передан указатель на структуру.
*/

// Еще один метод для структуры, в котором будут использоваться данные из структуры. 
func (acc *Account) structMethod2() {
	fmt.Printf("Метод для структуры со следующими полями:\nlogin: %v\npassword: %v\nurl: %v\n", acc.login, acc.password, acc.url)
}

// Показать креды аккаунта
func (acc *Account) ShowAccCreds() {
	cyanColor := color.New(color.FgCyan, color.Bold)
	cyanColor.Printf("Сервис: %v\nЛогин: %v\nПароль: %v\n", acc.url, acc.login, acc.password)
}

// Метод для создания пароля аккаунта
func (acc *Account) GenerateAccPassword(passLen int) {
	randPass := make([]string, passLen)
	
	for indx:=0; indx < int(passLen); indx++ {
		randElem := rand.IntN(65535)
		randPass[indx] = string(randElem)
	}
	acc.password = strings.Join(randPass, "")
}

// Проверка, что у аккаунта НЕ указан пароль
func (acc *Account) CheckPassIsEmpty() bool {
	if acc.password == "" {
		return true
	}
	return false
}

// Проверка, что у аккаунта НЕ указан логин
func (acc *Account) CheckLoginIsEmpty() bool {
	if acc.login == "" {
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
	createdAt time.Time
	updatedAt time.Time
	Acc Account // Встраиваем структуру account

	/*
	Допустим и такой вариант встраивания (явно именнованным полем). Но в таком случае будет недоступен короткий синтаксис вызова методов встроенной структуры и надо будет делать только вот так:

	newacc.acc.checkPassIsEmpty()
	*/ 	
	// acc account 
}


// Конструктор структуры account
func initAccount(login, password, urlString string) (*Account, error) {
	
	// Валидация url адреса
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL address.")
	}

	return &Account {
		login: login,
		password: password,
		url: urlString,
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
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Acc: Account{
			login: login,
			password: password,
			url: urlString,
		},
	}, nil
}

