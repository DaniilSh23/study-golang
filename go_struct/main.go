package main

import (
	"DaniilSh23/go_struct/account"
	"DaniilSh23/go_struct/encrypter"
	"DaniilSh23/go_struct/files"
	"DaniilSh23/go_struct/output"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Crypter encrypter.Encrypter

func main() {
	loadEnv()
	mainMenu()
}

// Загрузка переменных окружения
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		output.PrintError("Не удалось прочитать переменные окружения.")
		os.Exit(11)
	}
	cryptoKey := os.Getenv("CRYPTO_KEY")
	if cryptoKey == "" {
		output.PrintError("Отсутствует переменная окружения CRYPTO_KEY")
		os.Exit(12)
	}
	Crypter = *encrypter.InitEncrypter(cryptoKey)
}

// Главное меню
func mainMenu() {

	for {
		// Вывод меню и запрос выбора пользователя
		userChoice, err := getUserData("1 - Создать аккаунт\n2 - Найти аккаунт\n3 - Удалить аккант\n4 - Выход\n>>> ")
		if err != nil {
			continue
		}

		// Обработка выбора юзера
		choiceFunc := handleMenuChoice(userChoice)
		if choiceFunc == nil {
			fmt.Printf("Ваш выбор: %#q | Выход...\n", userChoice)
			return
		}

		// Вызов функции с логикой, соответствующей выбранному пункту меню
		choiceFunc()
	}
}

// Обработка выбора пункта меню
func handleMenuChoice(userChoice string) func() {
	choiceMapping := map[string]func(){
		"1": createAccount,
		"2": findAccount,
		"3": deleteAccount,
	}
	return choiceMapping[userChoice]
}

// Ядро приложения "хранилище паролей"
func createAccount() {

	newacc, err := getAccountCreds()
	if err != nil {
		return
	}

	// Если у аккаунта не указан пароль, то генерируем его с длиной в 8 символов
	emptyPwd := newacc.CheckPassIsEmpty() // К "композитному" методу можно обращаться так
	if emptyPwd {
		newacc.GenerateAccPassword(8) // Или вот так
	}

	// Если у аккаунта не указан логин, то пишем ошибку и заканчиваем работу программы
	emptyLogin := newacc.CheckLoginIsEmpty()
	if emptyLogin {
		output.PrintError("Ошибка: логин не указан!")
		return
	}

	// Создаем хранилище аккаунтов и добавляем туда новый аккаунт
	vault := account.InitVault(files.InitJsonDB("encrypted_data"), &Crypter)
	vault.AddAccount(newacc)
}

// Найти аккаунт
func findAccount() {

	// Запросить у юзера URL аккаунта, который надо найти
	usrInput, err := getUserData("Введите URL или логин аккаунта >>> ")
	if err != nil {
		return
	}

	// Вызвать из Vault метод для поиска URL в структуре
	vault := account.InitVault(files.InitJsonDB("encrypted_data"), &Crypter)
	searchResult := vault.SearchAccount(usrInput)

	// Вывод результатов поиска
	if len(searchResult) == 0 {
		fmt.Printf("Ничего не найдено по запросу %s...\n", usrInput)
	}
	fmt.Printf(strings.Repeat("=", 10) + "\n")
	for _, acc := range searchResult {
		acc.ShowAccCreds()
		fmt.Printf(strings.Repeat("=", 10) + "\n")
	}
}

// Удалить аккаунт
func deleteAccount() {

	// Запрос URL аккаунта для удаления
	usrInput, err := getUserData("Введите URL аккаунта >>> ")
	if err != nil {
		return
	}

	// Вызов метода удаления у Vault
	vault := account.InitVault(files.InitJsonDB("encrypted_data"), &Crypter)
	delResult := vault.DeleteAccount(usrInput)

	// Информирование о результате
	switch delResult {
	case true:
		fmt.Printf("Успешное удаление.\n")
	case false:
		fmt.Printf("Удаление не удалось, возможно аккаунт не существует.\n")
	}
}

// Получение кредов для нового аккаунта
func getAccountCreds() (*account.Account, error) {
	login, _ := getUserData("Введите логин >> ")
	password, _ := getUserData("Введите пароль (ENTER чтобы сгенерировать) >> ")
	url, _ := getUserData("Введите адрес сервиса >> ")

	// Получаем аккаунт с датой и временем
	newacc, err := account.InitAccount(login, password, url)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return nil, err
	}

	// Получение метаинформации о структуре
	// field, _ := reflect.TypeOf(newacc).Elem().FieldByName("Login")
	// fmt.Printf("Метаинформация (тэг) для поля login структуры Account: %v\n", string(field.Tag))

	return newacc, nil
}

// Запрос данных у пользователя
func getUserData(request_text string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(request_text)

	/*
		Считываем строку до нажатия ENTER (обязательно одинарные кавычки, так как ReadString принимает байты. В двойных значение \n было бы строкой)
	*/
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	input = strings.TrimSpace(input)
	return input, err
}
