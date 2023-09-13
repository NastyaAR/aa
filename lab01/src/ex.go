package main

import (
	"fmt"

	"lab1.com/algs1"
	"lab1.com/io"
	"lab1.com/matrix"
)

const (
	SUCCESS        = 0
	INP_STR_ERR    = -1
	INP_OPTION_ERR = -2
)

func ImputStringsOption() ([]rune, []rune, int) {
	string1 := io.InputString(1)
	string2 := io.InputString(2)
	rstring1, rstring2 := []rune(string1), []rune(string2)
	if string1 == "" || string2 == "" {
		fmt.Println("\nОшибка ввода строк!")
		return rstring1, rstring2, INP_STR_ERR
	}

	return rstring1, rstring2, SUCCESS
}

func ChangeStringsOption(rstr1, rstr2 []rune) ([]rune, []rune, int) {
	var choice int
	fmt.Print("Введите 1, если хотите изменить первую строку, 2 - вторую: ")
	fmt.Scanln(&choice)
	if choice != 1 && choice != 2 {
		fmt.Println("\nОшибка выбора!")
		return nil, nil, INP_OPTION_ERR
	}

	str := io.InputString(choice)
	if str == "" {
		fmt.Println("\nОшибка ввода строки!")
		return nil, nil, INP_STR_ERR
	}

	if choice == 1 {
		rstr1 = []rune(str)
	} else {
		rstr2 = []rune(str)
	}

	return rstr1, rstr2, SUCCESS
}

func main() {
	choice := 1
	var str1, str2 []rune
	var ok int
	var answers [4]int

	io.OutputMenu()

	for choice != 0 {
		choice = io.ChooseOption()

		if choice < 0 || choice > 7 {
			fmt.Println("\nОшибка выбора опции!")
			continue
		}

		switch choice {
		case 1:
			str1, str2, ok = ImputStringsOption()
			if ok < 0 {
				continue
			}
		case 2:
			str1, str2, ok = ChangeStringsOption(str1, str2)
			if ok < 0 {
				continue
			}
		case 3:
			answers[0] = algs1.ResursiveLeven(str1, str2)
			fmt.Printf("\nРедакционное расстояние между строками = %d\n", answers[0])
		case 4:
			answers[1] = algs1.DamerauLeven(str1, str2)
			fmt.Printf("\nРедакционное расстояние между строками = %d\n", answers[1])
		case 5:
			matr1 := matrix.CreateMatrix(len(str1)+1, len(str2)+1, false)
			answers[2] = algs1.MatrixDamerauLeven(str1, str2, &matr1)
			fmt.Printf("\nРедакционное расстояние между строками = %d\n", answers[2])
			fmt.Println("Матрица:\n")
			matr1.OutputMatrix()
		case 6:
			matr2 := matrix.CreateMatrix(len(str1)+1, len(str2)+1, true)
			answers[3] = algs1.RecursiveDamerauLevenWithCache(str1, str2, &matr2)
			fmt.Printf("\nРедакционное расстояние между строками = %d\n", answers[3])
			fmt.Println("Матрица:\n")
			matr2.OutputMatrix()
		case 7:
			io.OutputMenu()
		}
	}
}