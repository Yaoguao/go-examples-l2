package main

//	TASK 5
/*
Что выведет программа?

Объяснить вывод программы.

*/

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

// есть схожесть с прошлой задачей(task 4)
// сам метод тест возращает указатель
// и когда он присваивается интерфейсу error, то интерфейс теперь имеет в себе информацию о типе(*customError) и конкретном значении(nil)
// И получается что интерфейс не является nil
func main() {
	var err error
	err = test()
	if err != nil {
		println("error") // Будет ERROR
		return
	}
	//err := test()
	//if err != nil {
	//	println("error") // не будет ERROR
	//	return
	//}
	println("ok")
}
