package account

import (
	"DaniilSh23/go_struct/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)


// Структура, отражающая данные из хранилища
type VaultData struct {
	Accounts []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Структура, отражающая хранилище с БД
type Vault struct {
	Data VaultData `json:"data"`
	DB *files.JsonDB `json:"jsonDB"`
}

// Функция - конструктор структуры Vault
func InitVault(db *files.JsonDB) *Vault {

	// Пробуем достать данные Vault из JSON файла
	// db := files.InitJsonDB("data.json")
	data, err := db.ReadFile()
	
	vault := Vault {
			Data: VaultData{
				Accounts: []Account{},
				UpdatedAt: time.Now(),
			},
			DB: db,
		}

	// Возвращаем пустую структуру Vault, если достать данные из JSON не удалось
	if err != nil {
		return &vault
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
	
	// Записываем обновленную структуру в файл
	content := vault.ToBytes()
	vault.DB.WriteFile(content)
}

// Приведение структуры данных хранилища к байтам для дальнейшей записи в файл 
func (vault *Vault) ToBytes() []byte  {
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
		searched := strings.Contains(acc.Url, searchText)
		if searched {
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
			vault.Data.Accounts = append(vault.Data.Accounts[:indx], vault.Data.Accounts[indx + 1:]...) // Удаление - создаем новый слайс, исключив элемент, подлежащий удалению
			isDeleted = true
			vault.Data.UpdatedAt = time.Now()
			break
		}
	}
	
	// Записываем обновленную структуру в файл
	content := vault.ToBytes()
	vault.DB.WriteFile(content)
	
	return isDeleted
}

