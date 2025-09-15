package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Что выведет программа?
//
//Объяснить работу конвейера с использованием select.

// Данная функция возращает канал на чтение, ну просто реализация паттерна генератор, который просто наполняет канал значениями
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

// Что-то типо Fat-in
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			//	оператор select сам по себе блокирующий, если не использовать default
			// То есть, он ждет если хоть в один из двух каналов придет значение, то тот case он и будет обрабатывать
			// После обработки одного из case, происходит выход из него
			select {
			case v, ok := <-a:
				// как только канал закроется, то ok станет false и мы обнуляем его
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				// после проверки и закрытия двух каналов, закрываем канал c
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	// for range итерируется по каналу до его закрытия
	for v := range c {
		fmt.Print(v) // Будет не последовательный вывод из двух каналов
	}
}
