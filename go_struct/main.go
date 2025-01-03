package main

import (
	"DaniilSh23/go_struct/account"
	"DaniilSh23/go_struct/files"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mainMenu()
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
		"2": searchAccount,
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
		fmt.Println("Ошибка: логин не указан!")
		return
	}

	// Преобразуем структуру аккаунта в байты и сохраняем в файл JSON
	byteArr := newacc.ToBytes()
	if byteArr == nil {
		return
	}
	files.WriteFile(byteArr, "data.json")
}

// Найти аккаунт
func searchAccount() {}

// Удалить аккаунт
func deleteAccount() {}

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


