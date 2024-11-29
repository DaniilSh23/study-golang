package main

import (
	"fmt"
)

func main() {
	showRefernces()
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
