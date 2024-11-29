package main

import (
	"fmt"
)

func main() {
	reverseArrExcercise()
}

// Иллюстрация работы указателей (references)

type strSlice = []string
type strArr = [1]string

func showRefernces() {

	// Создаем слайс и меняем его в функцие
	mySlice := strSlice{"1"}
	addElemToSlice(mySlice)
	fmt.Println("mySlice == ", mySlice)

	// Создаем массив и меняем его в функцие
	myArr := strArr{"a"}
	addElemToArr(myArr)
	fmt.Println("myArr == ", myArr)

	// Создаем переменную и меняем ее в функцие
	value := 1
	changeValue(value)
	fmt.Println("value == ", value)
	fmt.Println("&value == ", &value)

	// Создаем переменную и меняем ее в функцие, передав по ссылке
	newValue := 2
	fmt.Println("&newValue == ", &newValue, " | newValue == ", newValue)
	changeValueThroughReference(&newValue)
	fmt.Println("После изменения newValue == ", newValue)
}

func addElemToSlice(inputSlice strSlice) {
	/*
		Слайс является ссылочным типом, поэтому когда мы передаем его в функцию, то он меняется везде,
		так как хранится в одной и той же ячейке памяти.
	*/
	inputSlice[0] = "2"
}

func addElemToArr(inputArr strArr) {
	/*
		Массив неявляется ссылочным типом. Когда мы передаем его в функцию, то фактически создается копия массива,
		которая хранится в отдельной ячейки памяти.
	*/
	inputArr[0] = "b"
}

func changeValue(value int) {
	/*
		Переменные, при передаче в функцию копируются, а соответственно хранятся в иной ячейке памяти.
		Изменения value внутри этой функции не повлияют на оригинал, который был сюда передан.
	*/
	value = 2
	valueAddress := &value
	fmt.Printf("Переменная value в функцие changeValue == %v | Адрес в памяти: %v\n", value, valueAddress)
}

func changeValueThroughReference(value *int) {
	/*
		Изменяем переменную, через передачу значения по ссылке.
		Синтаксис в определении функции value *int - означает, что переменная value - это ссылка на int
	*/
	fmt.Println("changeValueThroughReference | value == ", value)
	fmt.Println("changeValueThroughReference | *value == ", *value)
	*value += 1
	fmt.Println("changeValueThroughReference | после изменения *value == ", *value)
}

// Упражнение reverse array
func reverseArrExcercise() {
	myarr := [4]int{1, 2, 3, 4}
	fmt.Println("myarr == ", myarr)
	reverseArr(&myarr)
	fmt.Println("myarr == ", myarr)
}

func reverseArr(arr *[4]int) {
	// Обращаемся к значению массива через *arr
	last_index := len(*arr) - 1
	half_array_index := len(*arr) / 2
	// Идем до половины массива и на каждой итерации меняем местами элементы (1й с последним, 2й с предпоследним и т.д.)
	for index := 0; index <= half_array_index; index++ {
		replace_index := last_index - index
		(*arr)[replace_index], (*arr)[index] = (*arr)[index], (*arr)[replace_index]
	}
}
