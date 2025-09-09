package main

import "fmt"

//	TASK 2
/*
Что выведет программа?

Объяснить порядок выполнения defer функций и итоговый вывод.
*/

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

// Оператор defer выполняется после return
// т.к в функции test() именнованное возвращаемое значение, то при выходе defer меняет его
// ну в anotherTest(), в defer меняется локальная переменная, после возврата значения
func main() {
	fmt.Println(test())        // x = 2
	fmt.Println(anotherTest()) // x = 1
}
