package main

import (
	"fmt"
	"sort"
)

//	TASK 11
/*
Поиск анаграмм в словаре
Напишите функцию, которая находит все множества анаграмм по заданному словарю.

Требования
На вход подается срез строк (слов на русском языке в Unicode).

На выходе: map-множество -> список, где ключом является первое встреченное слово множества, а значением — срез из всех слов, принадлежащих этому множеству анаграмм, отсортированных по возрастанию.

Множества из одного слова не должны выводиться (т.е. если нет анаграмм, слово игнорируется).

Все слова нужно привести к нижнему регистру.

Пример:

Вход: ["пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"]
Результат (ключи в примере могут быть в другом порядке):
– "пятак": ["пятак", "пятка", "тяпка"]
– "листок": ["листок", "слиток", "столик"]

Слово «стол» отсутствует в результатах, так как не имеет анаграмм.

Для решения задачи потребуется умение работать со строками, сортировать
и использовать структуры данных (map).

Оценим эффективность: решение должно работать за линейно-логарифмическое время относительно количества слов (допустимо n * m log m, где m — средняя длина слова для сортировки букв).
*/

func SearchAnagram(input []string) map[string][]string {
	groups := make(map[string][]string)
	first := make(map[string]string)

	for _, w := range input {
		sorted := sortStringUnicode(w)

		if _, ok := first[sorted]; !ok {
			first[sorted] = w
		}

		already := false
		for _, v := range groups[sorted] {
			if v == w {
				already = true
				break
			}
		}
		if !already {
			groups[sorted] = append(groups[sorted], w)
		}
	}

	res := make(map[string][]string)
	for s, words := range groups {
		if len(words) < 2 {
			continue
		}
		sort.Strings(words)
		res[first[s]] = words
	}

	return res
}

func sortStringUnicode(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func main() {
	str := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	fmt.Println(SearchAnagram(str))

}
