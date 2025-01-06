package output

import "fmt"

/*
	Функция - иллюстрация generic. Она принимает две переменные и возвращает их сумму.

Для переменных определены два типа: все варианты int - это тип T и все варианты float - это тип V. В теле функции числа приводятся к типу float32 и возвращается их сумма
*/
func SumNumbs[T int | int8 | int16 | int32 | int64, V float32 | float64](a T, b V) V {
	return V(a) + V(b)
}

/*
	Одним из ограничений generic является невозможность получения типа в switch-case.

Но это можно обойти за счет приведения типов. Т.е. мы приведем значение переменной к типу value и все сработает.
*/
func TypeSwitchInGeneric[T int | float32 | string](value T) {
	switch val := any(value).(type) {
	case string:
		fmt.Printf("Получена строка %v\n", val)
	case float32:
		fmt.Printf("Получено float32 %v\n", val)
	case int:
		fmt.Printf("Получен int %v\n", val)
	}
}

/*
	Представим, что нам нужен такой объект как список элементов.

Создадим его в виде структуры, однако он сможет содержать внутри данные любого из типов
*/
type List[T any] struct {
	items []T
}

// Создание структуры с generic типом
var arr = List[string]{
	items: []string{"hello", "world"},
}
