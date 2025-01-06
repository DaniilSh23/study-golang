package output

import (
	"github.com/fatih/color"
)

// Функция - пример type switch через проверку типа
func TypeSwitchExample(value any) {

	// Создаем цвет текста для печати в терминал
	redPrint := color.New(color.FgRed, color.Bold)

	// Получение типа данных value (синтаксис только для switch-case)
	switch value := value.(type) {
	case string:
		redPrint.Printf(value + "\n")
	case int:
		redPrint.Printf("Код ошибки: %d\n", value)
	case error:
		redPrint.Printf("Error: %s\n", value.Error())
	default:
		redPrint.Printf("Неизвестная ошибка\n")
	}
}

// Функция пример определения типа переменной через if-else
func TypeSwitchByIf(value any) {

	// Создаем цвет текста для печати в терминал
	redPrint := color.New(color.FgRed, color.Bold)

	// Получение типа переменной и проверка через if
	value, ok := value.(string) // ok - bool об успешном извлечении значения заданного типа
	if ok {
		redPrint.Printf("Ошибка: %s\n", value)
		return
	}

	value, ok = value.(int)
	if ok {
		redPrint.Printf("Код ошибки: %d\n", value)
		return
	}

	redPrint.Printf("Неизвестная ошибка\n")
}
