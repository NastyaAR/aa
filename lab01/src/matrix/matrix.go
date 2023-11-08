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

func (m *Matrix) MakeMatrix() {
	m.Matr = make([][]int, m.M)
	for i := range m.Matr {
		m.Matr[i] = make([]int, m.N)
	}
}

func (m *Matrix) InitMatrix() {
	for j := 0; j < m.N; j++ {
		m.Matr[0][j] = j
	}

	for i := 0; i < m.M; i++ {
		m.Matr[i][0] = i
	}

	if m.Inf {
		for i := 1; i < m.M; i++ {
			for j := 1; j < m.N; j++ {
				m.Matr[i][j] = -1
			}
		}
	}
}

func (m *Matrix) OutputMatrix() {
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			fmt.Printf("%5d ", m.Matr[i][j])
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
