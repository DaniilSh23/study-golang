package geo_test

import (
	"DaniilSh23/weather-cli/geo"
	"reflect"
	"testing"
)

// Тест конструктора структуры geo.GeoData
func TestInitGeoData(test *testing.T) {

	// Arrange - подготовка + определяем ожидаемый результат
	city := "Sevastopol"
	expectedType := geo.GeoData {City: city}

	// Act - выполняем действия
	result, err := geo.InitGeoData(city)

	// Assert - сверяем результат
	if err != nil {test.Errorf("Ошибка при создании объекта geo.GeoData: %v", err)}

	if reflect.TypeOf(result) != reflect.TypeOf(&expectedType) {
		test.Errorf("Неверный тип объекта в конструкторе geo.GeoData! Ожидался: %v, получен: %v", expectedType, result)
	}

	if city != result.City {
		test.Errorf("Неверный параметр City в конструкторе объекта geo.geo.GeoData! Ожидалось: %v, получено: %v", city, result.City)
	}
}
