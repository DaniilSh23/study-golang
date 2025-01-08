package files

import (
	"DaniilSh23/go_struct/output"
	"os"

	"github.com/fatih/color"
)

type JsonDB struct {
	filename string
}

// Конструктор структуры JsonDB
func InitJsonDB(filename string) *JsonDB {
	
	// Проверка наличия файла с данными
	_, err := os.Stat(filename)
	fileIsNotExist := os.IsNotExist(err)
	
	// Создание файла
	if fileIsNotExist {
		_, err := os.Create(filename)
		if err != nil {
			output.PrintError("Ошибка при создании файла\n" + filename + err.Error())
			os.Exit(5)
		}
	}

	return &JsonDB{
		filename: filename,
	}
}

// Функция для записи файла
func (db *JsonDB) Write(content []byte) {

	// Цвета для текста в терминале
	redColor := color.New(color.FgRed, color.Bold)
	greenColor := color.New(color.FgGreen, color.Faint)
	
	file, err := os.OpenFile(db.filename, os.O_RDWR, 0644)

	defer file.Close() // Откладываем выполнение закрытия файла на конец функции (defer закидывает операции в stack frames, в случае нескольких вызовов defer, отложенные операции будут выполняться по принципу стека LIFO)

	// Запись строки в файл
	size, err := file.Write(content)
	if err != nil {
		redColor.Printf("Ошибка при записи в файл %v\n", err)
		return
	}
	greenColor.Printf("Успешно записано %v байт в файл %v\n", size, file.Name())
}

// Функция для чтения файла
func (db *JsonDB) Read() ([]byte, error) {

	// Цвета для текста в терминале
	redColor := color.New(color.FgRed, color.Bold)
	greenColor := color.New(color.FgGreen, color.Faint)

	data, err := os.ReadFile(db.filename) // чтение файла целиком
	if err != nil {
		redColor.Printf("Ошибка при чтении файла %v: %v\n", db.filename, err)
		return nil, err
	}

	greenColor.Printf("Успешное чтение файла %v:\n", db.filename)
	return data, nil
}

