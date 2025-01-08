package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)


type GeoData struct {
	City string `json:"city"`
}

type ValidateCityResponse struct {
	Error bool `json:"error"`
}

// Функция-конструктор объекта геолокации
func InitGeoData(city string) (*GeoData, error) {
	
	// Самостоятельно формируем структуру геолокации, если город был передан в функцию
	if city != "" {

		// Валидация названия города
		isValid, err := IsValidCity(city)
		if err != nil {
			fmt.Printf("Не удалось выполнить валидацию города: %v\n", err)	
			os.Exit(21)
		}
		if !isValid {
			fmt.Printf("Невалидное название города: %v\n", city)
			os.Exit(22)
		}

		return &GeoData {
			City: city,
		}, nil
	}

	// Получаем геолокацию по IP адресу юзера через API
	response, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, ErrorNot200Status(response.StatusCode)
	}
	
	// Читаем все из тела ответа (тело ответа является ридером потока данных)
	response_data, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, errors.New("Faild to read response body while GET geo by IP.")
	}

	// Упаковываем данные ответа в структуру (там будут только описанные в структуре поля)
	var geo GeoData
	json.Unmarshal(response_data, &geo)
	return &geo, nil 
}

// Валидация названия города
func IsValidCity(city string) (bool, error) {

	// Готовим данные для запроса
	apiHost := "https://countriesnow.space/api/v0.1/countries/population/cities"
	reqBody, err := json.Marshal(
		map[string]string{"city": city},
	)
	if err != nil {
		return false, err
	}
	reqContentType := "application/json"
	
	// Выполняем запрос
	response, err := http.Post(apiHost, reqContentType, bytes.NewBuffer(reqBody))
	if err != nil {
		return false, err
	}

	response_data, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return false, err
	}

	var validateCityResp ValidateCityResponse
	json.Unmarshal(response_data, &validateCityResp)
	return !validateCityResp.Error, nil
}

