package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

//	TASK 8
/*
Получение точного времени (NTP)
Создать программу, печатающую точное текущее время с использованием NTP-сервера.

Реализовать проект как модуль Go.

Использовать библиотеку ntp для получения времени.

Программа должна выводить текущее время, полученное через NTP (Network Time Protocol).

Необходимо обрабатывать ошибки библиотеки: в случае ошибки вывести её текст в STDERR и вернуть ненулевой код выхода.

Код должен проходить проверки (vet и golint), т.е. быть написан идиоматически корректно.

* проверки проходят
➜  time-ntp git:(master) ✗ golint
➜  time-ntp git:(master) ✗ go vet
*/

func main() {
	timeNTP, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error by get time: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(timeNTP.Local().Format(time.DateTime))
}
