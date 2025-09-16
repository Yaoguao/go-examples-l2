package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//	TASK 9
/*
Распаковка строки
Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.

Примеры работы функции:

Вход: "a4bc2d5e"
Выход: "aaaabccddddde"

Вход: "abcd"
Выход: "abcd" (нет цифр — ничего не меняется)

Вход: "45"
Выход: "" (некорректная строка, т.к. в строке только цифры — функция должна вернуть ошибку)

Вход: ""
Выход: "" (пустая строка -> пустая строка)

Дополнительное задание
Поддерживать escape-последовательности вида \:

Вход: "qwe\4\5"
Выход: "qwe45" (4 и 5 не трактуются как числа, т.к. экранированы)

Вход: "qwe\45"
Выход: "qwe44444" (\4 экранирует 4, поэтому распаковывается только 5)

Требования к реализации
Функция должна корректно обрабатывать ошибочные случаи (возвращать ошибку, например, через error), и проходить unit-тесты.

Код должен быть статически анализируем (vet, golint).
*/

type cSymbol struct {
	s          rune
	multiplier []rune
}

// UnpackingString decompresses a string containing duplicate characters
func UnpackingString(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	if !strings.ContainsAny(str, "0123456789") {
		return str, nil
	}
	rStr := []rune(str)

	if unicode.IsDigit(rStr[0]) {
		return "", errors.New("incorrect format")
	}

	multSym := make([]cSymbol, 0, len(rStr))

	fPass := false
	for _, s := range str {
		if unicode.IsDigit(s) && !fPass {
			idx := len(multSym) - 1
			multSym[idx].multiplier = append(multSym[idx].multiplier, s)
		} else if string(s) == "\\" {
			fPass = true
			continue
		} else {
			fPass = false
			multSym = append(multSym, cSymbol{
				s:          s,
				multiplier: make([]rune, 0),
			})
		}
	}

	res := make([]rune, 0, len(rStr))
	for _, val := range multSym {
		strSize := string(val.multiplier)
		var size int

		if strSize == "" {
			size = 1
		} else {
			size, _ = strconv.Atoi(strSize)
		}
		s := val.s
		for i := 0; i < size; i++ {
			res = append(res, s)
		}
	}

	return string(res), nil
}

func main() {
	str := "qwe\\45"

	us, err := UnpackingString(str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(us)
}
