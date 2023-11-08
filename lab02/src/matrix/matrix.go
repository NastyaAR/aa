package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Matrix struct {
	Matr [][]int
	M    int
	N    int
}

func NilMatrix() Matrix {
	var m Matrix
	m.Matr = nil
	m.M = 0
	m.N = 0
	return m
}

func (m *Matrix) MakeMatrix() {
	m.Matr = make([][]int, m.M)
	for i := range m.Matr {
		m.Matr[i] = make([]int, m.N)
	}
}

func (m *Matrix) FillMatrix() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			m.Matr[i][j] = rand.Intn(100)
		}
	}
}

func CreateMatrix(m int, n int, fill bool) Matrix {
	matrix := Matrix{M: m, N: n}
	matrix.MakeMatrix()
	if fill {
		matrix.FillMatrix()
	}

	return matrix
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

func AddMatrixes(m1, m2 *Matrix) Matrix {
	m := CreateMatrix(m1.M, m1.N, false)

	for i := 0; i < m1.M; i++ {
		for j := 0; j < m1.N; j++ {
			m.Matr[i][j] = m1.Matr[i][j] + m2.Matr[i][j]
		}
	}

	return m
}

func SubMatrixes(m1, m2 *Matrix) Matrix {
	m := CreateMatrix(m1.M, m1.N, false)

	for i := 0; i < m1.M; i++ {
		for j := 0; j < m1.N; j++ {
			m.Matr[i][j] = m1.Matr[i][j] - m2.Matr[i][j]
		}
	}

	return m
}

func (m *Matrix) completeMatrixToSq() {
	if m.M > m.N {
		for j := 0; j < m.M-m.N; j++ {
			for i := range m.Matr {
				m.Matr[i] = append(m.Matr[i], 0)
			}
		}
		m.N = m.M
	} else if m.N > m.M {
		for j := 0; j < m.N-m.M; j++ {
			m.Matr = append(m.Matr, make([]int, m.N))
		}
		m.M = m.N
	}
}

func (m *Matrix) completeMatrixToDeg(deg int) {
	for j := 0; j < deg-m.M; j++ {
		m.Matr = append(m.Matr, make([]int, deg))
	}

	for j := 0; j < deg-m.N; j++ {
		for i := range m.Matr {
			m.Matr[i] = append(m.Matr[i], 0)
		}
	}

	m.M = deg
	m.N = deg
}

func getDegSize(m, n int) int {
	k := int(math.Max(float64(m), float64(n)))
	if k == 0 {
		return 0
	}

	var i int
	for i = 1; i < k; i *= 2 {
	}

	return i
}

func CopyMatrix(matr *Matrix) Matrix {
	cm := CreateMatrix(matr.M, matr.N, false)

	for i := 0; i < matr.M; i++ {
		for j := 0; j < matr.N; j++ {
			cm.Matr[i][j] = matr.Matr[i][j]
		}
	}

	return cm
}

func CompleteMatrixes(m1, m2 *Matrix) {
	m1.completeMatrixToSq()
	m2.completeMatrixToSq()

	n := getDegSize(m1.M, m2.M)

	m1.completeMatrixToDeg(n)
	m2.completeMatrixToDeg(n)
}

func (m *Matrix) CorrectMatrix(k, n int) Matrix {
	res := CreateMatrix(k, n, false)

	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			res.Matr[i][j] = m.Matr[i][j]
		}
	}

	return res
}
