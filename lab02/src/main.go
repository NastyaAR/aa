package main

import (
	"fmt"
	"lab02/time_measure"

	"lab02/algs"
	"lab02/inp_out"
	"lab02/matrix"
)

const nAlgs int = 6

func main() {
	var matr1, matr2, res matrix.Matrix
	var err error
	var b, e, step int

	choice := 1

	inp_out.OutputMenu()

	for choice != 0 {
		choice, err = inp_out.ChooseOption()

		if err != nil || choice < 0 || choice > 11 {
			fmt.Print(err)
			continue
		}

		switch choice {
		case 1:
			matr1, matr2, err = inp_out.GenerateMatrixes()
			if err != nil {
				fmt.Print(err)
				continue
			}
		case 2:
			matr1, matr2, err = inp_out.InputMatrixes()
			if err != nil {
				fmt.Print(err)
				continue
			}
		case 3:
			fmt.Println("\nПервая матрица: ")
			matr1.OutputMatrix()
			fmt.Println("\nВторая матрица: ")
			matr2.OutputMatrix()
		case 4:
			res = algs.BaseMulMatrixNoOpt(&matr1, &matr2)
			fmt.Println("\nРезультат: ")
			res.OutputMatrix()
		case 5:
			res = algs.BaseMulMatrixOpt1(&matr1, &matr2)
			fmt.Println("\nРезультат: ")
			res.OutputMatrix()
		case 6:
			res = algs.GrapeMulMatrixNoOpt(&matr1, &matr2)
			fmt.Println("\nРезультат: ")
			res.OutputMatrix()
		case 7:
			res = algs.GrapeMulMatrixOpt1and2(&matr1, &matr2)
			fmt.Println("\nРезультат: ")
			res.OutputMatrix()
		case 8:
			cm1 := matrix.CopyMatrix(&matr1)
			cm2 := matrix.CopyMatrix(&matr2)
			matrix.CompleteMatrixes(&cm1, &cm2)
			res = algs.StrassenMulMatrixNoOpt(&cm1, &cm2)
			res = res.CorrectMatrix(matr1.M, matr2.N)
			fmt.Println("\nРезультат: ")
			res.OutputMatrix()
		case 9:
			cm1 := matrix.CopyMatrix(&matr1)
			cm2 := matrix.CopyMatrix(&matr2)
			matrix.CompleteMatrixes(&cm1, &cm2)
			res = algs.StrassenMulMatrixOpt1(&cm1, &cm2)
			res = res.CorrectMatrix(matr1.M, matr2.N)
			fmt.Println("\nРезультат: ")
			res.OutputMatrix()
		case 10:
			b, e, step = 11, 500, 50
			seconds, sizes := time_measure.MeasureTime(b, e, step)
			inp_out.OutputTime(seconds, nAlgs, len(sizes))
			inp_out.OutputTimeIntoFiles(seconds, nAlgs, sizes)
			inp_out.GetFullCsv(seconds, sizes)
		case 11:
			inp_out.OutputMenu()
		}
	}
}
