package matrix

import (
	"fmt"
)

type Matrix struct {
	Matr [][]int
	M    int
	N    int
	Inf  bool
}

func (matr *Matrix) MakeMatrix() {
	matr.Matr = make([][]int, matr.M)
	for i := range matr.Matr {
		matr.Matr[i] = make([]int, matr.N)
	}
}

// receiver name (m)
func (matr *Matrix) InitMatrix() {
	for j := 0; j < matr.N; j++ {
		matr.Matr[0][j] = j
	}

	for i := 0; i < matr.M; i++ {
		matr.Matr[i][0] = i
	}

	if matr.Inf {
		for i := 1; i < matr.M; i++ {
			for j := 1; j < matr.N; j++ {
				matr.Matr[i][j] = -1
			}
		}
	}
}

func (matr *Matrix) OutputMatrix() {
	for i := 0; i < matr.M; i++ {
		for j := 0; j < matr.N; j++ {
			fmt.Printf("%5d ", matr.Matr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func CreateMatrix(m, n int, inf bool) Matrix {
	matr := Matrix{M: m, N: n, Inf: inf}
	matr.MakeMatrix()
	matr.InitMatrix()
	return matr
}
