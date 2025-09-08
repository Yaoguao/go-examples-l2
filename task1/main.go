package main

import "fmt"

//	TASK 1
/*
Что выведет программа?

Объясните вывод программы.

package main

import "fmt"

func main() {
  a := [5]int{76, 77, 78, 79, 80}
  var b []int = a[1:4]
  fmt.Println(b)
}
*/

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	b := a[1:4]

	fmt.Println(b) // [77, 78, 79]
	// при взятие подслайса [a:b], вторым значением b - будет индекс b-1, a - включительно
}
