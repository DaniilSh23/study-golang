package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	// Println занимает всю строку (линию), т.е. следующий принт будет на новой строке без необходимости указывать \n
	fmt.Println("___ Расчет Вашего ИМТ ___")

	// var userKg float64 // если ничего не присвоено, то переменная имеет значение 0.0 (float64)
	// userHeight := 1.76

	for {
		programWork()

		var stopWord string
		fmt.Print("\n\nНажмите любую клавишу для повторного расчета.\nВведите STOP для завершения работы программы >>> ")
		fmt.Scan(&stopWord)
		if stopWord == "STOP" {
			break
		}
		fmt.Println("Выполняем повторный расчет ИМТ")
	}

}

func programWork() {
	// Основная логика работы программы

	userHeight, userKg, err := getUserInput()
	// Проверяем наличие ошибки в выполненние функции getUserInput
	if err != nil {
		// fmt.Println(err)
		// return
		panic(err)
	}

	IMT := calculateIMT(userKg, userHeight)

	checkIMTByIf(IMT)
	checkIMTBySwitch(IMT)

	outputResult(IMT)
}

func checkIMTByIf(imt float64) {
	// Проверка ИМТ и уведомления пользователя о его весе через условные операторы

	if imt < 16 {
		fmt.Println("У Вас сильный дефицит массы тела")
	} else if 16 <= imt && imt < 18.5 {
		fmt.Println("У Вас дефицит массы тела")
	} else if 18.5 <= imt && imt < 25 {
		fmt.Println("У Вас нормальная масса тела")
	} else if 25 <= imt && imt < 30 {
		fmt.Println("У Вас избыточная масса тела")
	} else if 30 <= imt && imt < 35 {
		fmt.Println("У Вас 1 степень ожирения")
	} else if 35 <= imt && imt < 40 {
		fmt.Println("У Вас 2 степень ожирения")
	} else {
		fmt.Println("У Вас 3 степень ожирения")
	}
}

func checkIMTBySwitch(imt float64) {
	// Проверка ИМТ и уведомления пользователя о его весе через switch-case
	switch {
	case imt < 16:
		fmt.Println("У Вас сильный дефицит массы тела")
	case 16 <= imt && imt < 18.5:
		fmt.Println("У Вас дефицит массы тела")
	case 18.5 <= imt && imt < 25:
		fmt.Println("У Вас нормальная масса тела")
	case 25 <= imt && imt < 30:
		fmt.Println("У Вас избыточная масса тела")
	case 30 <= imt && imt < 35:
		fmt.Println("У Вас 1 степень ожирения")
	case 35 <= imt && imt < 40:
		fmt.Println("У Вас 2 степень ожирения")
	default:
		fmt.Println("У Вас 3 степень ожирения")
	}
}

func outputResult(imt float64) {
	// Вывод результата расчета ИМТ в терминал

	// Printf принтует форматированную строку, т.е. мы можем докинуть в нее переменные
	fmt.Printf("Ваш ИМТ == %v", imt)

	// Можно также вывести и простым принтом, т.к. он принимает несколько аргументов
	fmt.Print("Ваш ИМТ == ", imt)

	// А можно float64 выводить форматировано, указывая сколько знаков после запятой хотим видеть (произойдет округление)
	fmt.Printf("Ваш ИМТ == %.0f", imt) // Здесь мы выведем 0 знаков после запятой в float64 (подробнее в доке модуля fmt)

	// Вот так можно сохранить результат форматирования строки в переменную (тип будет строкой) без вывода в консоль
	result := fmt.Sprintf("Ваш ИМТ == %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	// Функция для расчета ИМТ

	// IMTPower - это untyped int, что говорит о том, что это значение не строго типизировано, как int.
	// В данном контексте IMTPower может вести себя, как float64.
	// Если бы указали явно const IMTPower = 2, то уже получили бы ошибку
	return userKg / math.Pow(userHeight/100, IMTPower)
}

func getUserInput() (float64, float64, error) {
	// Получение ввода веса и роста от пользователя

	var userHeight, userKg float64 // Определяем одной строкой две переменные с одинаковым типом данных

	// fmt.Scan принимает ввод пользователя из терминала и кладет его в переменную, которую мы
	// передаем в функцию. Однако в саму функцию Scan нужно передавать не саму переменную, а указатель на переменную,
	// что мы и делаем знаком &
	fmt.Print("Введите свой рост (см): ")
	fmt.Scan(&userHeight)
	fmt.Print("\nВведите свой вес (кг): ")
	fmt.Scan(&userKg)

	// Проверка валидности роста и веса
	if userHeight <= 0 || userKg <= 0 {
		return 0, 0, errors.New("Неверно указан рост или вес")
	}

	return userHeight, userKg, nil
}

// Как можно определять функции

func myFunc1(numb1 float64) {
	// Функция с 1 аргументов
}

func myFunc2(numb1, numb2 float64) {
	// Функция с 2 аргументами одного типа данных
}

func myFunc3(numb1 uint16, str1 string) {
	// Функция с аргументами разного типа данны
}

func myFunc4(numb1 uint16, numb2 uint16) uint16 {
	// Функция с аргументами, которая возвращает значение типа uint16
	return numb1 + numb2
}

func myFunc5(numb1 uint16, flag bool) (uint16, bool) {
	// Функция с аргументами, которая возвращает два значения типа uint16 и bool
	return numb1, flag
}

func myFunc6(numb1 uint16) (returnVal uint16) {
	// Функция с аргументом и альтернативным синтаксисом возврата значений.
	// Мы определяем переменную, которая вернется из функции там же, где указываем тип данных возвращаемого значения

	returnVal = numb1 + 2 // Присваиваем новое значение переменной возврата
	return                // Переменную returnVal можно не указывать, Go сам ее подставит в return
}

func myFunc7(numb1 uint16) (returnVal1 uint16, returnVal2 string) {
	// Функция с аргументами и альтернативным синтаксисом возврата нескольких значений.

	// Присваиваем новое значение переменным возврата
	returnVal1 = numb1 + 2
	returnVal2 = "Hello world"
	return // Переменные returnValN можно не указывать, Go сам их подставит в return
}

func cyclesExample() {
	// Функция с примерами записей циклов

	for i := 0; i < 10; i++ {
		fmt.Printf("Итерация цикла: %d\n", i)
	}

	// Можно запустить цикл и так
	iteration := 0
	for iteration < 10 {
		fmt.Printf("Итерация цикла: %d\n", iteration)
		iteration++
	}

	// Цикл с break
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("Итерация %d | Вызываем break\n", i)
			break
		}
		fmt.Printf("Итерация цикла: %d\n", i)
	}

	// Цикл с continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("Итерация %d | Вызываем continue\n", i)
			continue
		}
		fmt.Printf("Итерация цикла: %d\n", i)
	}

	// Бесконечный цикл
	for {
		fmt.Println("Infinity cycle for")
	}
}
