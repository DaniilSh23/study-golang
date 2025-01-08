package weather

import (
	"DaniilSh23/weather-cli/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Получение погоды по геолокации
func GetWeather(geo geo.GeoData, showWeatherFormat int) (string, error) {
	
	// Используем модуль url для создания объекта URL адреса
	apiHost := "https://wttr.in/"
	baseUrl, err := url.Parse(apiHost + geo.City)
	if err != nil {
		fmt.Printf("Не удалось сформировать URL для получения погоды из: %v + %v\n", apiHost, geo.City)
		return "", err
	}

	// Формируем query параметры запроса
	queryParams := url.Values{}
	queryParams.Add("format", fmt.Sprint(showWeatherFormat))

	// Добавляем query параметры к нашему URL
	baseUrl.RawQuery = queryParams.Encode()

	// Выполняем запрос
	response, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}

	// Читаем тело ответа
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("Faild to read response body while GET weather.")
	}

	// Приводим responseBody к строке из байтов и возвращаем из функции
	return string(responseBody), nil
}

