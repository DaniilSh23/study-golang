package encrypter

import (
	"DaniilSh23/go_struct/output"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

// Структура для шифрования и дешифрования паролей
type Encrypter struct {
	CryptoKey string
}


// Конструктор структуры Encrypter
func InitEncrypter(cryptoKey string) *Encrypter {
	return &Encrypter {
		CryptoKey: cryptoKey,
	}
}

// Метод шифрования строки
func (crypt *Encrypter) Encrypt(rawStr []byte) []byte {
	
	// Создается AES блочный шифр
	block, err := aes.NewCipher([]byte(crypt.CryptoKey))
	if err != nil {
		output.PrintError("Не удалось зашифровать данные аккаунтов")
		os.Exit(31)
	}
	
	// Устанавливается режим GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		output.PrintError("Не удалось зашифровать данные аккаунтов")
		os.Exit(32)
	}
	
	// Генерируется случайный nonce.
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		output.PrintError("Не удалось зашифровать данные аккаунтов")
		os.Exit(33)
	}
	
	// Данные шифруются с использованием AES-GCM, а результат включает nonce и шифротекст.
	return aesGCM.Seal(nonce, nonce, rawStr, nil)
}

// Метод дешифрования строки
func (crypt *Encrypter) Decrypt(cryptoStr []byte) []byte {

	// Создается AES блочный шифр.
	block, err := aes.NewCipher([]byte(crypt.CryptoKey))
	if err != nil {
		output.PrintError("Не удалось расшифровать данные аккаунтов " + err.Error())
		os.Exit(34)
	}
	
	// Устанавливается режим GCM.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		output.PrintError("Не удалось расшифровать данные аккаунтов " + err.Error())
		os.Exit(35)
	}
	
	/* Читается nonce из зашифрованной строки. 
	Остальная часть строки интерпретируется как шифротекст. 
	Используется nonce и шифротекст для восстановления исходных данных. */ 
	nonceSize := aesGCM.NonceSize()
	nonce, cypherText := cryptoStr[:nonceSize], cryptoStr[nonceSize:]
	originalText, err := aesGCM.Open(nil, nonce, cypherText, nil)
	if err != nil {
		output.PrintError("Не удалось расшифровать данные аккаунтов " + err.Error())
		os.Exit(36)
	}
	return originalText
}



