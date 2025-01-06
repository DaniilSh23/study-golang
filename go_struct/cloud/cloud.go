package cloud

import "fmt"

// Структура, отражающая облачное хранилище данных
type CloudDB struct {
	host string
}

// Конструктор для создания структуры облачного хранилища
func InitCloudDB(host string) *CloudDB {
	return &CloudDB{
		host: host,
	}
}

// Метод для чтения данных из облачного хранилища
func (db *CloudDB) Read() ([]byte, error) {
	fmt.Printf("Читаем данные из облака\n")
	return []byte("Hello world!"), nil
}

// Метод для записи данных в облачное хранилище
func (db *CloudDB) Write(content []byte) {
	fmt.Printf("Записываем данные в облако | %v\n", content)
}
