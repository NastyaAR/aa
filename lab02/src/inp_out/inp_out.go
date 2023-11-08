package inp_out

import (
	"errors"
	"fmt"
	"lab02/matrix"
	"os"
)

func inputSizes3x() (int, int, int, error) {
	var m, n, k int

	fmt.Printf("Введите количество строк и столбцов для 1 матрицы и количество столбцов для 2й: ")
	l, err := fmt.Scanln(&m, &n, &k)
	if l < 3 && err != nil {
		return 0, 0, 0, errors.New("\nНекорректно введены размерности матриц!\n")
	}

	return m, n, k, nil
}

func inputSizes2x() (int, int, error) {
	var m, n int

	fmt.Printf("Введите количество строк и столбцов для матрицы: ")
	l, err := fmt.Scanln(&m, &n)
	if l < 2 && err != nil {
		return 0, 0, errors.New("\nНекорректно введены размерности матрицы!\n")
	}

	return m, n, nil
}

func inputElem() (int, error) {
	var num int
	l, err := fmt.Scanf("%d", &num)
	if l < 1 && err != nil {
		return 0, errors.New("\nНекорректно введён элемент матрицы!\n")
	}
	return num, nil
}

func inputMatrix(prev_col int) (matrix.Matrix, error) {
	var err error
	m, n, err := inputSizes2x()

	if err != nil {
		return matrix.NilMatrix(), err
	}

	if prev_col != 0 && prev_col != m {
		return matrix.NilMatrix(), errors.New("\nНекорректный размер второй матрицы!\n")
	}

	mtr := matrix.CreateMatrix(m, n, false)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mtr.Matr[i][j], err = inputElem()
			if err != nil {
				return matrix.NilMatrix(), err
			}
		}
	}

	return mtr, nil
}

func InputMatrixes() (matrix.Matrix, matrix.Matrix, error) {
	var mtr1, mtr2 matrix.Matrix
	var err error
	mtr1, err = inputMatrix(0)

	if err != nil {
		return matrix.NilMatrix(), matrix.NilMatrix(), err
	}

	mtr2, err = inputMatrix(mtr1.N)
	if err != nil {
		return matrix.NilMatrix(), matrix.NilMatrix(), err
	}

	return mtr1, mtr2, nil
}

func GenerateMatrixes() (matrix.Matrix, matrix.Matrix, error) {
	var err error
	m, n, k, err := inputSizes3x()

	if err != nil {
		return matrix.NilMatrix(), matrix.NilMatrix(), err
	}

	mtr1 := matrix.CreateMatrix(m, n, true)
	mtr2 := matrix.CreateMatrix(n, k, true)

	return mtr1, mtr2, nil
}

func GenerateSquareMatrixes(m int) (matrix.Matrix, matrix.Matrix) {
	mtr1 := matrix.CreateMatrix(m, m, true)
	mtr2 := matrix.CreateMatrix(m, m, true)

	return mtr1, mtr2
}

func OutputMenu() {
	fmt.Println("\nДобро пожаловать в главное меню!")
	fmt.Println(`
	1. Ввести размерности и сгенерировать матрицы, заполненные случайными числами
	2. Ввести матрицы вручную
	3. Вывести матрицы
	
	Вычислить произведение матриц:
	4. классическим алгоритмом умножения матриц без оптимизаций
	5. классическим алгоритмом умножения матриц без оптимизаций с +=
	6. алгоритмом Винограда без оптимизаций
	7. алгоритмом Винограда с += и с << вместо *2
	8. алгоритмом Штрассена без оптимизаций
	9. алгоритмом Штрассена с >> вместо /2 и предвычислением некоторых слагаемых
	
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

func outputTimeIntoFile(fname string, sizes []int, seconds []float32) {
	file, err := os.Create(fmt.Sprintf("../report/inc/csv/%s.csv", fname))
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

func OutputTimeIntoFiles(seconds [][]float32, m int, sizes []int) {
	algs := []string{"Classic",
		"OptClassic",
		"Grape",
		"OptGrape",
		"Strassen",
		"OptStrassen",
	}

	for i := range seconds {
		outputTimeIntoFile(algs[i], sizes, seconds[i])
	}
}

func GetFullCsv(seconds [][]float32, sizes []int) {
	file, err := os.Create("../report/inc/csv/time.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for i := range sizes {
		fmt.Fprintf(file, "%d x %d", sizes[i])
		for j := range seconds {
			fmt.Fprintf(file, ",%8.7f", seconds[j][i])
		}
		fmt.Fprintln(file)
	}
}
