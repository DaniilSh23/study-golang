package main

import "fmt"

func main() {

}

func arrayFunc() {
	// Функция для иллюстрации работы с массивами

	// Создание массива целых чисел, длиной 3 и с начальными эл-ми 5, 7, -3
	myarr := [3]int{5, 7, -3}

	// Массив строк, длиной 2, который мы отдельно определяем, а потом присваеиваем ему значение
	var strArr [2]string
	strArr = [2]string{"hello", "world"}

	// Объявление пустого массива строк
	emptyArr := [2]string{}

	// Записываем в переменную нулевой элемент массива myarr
	elem0 := myarr[0]

	// Задаем нулевой элемент массиву emptyArr
	emptyArr[0] = "Г-банк"

	fmt.Println(myarr)
	fmt.Println((strArr))
	fmt.Println(emptyArr)
	fmt.Println(elem0)
}
