package output

import "github.com/fatih/color"

func PrintError(value any) {

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
