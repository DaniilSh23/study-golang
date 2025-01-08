package main

import (
	"DaniilSh23/weather-cli/geo"
	"DaniilSh23/weather-cli/weather"
	"flag"
	"fmt"
	"os"
)


func main() {
	startApp()
}


// Функция с логикой, необходимой до запуска приложения
func prepareToStart() {}

// Функция с логикой для запуска приложения
func startApp() {
	fmt.Println("___Погода___")
	
	city, weatherShowFormat := parseFlags()
	geoData, err := geo.InitGeoData(city)
	if err != nil {
		fmt.Printf("Не удалось получить Ваши геоданные. %v\n", err.Error())
		os.Exit(11)
	}

	weatherData, err := weather.GetWeather(*geoData, weatherShowFormat)
	fmt.Printf("Погода в городе %v: %v\n", geoData.City, weatherData)
}

// Функция для парсинга флагов, переданных при старте программы. Возвращает (city, weatherShowFormat) 
func parseFlags() (string, int) {

	// Создаем флаги
	city := flag.String("city", "", "Город юзера")
	weatherShowFormat := flag.Int("show-format", 1, "Формат вывода погоды")
	
	// Парсим флаги из ввода пользователя в терминале
	flag.Parse()
	
	return *city, *weatherShowFormat
}


