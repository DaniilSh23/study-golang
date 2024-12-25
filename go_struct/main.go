package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"os"
	"strings"
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

/*
Если указать (acc account) без звездочки, т.е. вот так: (acc *account), то будет создана копия структуры и передана в функцию (метод структуры), а со звездочкой - будет передан указатель на структуру.
*/

// Еще один метод для структуры, в котором будут использоваться данные из структуры. 
func (acc *account) structMethod2() {
	fmt.Printf("Метод для структуры со следующими полями:\nlogin: %v\npassword: %v\nurl: %v\n", acc.login, acc.password, acc.url)
}

// Показать креды аккаунта
func (acc *account) showAccCreds() {
	fmt.Printf("Сервис: %v\nЛогин: %v\nПароль: %v\n", acc.url, acc.login, acc.password)
}

// Метод для создания пароля аккаунта
func (acc *account) generateAccPassword(passLen int) {
	randPass := make([]string, passLen)
	
	for indx:=0; indx < int(passLen); indx++ {
		randElem := rand.IntN(65535)
		randPass[indx] = string(randElem)
	}
	acc.password = strings.Join(randPass, "")
}

// Проверка, что пароль у аккаунта НЕ указан
func (acc *account) checkPassIsEmpty() bool {
	if acc.password == "" {
		return true
	}
	return false
}

// Конструктор структуры account
func initAccount(login, password, urlString string) (*account, error) {
	
	// Валидация url адреса
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL address.")
	}

	return &account {
		login: login,
		password: password,
		url: urlString,
	}, nil
}


func main() {
	corePasswordStorage()
	// runeExample()
	// _ := generateRandPassword(7)
}

// Ядро приложения "хранилище паролей"
func corePasswordStorage() {
	newacc, err := getAccountCreds()
	if err != nil {
		return
	}
	emptyPwd := newacc.checkPassIsEmpty()
	
	// Если у аккаунта не указан пароль, то генерируем его с длиной в 8 символов
	if emptyPwd {
		newacc.generateAccPassword(8)
	}
	
	newacc.showAccCreds()
}

// Получение кредов для нового аккаунта
func getAccountCreds() (*account, error) {
	login, _ := getUserData("Введите логин >> ")
	password, _ := getUserData("Введите пароль (ENTER чтобы сгенерировать) >> ")
	url, _ := getUserData("Введите адрес сервиса >> ")
	newacc, err := initAccount(login, password, url)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return nil, err
	}
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

