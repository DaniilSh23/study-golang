package main

import "fmt"

func main() {
	makeFuncExample()
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

func sliceFunc() {
	// Функция для иллюстрации работы со слайсами (срезами)

	myarr := [4]int{0, 1, 2, 3} // Создаем массив
	arrSlice := myarr[1:4]      // Берем срез массима с 1 по 4 эле-т (невключительно)
	fmt.Printf("Слайс с 1 по 4 эл-т: %v\n", arrSlice)

	arrSlice2 := myarr[:3] // Слайс сначала массива и до 3 эл-та [0 1 2]
	fmt.Printf("Слайс сначала и по 3-й эл-т: %v\n", arrSlice2)

	arrSlice3 := myarr[2:] // Слайс со 2 индекса и до конца
	fmt.Printf("Слайс со 2-го эл-та и до конца: %v\n", arrSlice3)
}

func workWithSlicesAndArrays() {
	// Функция для иллюстрации работы со слайсами и массивами

	myarr := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("Массив myarr %v\n", myarr)

	fmt.Println("====================")

	// Присваиваем массив новой переменной, при присвоении он копируется в новую ячейку памяти
	myarrNew := myarr
	myarr[0] = 10 // Изменения массимв myarr не влияют на myarrNew
	fmt.Printf("Массив myarrNew: %v\n", myarrNew)
	fmt.Printf("Массив myarr %v\n", myarr)

	fmt.Println("====================")

	// Создаем слайс из массива myarr: cлайс при создании ссылается на оригинальный массив!
	myarrPart := myarr[1:]
	myarrPart[0] = 99 // Меняем 0-й эл-т слайса (это 1-й эл-т myarr) на 99, изменения повлияют также на массив myarr!
	fmt.Printf("Слайс myarrPart %v\n", myarrPart)
	fmt.Printf("Массив myarr %v\n", myarr)

	fmt.Println("====================")

	/*
		Создаем слайс на основе другого слайса:
			! все слайсы будут созданы из массива myarr и изменения любого из слайсов меняют оригинальный массив.
	*/

	// Создаем новый слайс и меняем его 1-й эл-т (является 1-м эл-том myarrPart и 2-м для myarr)
	myarrPartNew := myarrPart[:2]
	myarrPartNew[1] = 88

	fmt.Printf("Слайс myarrPart %v\n", myarrPart)       // Видим, что изменения в myarrPartNew коснулись и myarrPart
	fmt.Printf("Массив myarr %v\n", myarr)              // Изменения, внесенные в myarrPartNew применяются на myarr
	fmt.Printf("Слайс myarrPartNew %v\n", myarrPartNew) // А любой слайс строится на основании своего массива

	fmt.Println("====================")

	/*
		Длина (len) и вместимость (cap):
			! Вместимость слайса (cap) основана на размере массива и считается от эл-та с которого был взят слайс,
			до конца масива. Например:
				Индексы массива:_____0__1__2__3_4

				Массив myarr        [10 99 88 3 4] - массив состоит из 5-и эл-в
				Слайс myarrPart        [99 88 3 4] - слайс, взятый с 1-го индекса и до конца, вмещает 4 эл-та
				Слайс myarrPartNew     [99 88 _ _] - слайс с 1-м и 2-м эл-ми массива вмещает также еще 2 эл-та, т.е. 4
	*/
	fmt.Println("Длина слайса myarrPart: ", len(myarrPart), "| Вместимость слайса myarrPart: ", cap(myarrPart))
	fmt.Println("Длина слайса myarrPartNew: ", len(myarrPartNew), "| Вместимость слайса myarrPartNew: ", cap(myarrPartNew))

	// Мы можем расширять слайс вправо, т.к. он полностью основан на массиве
	myarrPartExt := myarrPart[1:3]
	fmt.Printf("Слайс myarrPartExt %v | cap: %v\n", myarrPartExt, cap(myarrPartExt))

	// Расширяем слайс вправо до конца основного массива. Делаем это на основе того же слайса, т.к. он имеет cap == 3
	myarrPartExt = myarrPartExt[0:3]
	fmt.Printf("Слайс myarrPartExt %v\n", myarrPartExt)
}

func dynamicArray() {
	// Функция, которая иллюстрирует работу с динамическими массивами

	/*
	 Создадим слайс, без указания длины.
	 Под капотом GO создает массив, выделяет под него пространство в памяти, но мы не ограничиваемся заранее заданной длиной.
	*/
	mySlice := []int{0, 1, 2, 3, 4}

	/*
		Добавление элемента в слайс без длины.
		Так нельзя сделать: mySlice[5] = 100
		append() проверяет capacity (вместимость) слайса, и, если вместимости недостаточно, то он ее увеличивает, а затем
		добавляет новый эл-т в слайс. Данная функция возвращает новый слайс, это значит, что для добавления эл-та в существующий
		слайс мы можем его переприсвоить, или же создать новый слайс с добавленным эл-том.
	*/
	newSlice := append(mySlice, 100) // Создаем новый слайс с добавленным эл-том
	mySlice = append(mySlice, 88)    // Переопределяем существующий слайс, добавляя новый эл-т

	newSlice = append(newSlice, 99, 777, 333) // Добавление нескольких элементов в слайс

	// Добавление в слайс другого слайса. Троеточие (...) - это распаковка эл-тов слайса
	otherSlice := []int{11, 22, 33}
	otherSlice = append(otherSlice, newSlice...)

	// Добавление в слайс другого слайса, как отдельный эл-т (без распаковки)
	sliceInSliceLikeElem := [][]int{} // Создаем слайс и говорим, что в нем будут хранится значения с типо: слайс int
	sliceInSliceLikeElem = append(sliceInSliceLikeElem, mySlice, newSlice)

	fmt.Println("mySlice == ", mySlice)
	fmt.Println("newSlice == ", newSlice)
	fmt.Println("otherSlice == ", otherSlice)
	fmt.Println("sliceInSliceLikeElem == ", sliceInSliceLikeElem)
}

func transactionsArrayExcercise() {
	/*
		Функция по упражнению: массив транзакций.
		1) Запросить у пользователя его транзакции (по одному значению из терминала, если 0, то останавливаемся)
		2) Вывести итоговую сумму всех транзакций (подсчитать итоговый баланс)
	*/

	transactionSlice := []float64{}

	var balance float64
	for {
		transaction := scanTransaction()

		// Останавливаем цикл, если транзакция равна нулю
		if transaction == 0 {
			fmt.Printf("Итоговый баланс: %.2f\n", balance)
			break
		}

		transactionSlice = append(transactionSlice, transaction)
		balance += transaction
	}
	fmt.Println("Слайс транзакций == ", transactionSlice)
}

func scanTransaction() float64 {
	// Функция для запроса у юзера транзакции (какого-либо числа)

	var transaction float64
	fmt.Print("Введите транзакцию >> ")
	fmt.Scan(&transaction)
	fmt.Println("Введено: ", transaction)
	return transaction
}

func arrayCicle() {
	// Функция, демонстрирующая работу циклов с массивами (слайсами)

	myarr := [5]int{9, 8, 7, 6, 5}

	// Итерируемся по массиву
	for index, value := range myarr {
		fmt.Printf("index == %v | value == %v\n", index, value)
	}

	// Итерируемся по массиву, получая только эл-т массива (без индекса)
	for _, value := range myarr {
		fmt.Printf("Только value == %v\n", value)
	}

	// Итерируемся по массиву, получая только индекс (без эл-та массива)
	for index := range myarr {
		fmt.Printf("Тольо index == %v\n", index)
	}
}

func makeFuncExample() {
	// Функция для примера, как работает выделение памяти в GO и как работает функция make()

	// Создаем пустой массив строк, но заранее указываем, что его длина (len) == 2 эл-м
	myarr := make([]string, 2)
	fmt.Printf("Вместимость myarr == %v\n", cap(myarr))

	/*
		В myarr заданы два элемента - пустые строки.
		По дефолту тип string имеет значение пустой строки, поэтому они были установлены в массив при его создании
	*/
	fmt.Printf("Содержимое myarr == %v\n", myarr)

	/*
		Чтобы установить в myarr нужные элементы массива, нужно переопределить их по индексу
	*/
	myarr[0] = "string0"
	myarr[1] = "string1"
	fmt.Printf("Массив myarr == %v\n", myarr)

	/*
		Если для добавления в пустой массив элементов использовать append(), то новые элементы добавятся после дефолтных
		значений в виде пустых строк.
	*/
	emptyArr := make([]string, 2)
	emptyArr = append(emptyArr, "appendString1")
	emptyArr = append(emptyArr, "appendString2")
	fmt.Printf("Массив emptyArr == %v | len(emptyArr) == %v\n", emptyArr, len(emptyArr))

	/*
		make() также принимает и capacity массива, т.е. его максимальную вместимость.
		Таким образом, мы можем создать массив, в котором дефолтное значение длины,к примеру, по умолчанию будет равно нулю,
		но вместимость, например 2.
	*/
	newArr := make([]string, 0, 2)
	fmt.Printf("Массив newArr == %v | len(newArr) == %v | cap(newArr) == %v\n", newArr, len(newArr), cap(newArr))

	/*
		Теперь мы можем добавлять эл-ты в массив через append().
		Однако в данном случае cap == 2 - это означает, что при добавлениие 2-х эл-тов GO не будет тратить ресурсы на
		выделение памяти для расширения массива. Но если мы решим добавить 3-й эл-т, то GO придется занятся выделением
		памяти, т.к. вместимость (capacity) массива вышла за заданные нами же рамки.
	*/
	newArr = append(newArr, "new_elem0")
	newArr = append(newArr, "new_elem1")

	/*
		На этом шаге GO придется заняться выделением памяти для массива, т.к. изначально cap(newArr) == 2.
		И GO выделит память с запасом, т.е. если до добавления 3-го эл-та ("new_elem2") cap(newArr) было равно 2, то
		после добавления 3-го эл-та, cap(newArr) станет равно 4.
	*/
	newArr = append(newArr, "new_elem2")
	fmt.Printf("newArr == %v | len(newArr) == %v | cap(newArr) == %v\n", newArr, len(newArr), cap(newArr))
}
