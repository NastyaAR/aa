package io

import "fmt"

func InputString(num int) string {
	var str string
	fmt.Printf("Введите строку №%d: ", num)
	fmt.Scanln(&str)
	fmt.Println()
	return str
}

func OutputMenu() {
	fmt.Println("\nДобро пожаловать в главное меню!")
	fmt.Println(`
	1. Ввести строки
	2. Изменить строки
	3. Вычислить редакционное расстояние между строками алгоритмом Левенштейна (матрица)
	4. Вычислить редакционное расстояние между строками алгоритмом Дамерау-Левенштейна (рекурсия без кеша)
	5. Вычислить редакционное расстояние между строками алгоритмом Дамерау-Левенштейна (рекурсия с кешем)
	6. Вычислить редакционное расстояние между строками алгоритмом Дамерау-Левенштейна (матрица)
	7. Вывести результаты замеров времени
	8. Вывести меню
	0. Выход `)
}

func ChooseOption() int {
	choice := 0
	fmt.Println()
	fmt.Print("Выберите опцию: ")
	fmt.Scanln(&choice)
	return choice
}
