package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Глобальная мапа для хранения закладок (вне функции недопустимо короткое объявление переменной через ":=" )
var bookmarksMap = map[string]string{}

func main() {
	mainMenu()
}

/*
Функция - главное меню приложения "закладки". "Закладки" - двусмысленно как-то...
Сразу представляю себе наркоманов, которые голыми руками ковыряют землю, где предположительно их ждет драгоценный клад,
на полянке, где бегает за палкой и все кругом пометила моя собака.
*/
func mainMenu() {
	fmt.Println("=====Приложение закладки=====")

	// Крутим бесконечный цикл, чтобы запрашивать ввод, пока юзер не выйдет
	for {

		// Запрашиваем выбор пункта меню у пользователя
		fmt.Printf("Выберите действие:\n1 - Посмотреть все закладки\n2 - Добавить закладку\n3 - Удалить закладку\n4 - Выход\n>>>")
		var menuChoice int
		_, err := fmt.Scan(&menuChoice)
		if err != nil {
			fmt.Printf("Ошибка ввода %v", err)
		}

		// Обрабатываем выбор пользователя
		handleChoiceFunc := handleMenuChoice(menuChoice)
		if handleChoiceFunc == nil {
			fmt.Println("Выход из приложения.")
			return
		}

		// Вызываем функцию с логикой выбора юзера
		handleChoiceFunc()
	}
}

// Функция обработки выбора юзера в главном меню
func handleMenuChoice(choice int) func() {
	choiceMapping := map[int]func(){
		1: showBookmarks,
		2: addBookmarks,
		3: deleteBookmarks,
		// 4: func() { fmt.Println("Выход из приложения.") },
	}
	return choiceMapping[choice] // Возвращаем функцию, которая содержит логику, согласно выбора юзера
}

// Функция для нажатия 1 - выводит все закладки
func showBookmarks() {
	fmt.Println("---Отображение всех закладок---")

	// Итерируемся по глобальной мапе с закладками и отображаем их
	for key, value := range bookmarksMap {
		// %q - форматирование строки для оборачивания ее в двойные кавычки
		fmt.Printf("Закладка: %q\nСодержание: %q\n----------\n", key, value)
	}
}

// Функция для нажатия 2 - добавление новой закладки
func addBookmarks() {
	fmt.Println("---Добавление новой закладки---")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите заголовок закладки\n>>>")
	// Если указать \n в "двойных кавычках" - это будет тип string, но метод ReadString принимает байты, поэтому 'одинарные'
	bookmarkTitle, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода! ", err)
	}
	bookmarkTitle = strings.TrimSpace(bookmarkTitle)

	fmt.Print("Введите текст закладки\n>>>")
	bookmarkText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода! ", err)
	}
	bookmarkText = strings.TrimSpace(bookmarkText)

	bookmarksMap[bookmarkTitle] = bookmarkText
}

// Функция для нажатия 3 - удаление закладки
func deleteBookmarks() {
	fmt.Println("---Удаление закладки---")

	fmt.Print("Введите заголовок закладки для удаления\n>>>")
	reader := bufio.NewReader(os.Stdin)       // Создаем ридер для считывания ввода из терминала
	userInput, err := reader.ReadString('\n') // Считываем все, пока юзер не введет ENTER (\n)
	if err != nil {
		fmt.Println("Ошибка ввода! ", err)
	}

	/*
		Удаляем все пробельные символы вначале и в конце ввода юзера
		(в т.ч. и \n который он ввел, нажав ENTER после завершения ввода)
	*/
	userInput = strings.TrimSpace(userInput)
	delete(bookmarksMap, userInput)
}

// Функция, иллюстрирующая синтаксис работы с map
func syntaxMap() {

	/*
		Создаем map, указываем в квадратных скобках тип данных ключа, за ними тип данных значение
		и в круглых скобках определяем содержимое.
	*/
	mymap := map[string]int{
		"key": 123,
	}
	fmt.Println("mymap == ", mymap)

	// Получаем значение мапы по ключу
	fmt.Printf("Значение mymap по ключу 'key' == %v\n", mymap["key"])

	// Переопределение значение мапы по ключу
	mymap["key"] = 321
	fmt.Printf("Значение mymap по ключу 'key' == %v\n", mymap["key"])

	// Добавление нового элемента в мап
	mymap["newElem"] = 777
	fmt.Println("mymap == ", mymap)

	/*
		Удаление элемента из мапы. Для этого нужно использовать встроенный метод delete()
		Ниже, для иллюстрации, я добавлю элемент в мап и удалю его потом
		Функция delete() принимает объект мапы, на котором надо произвести удаление и ключ для удаления элемента.
		Если попробовать удалить эл-т по несуществующему ключу, то ничего не будет.
		Если попробовать использовать (например вывести в терминал) элемент мапы по несуществубщему ключу,
		то будет просто дефолтное значение указанного типа данных, например для string - пустая строка, int - это ноль, ...
	*/
	fmt.Println("Добавляем в mymap пятый элемент.")
	mymap["fiveElement"] = 5
	fmt.Println("mymap == ", mymap)
	delete(mymap, "fiveElement")
	fmt.Println("mymap == ", mymap)

	delete(mymap, "elem404") // Пробуем удалить эл-т по несуществующему ключу
	// Принтуем несуществующий эл-т и получаем 0, т.к. это дефолтное значение для int
	fmt.Printf("Значение mymap по ключу 'elem404' == %v\n", mymap["elem404"])
}

// Функция для иллюстрации итерации по map
func iterationMap() {
	mymap := map[string]int{"a": 1, "b": 2}
	for key, val := range mymap {
		fmt.Println(key, val)
	}
}
