package files

import (
	"os"

	"github.com/fatih/color"
)

// Функция для записи файла
func WriteFile(content []byte, filename string) {
	
	// Цвета для текста в терминале
	redColor := color.New(color.FgRed, color.Bold)
	greenColor := color.New(color.FgGreen, color.Faint)
	
	// Создание файла
	newfile, err := os.Create(filename)
	if err != nil {
		redColor.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	
	defer newfile.Close() // Откладываем выполнение закрытия файла на конец функции (defer закидывает операции в stack frames, в случае нескольких вызовов defer, отложенные операции будут выполняться по принципу стека LIFO)
	
	// Запись строки в файл
	writedBytes, err := newfile.Write(content)
	if err != nil {
		redColor.Printf("Ошибка при записи в файл %v\n", err)
		return
	}
	greenColor.Printf("Успешно записано %v байт в файл %v\n", writedBytes, newfile.Name())
}


// Функция для чтения файла
func ReadFile(filename string) {
	
	// Цвета для текста в терминале
	redColor := color.New(color.FgRed, color.Bold)
	greenColor := color.New(color.FgGreen, color.Faint)

	data, err := os.ReadFile(filename)	// чтение файла целиком
	if err != nil {
		redColor.Printf("Ошибка при чтении файла %v: %v\n", filename, err)
	}

	greenColor.Printf("Успешное чтение файла %v:\n\n%v\n", filename, string(data))
}
