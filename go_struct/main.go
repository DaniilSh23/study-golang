package main

import (
	"DaniilSh23/go_struct/account"
	"DaniilSh23/go_struct/files"
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

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

	// Если у аккаунта не указан пароль, то генерируем его с длиной в 8 символов
	emptyPwd := newacc.Acc.CheckPassIsEmpty() // К "композитному" методу можно обращаться так
	if emptyPwd {
		newacc.Acc.GenerateAccPassword(8) // Или вот так
	}

	// Если у аккаунта не указан логин, то пишем ошибку и заканчиваем работу программы
	emptyLogin := newacc.Acc.CheckLoginIsEmpty()
	if emptyLogin {
		fmt.Println("Ошибка: логин не указан!")
		return
	}
	
	newacc.Acc.ShowAccCreds()

	files.WriteFile()
}

// Получение кредов для нового аккаунта
func getAccountCreds() (*account.AccountWithTimestamp, error) {
	login, _ := getUserData("Введите логин >> ")
	password, _ := getUserData("Введите пароль (ENTER чтобы сгенерировать) >> ")
	url, _ := getUserData("Введите адрес сервиса >> ")

	// Получаем аккаунт с датой и временем
	// newacc, err := initAccount(login, password, url)
	newacc, err := account.InitAccountWithTimestamp(login, password, url)
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

