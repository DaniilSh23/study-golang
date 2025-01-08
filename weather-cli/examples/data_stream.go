package examples

import (
	"flag"
	"fmt"
	"strings"
)

// Функция для парсинга флагов, переданных при старте программы
func parseFlags() {

	// Создаем флаги
	city := flag.String("city", "", "Город юзера")
	weatherShowFormat := flag.Int("show-format", 1, "Формат вывода погоды")
	
	// Парсим флаги из ввода пользователя в терминале
	flag.Parse()
	
	// ПРинтуем (потом удалить)
	fmt.Println(*city)
	fmt.Println(*weatherShowFormat)
}

// Пример Reader
func ExampleReader() {

	// Создаем новый ридер
	strFlow := "Это некий поток данных в виде строки"
	myreader := strings.NewReader(strFlow)

	// Задаем блок слайсла байт длиной 4, в который будем считывать поток
	block := make([]byte, 4)

	// Запускаем цикл для считывания потока поблочно
	for {
		readedBytes, err := myreader.Read(block)
		fmt.Println("Считано из потока (байт): ", readedBytes)
		
		// Останавливаем цикл, если в потоке больше ничего не осталось
		if readedBytes == 0 {
			fmt.Println("Поток прочитан полностью!")
			break
		}
		
		// ПРинтуем данные из потока, предварительно отформатировав их (перевели байты в символы)
		fmt.Printf("Считанный блок данных из потока: %q\n", block)

		if err != nil {
			fmt.Println("Ошибка при чтении потока!", err.Error())
			break
		}
	}
}
