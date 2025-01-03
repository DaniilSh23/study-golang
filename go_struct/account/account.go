package account

import (
	"encoding/json"
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
	Login string `json:"login"` // Логин к сервису
	Password string `json:"password"` // Пароль к сервису
	Url string `json:"url"` // Адрес сервиса
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Показать креды аккаунта
func (acc *Account) ShowAccCreds() {
	cyanColor := color.New(color.FgCyan, color.Bold)
	cyanColor.Printf("Сервис: %v\nЛогин: %v\nПароль: %v\n", acc.Url, acc.Login, acc.Password)
}

// Метод для создания пароля аккаунта
func (acc *Account) GenerateAccPassword(passLen int) {
	randPass := make([]string, passLen)
	
	for indx:=0; indx < int(passLen); indx++ {
		randElem := rand.IntN(65535)
		randPass[indx] = string(randElem)
	}
	acc.Password = strings.Join(randPass, "")
}

// Проверка, что у аккаунта НЕ указан пароль
func (acc *Account) CheckPassIsEmpty() bool {
	if acc.Password == "" {
		return true
	}
	return false
}

// Проверка, что у аккаунта НЕ указан логин
func (acc *Account) CheckLoginIsEmpty() bool {
	if acc.Login == "" {
		return true
	}
	return false
}

// Конвертация структуры аккаунта в слайс байтов для записи в JSON файл
func (acc *Account) ToBytes() []byte {
	byteArray, err := json.MarshalIndent(acc, "", "	")
	if err != nil {
		fmt.Printf("Ошибка при конвертации структуры аккаунта к байтам: %v", byteArray)
		return nil
	}
	
	return byteArray
}


// Конструктор структуры 
func InitAccount(login, password, urlString string) (*Account, error) {
	
	// Валидация url адреса
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL address.")
	}

	return &Account {
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login: login,
		Password: password,
		Url: urlString,
	}, nil
}

