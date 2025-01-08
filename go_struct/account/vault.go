package account

import (
	"encoding/json"
	"fmt"
	"time"
)

// Интерфейс, определяющий требования к хранилищам данных
type AnyDB interface {
	Read() ([]byte, error) // Метод для чтения данных
	Write([]byte)          // Метод для записи данных
}

// Интерфейс, определяющий требования к структуре шифрования данных
type AnyCrypter interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

// Структура, отражающая данные из хранилища
type VaultData struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Структура, отражающая хранилище с БД
type Vault struct {
	Data VaultData `json:"data"`
	DB   AnyDB     
	Crypt AnyCrypter
}

// Функция - конструктор структуры Vault
func InitVault(db AnyDB, crypter AnyCrypter) *Vault {

	// Пробуем достать данные Vault из JSON файла
	data, err := db.Read()

	vault := Vault{
		Data: VaultData{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		},
		DB: db,
		Crypt: crypter,
	}

	// Возвращаем пустую структуру Vault, если достать данные из JSON не удалось
	if err != nil {
		fmt.Println("Не удалось достать данные из файла. Возвращаем пустое Vault.")
		return &vault
	}
	
	// Расшифровываем данные
	if len(data) > 0{
		data = crypter.Decrypt(data)	
	}

	// Создаем структуру Vault из данных в JSON файле
	err = json.Unmarshal(data, &vault.Data)
	if err != nil {
		fmt.Printf("Не удалось привести к структуре Vault данные из JSON файла. Создаем пустое хранилище Vault\n")
		return &vault
	}

	return &vault
}

// Добавление нового аккаунта в хранилище
func (vault *Vault) AddAccount(acc *Account) {

	// Добавляем новый аккаунт в структуру
	vault.Data.Accounts = append(vault.Data.Accounts, *acc)
	vault.Save()
}

// Сохранение данных в БД
func (vault *Vault) Save() {
	
	content := vault.ToBytes()
	
	// Шифруем данные
	content = vault.Crypt.Encrypt(content)
	
	// Записываем обновленную структуру в файл
	vault.DB.Write(content)
}


// Приведение структуры данных хранилища к байтам для дальнейшей записи в файл
func (vault *Vault) ToBytes() []byte {
	byteArr, err := json.MarshalIndent(vault.Data, "", "	")
	if err != nil {
		fmt.Printf("Ошибка при конвертации в байты структуры Vault: %v", err)
		return nil
	}
	return byteArr
}

// Поиск аккаунтов по URL (поиск подстроки в строке)
func (vault *Vault) SearchAccount(searchText string) []Account {
	// Создаем слайс для результатов поиска с 0 длинной и вместимостью, равной кол-ву хранящихся аккаунтов
	accResults := make([]Account, 0, len(vault.Data.Accounts))

	// Обходим все аккаунты в структуре и ищем подходящие
	for _, acc := range vault.Data.Accounts {
		if acc.CheckIsMatched(searchText) {
			accResults = append(accResults, acc)
		}
	}

	return accResults
}

// Удаление аккаунта по URL
func (vault *Vault) DeleteAccount(accUrl string) bool {

	// Обходим все аккаунты в структуре и находим нужный для удаления
	isDeleted := false
	for indx, acc := range vault.Data.Accounts {
		if acc.Url == accUrl {
			fmt.Printf("Найден аккаунт с URL %v. Удаляем...\n", accUrl)
			vault.Data.Accounts = append(vault.Data.Accounts[:indx], vault.Data.Accounts[indx+1:]...) // Удаление - создаем новый слайс, исключив элемент, подлежащий удалению
			isDeleted = true
			vault.Data.UpdatedAt = time.Now()
			break
		}
	}
	
	vault.Save()
	return isDeleted
}
