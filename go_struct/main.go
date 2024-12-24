package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "math/rand/v2"
)


// Структура, описывающая аккаунт, пароль для которого храним
type account struct {
	login string // Логин к сервису
	password string // Пароль к сервису
	url string // Адрес сервиса
}

// Метод для структуры account
func (account) structMethod() {
	fmt.Println("Я метод структуры account")
}

// Еще один метод для структуры, в котором будут использоваться данные из структуры
func (acc account) structMethod2() {
	fmt.Printf("Метод для структуры со следующими полями:\nlogin: %v\npassword: %v\nurl: %v\n", acc.login, acc.password, acc.url)
}

// Показать креды аккаунта
func (acc account) showAccCreds() {
	fmt.Printf("Сервис: %v\nЛогин: %v\nПароль: %v\n", acc.url, acc.login, acc.password)
}

func main() {
	// corePasswordStorage()
	// runeExample()
	newPass := generateRandPassword(7)
	fmt.Println("GENERATED NEW PASS: ", newPass)
}

// Ядро приложения "хранилище паролей"
func corePasswordStorage() {
	newacc := getAccountCreds()
	newacc.showAccCreds()
}

// Получение кредов для нового аккаунта
func getAccountCreds() account {
	login, _ := getUserData("Введите логин >> ")
	password, _ := getUserData("Введите пароль >> ")
	url, _ := getUserData("Введите адрес сервиса >> ")
	return account {
		login: login,
		password: password,
		url: url,
	}
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



// Пример работы с рунами
func runeExample() {
	/*
	Руна (rune) - это не что иное, как представление в виде unicode символов строки, по своей сути является int32.
	*/
	hello := "Hello world!)"
	for _, char := range hello {
		fmt.Println(char, string(char))
	}

	/*
	Мы можем также увидеть, что rune - это просто alias к int32. Давай сделаем массив rune из все той же приветственной строки. Если навестить на rune и узнать, что она из себя представляет, то увидим, что ее определение выглядит следующим образом:

	type rune = int32

	То есть я могу создать свою руну, как отдельный alias данного типа, выполнив, например, такой код:

	type myRune = int32

	Соответственно и в коде ниже если создать не массив rune, а массив int32, то ничего не изменится, эти записи эквивалентны.

	helloArr := []int32(hello)
	*/
	helloArr := []rune(hello)
	for _, char := range helloArr {
		fmt.Println(char, string(char))
	}
}


/*
	Упражнение "генерация пароля".
	Задача написать функцию, которая примет на вход целое число - количество символов в случайном пароле и сгенерирует этот самый пароль, вернув строку. Для этого понадобится импортировать "math/rand/v2"
*/

func generateRandPassword(passLen int) string {

	// Создаем слайс длиной, равной заданной длине случайного пароля
	randPass := make([]string, passLen)

	// Выполняем кол-во итераций, равное заданной длине случайного пароля
	for indx:=0; indx < int(passLen); indx++ {

		// На каждой итерации генерируем случайное число от 0 до 65535 (uint16)
		randElem := rand.IntN(65535)
		// И ставим его по нужному индексу в слайс случайного пароля, преобразовав в строку
		randPass[indx] = string(randElem)
	}

	// Возвращаем сгенерированный пароль, объединив слайс с символами пароля в одну строку
	return strings.Join(randPass, "")
}

