package examples

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// Пример работы с рунами
func runeExample() {
	/*
	Руна (rune) - это не что иное, как представление в виде unicode символов строки, по своей сути является int32.
	*/
	hello := "Hello world!)"
	for _, char := range hello {
		fmt.Println(char, string(char))
	}

	/*
	Мы можем также увидеть, что rune - это просто alias к int32. Давай сделаем массив rune из все той же приветственной строки. Если навестить на rune и узнать, что она из себя представляет, то увидим, что ее определение выглядит следующим образом:

	type rune = int32

	То есть я могу создать свою руну, как отдельный alias данного типа, выполнив, например, такой код:

	type myRune = int32

	Соответственно и в коде ниже если создать не массив rune, а массив int32, то ничего не изменится, эти записи эквивалентны.

	helloArr := []int32(hello)
	*/
	helloArr := []rune(hello)
	for _, char := range helloArr {
		fmt.Println(char, string(char))
	}
}


/*
	Упражнение "генерация пароля".
	Задача написать функцию, которая примет на вход целое число - количество символов в случайном пароле и сгенерирует этот самый пароль, вернув строку. Для этого понадобится импортировать "math/rand/v2"
*/

func generateRandPassword(passLen int) string {

	// Создаем слайс длиной, равной заданной длине случайного пароля
	randPass := make([]string, passLen)

	// Выполняем кол-во итераций, равное заданной длине случайного пароля
	for indx:=0; indx < int(passLen); indx++ {

		// На каждой итерации генерируем случайное число от 0 до 65535 (uint16)
		randElem := rand.IntN(65535)
		// И ставим его по нужному индексу в слайс случайного пароля, преобразовав в строку
		randPass[indx] = string(randElem)
	}

	// Возвращаем сгенерированный пароль, объединив слайс с символами пароля в одну строку
	return strings.Join(randPass, "")
}
