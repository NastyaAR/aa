package inp_out

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func inputNumber() (int, error) {
	var number int
	l, err := fmt.Scanf("%d", &number)
	if l < 1 && err != nil {
		return 0, errors.New("\nНекорректно введён размер массива!\n")
	}
	return number, nil
}

func inputElem() (int, error) {
	var num int
	l, err := fmt.Scanf("%d", &num)
	if l < 1 && err != nil {
		return 0, errors.New("\nНекорректно введён элемент массива!\n")
	}
	return num, nil
}

func InputArray() ([]int, error) {
	var err error
	m, err := inputNumber()

	if err != nil {
		return nil, err
	}

	fmt.Print("\nВведите элементы массива: ")

	arr := make([]int, m, m)
	for i := range arr {
		arr[i], err = inputElem()
		if err != nil {
			return nil, err
		}
	}

	return arr, nil
}

func GenerateArray(equal bool) ([]int, error) {
	var err error
	fmt.Print("\nВведите количество элементов массива: ")
	number, err := inputNumber()

	if err != nil {
		return nil, err
	}

	num := rand.Intn(1000)

	arr := make([]int, number, number)
	for i := range arr {
		if equal {
			arr[i] = num
		} else {
			flag := rand.Intn(10)
			if flag > 5 {
				arr[i] = rand.Intn(1000) * -1
			} else {
				arr[i] = rand.Intn(1000)
			}
		}
	}

	return arr, nil
}

func GenerateArrayForTest(equal bool, number int) []int {
	num := rand.Intn(1000)

	arr := make([]int, number, number)
	for i := range arr {
		if equal {
			arr[i] = num
		} else {
			flag := rand.Intn(10)
			if flag > 5 {
				arr[i] = rand.Intn(1000) * -1
			} else {
				arr[i] = rand.Intn(1000)
			}
		}
	}

	return arr
}

func OutputMenu() {
	fmt.Println("\nДобро пожаловать в главное меню!")
	fmt.Println(`
	1. Ввести размерность массива и заполнить его случайными числами (неотсорт.)
	2. Ввести размерность массива и заполнить его случайным числом
	3. Уже отсортированный массив, заполненный случайными числами
	4. Отсортированный в обратном порядке массив, заполненный случайными числами
	5. Ввести массив вручную
	6. Вывести массив
	
	Отсортировать массив:
	7. блочной сортировкой
	8. поразрядной сортировкой
	9. сортировкой слиянием

	10. Замеры времени
	
	11. Вывести меню
	0. Выход `)
}

func ChooseOption() (int, error) {
	choice := 0
	fmt.Print("\nВыберите опцию: ")
	l, err := fmt.Scanln(&choice)
	if l < 1 && err != nil {
		return -1, errors.New("\nНекорректно введена опция!\n")
	}
	return choice, nil
}

func OutputTime(seconds [][]float32, m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%8.7f   ", seconds[i][j])
		}
		fmt.Println()
	}
}

func outputTimeIntoFile(dir string, fname string, sizes []int, seconds []float32) {
	file, err := os.Create(fmt.Sprintf("../report/inc/csv/%s/%s.csv", dir, fname))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file, "len,%s\n", fname)
	for i := range sizes {
		fmt.Fprintf(file, "%d,%8.7f\n", sizes[i], seconds[i])
	}
}

func OutputTimeIntoFiles(seconds [][][]float32, sizes []int) {
	algs := []string{"Блочная_сортировка",
		"Порязрядная_сортировка",
		"Сортировка_слиянием",
	}

	typesData := []string{"Случайные",
		"Одинаковые",
		"Случ_отсорт",
		"Случ_отсорт_обр",
	}

	for i := range typesData {
		for j := range algs {
			outputTimeIntoFile(typesData[i], algs[j], sizes, seconds[i][j])
		}
	}
}

func getFullCsv(fname string, seconds [][]float32, sizes []int) {
	file, err := os.Create(fmt.Sprintf("../report/inc/csv/%s.csv", fname))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for i := range sizes {
		fmt.Fprintf(file, "%d", sizes[i])
		for j := range seconds {
			fmt.Fprintf(file, ",%8.7f", seconds[j][i])
		}
		fmt.Fprintln(file)
	}
}

func GetFullCsvs(seconds [][][]float32, sizes []int) {
	typesData := []string{"Случайные",
		"Одинаковые",
		"Случ_отсорт",
		"Случ_отсорт_обр",
	}

	for i := range typesData {
		getFullCsv(typesData[i], seconds[i], sizes)
	}
}
