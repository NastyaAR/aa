package algs1

import (
	"lab1.com/matrix"
)

func GetMin(values ...int) int {
	// камель кейс
	cur_min := values[0]

	for _, val := range values {
		if val < cur_min {
			cur_min = val
		}
	}

	return cur_min
}

func ResursiveLeven(str1, str2 []rune) int {
	m, n := len(str1), len(str2)
	var operation int

	if m == 0 && n == 0 {
		return 0
	} else if m == 0 && n > 0 {
		return n
	} else if m > 0 && n == 0 {
		return m
	} else {
		if str1[m-1] != str2[n-1] {
			operation = 1
		}
		return GetMin(ResursiveLeven(str1, str2[0:n-1])+1,
			ResursiveLeven(str1[0:m-1], str2)+1,
			ResursiveLeven(str1[0:m-1], str2[0:n-1])+operation)
	}

	return 0
}

func MatrixLeven(str1 []rune, str2 []rune) (int, matrix.Matrix) {
	operation := 0
	matr := matrix.CreateMatrix(len(str1)+1, len(str2)+1, false)

	for i := 1; i < matr.M; i++ {
		for j := 1; j < matr.N; j++ {
			if str1[i-1] != str2[j-1] {
				operation = 1
			} else {
				operation = 0
			}
			matr.Matr[i][j] = GetMin(matr.Matr[i][j-1]+1, matr.Matr[i-1][j]+1, matr.Matr[i-1][j-1]+operation)
		}
	}

	return matr.Matr[matr.M-1][matr.N-1], matr
}

func MatrixDamerauLeven(str1 []rune, str2 []rune) (int, matrix.Matrix) {
	operation := 0
	matr := matrix.CreateMatrix(len(str1)+1, len(str2)+1, false)

	for i := 1; i < matr.M; i++ {
		for j := 1; j < matr.N; j++ {
			if str1[i-1] != str2[j-1] {
				operation = 1
			} else {
				operation = 0
			}
			if i >= 2 && j >= 2 && str1[i-1] == str2[j-2] && str1[i-2] == str2[j-1] {
				matr.Matr[i][j] = GetMin(matr.Matr[i][j-1]+1, matr.Matr[i-1][j]+1,
					matr.Matr[i-1][j-1]+operation, matr.Matr[i-2][j-2]+1)
			} else {
				matr.Matr[i][j] = GetMin(matr.Matr[i][j-1]+1, matr.Matr[i-1][j]+1, matr.Matr[i-1][j-1]+operation)
			}
		}
	}

	return matr.Matr[matr.M-1][matr.N-1], matr
}

func GetDistance(matr *matrix.Matrix, i int, j int, str1 []rune, str2 []rune) int {
	if matr.Matr[i][j] == -1 {
		matr.Matr[i][j] = RecursivePartOfLeven(str1, str2, matr)
	}
	return matr.Matr[i][j]
}

func RecursivePartOfLeven(str1 []rune, str2 []rune, matr *matrix.Matrix) int {
	var operation int

	m, n := len(str1), len(str2)

	if m == 0 && n == 0 {
		return 0
	} else if m == 0 && n > 0 {
		return n
	} else if m > 0 && n == 0 {
		return m
	} else {
		if str1[m-1] != str2[n-1] {
			operation = 1
		}
		moveLeft := GetDistance(matr, m, n-1, str1, str2[:n-1]) + 1
		moveRight := GetDistance(matr, m-1, n, str1[:m-1], str2) + 1
		moveDiag := GetDistance(matr, m-1, n-1, str1[:m-1], str2[:n-1]) + operation
		if m >= 2 && n >= 2 && str1[m-1] == str2[n-2] && str1[m-2] == str2[n-1] {
			transposition := GetDistance(matr, m-2, n-2, str1[:m-2], str2[:n-2]) + 1
			return GetMin(moveLeft, moveRight, moveDiag, transposition)
		}
		return GetMin(moveLeft, moveRight, moveDiag)
	}
}

func RecursiveDamerauLevenWithCache(str1 []rune, str2 []rune) (int, matrix.Matrix) {
	matr := matrix.CreateMatrix(len(str1)+1, len(str2)+1, true)
	ans := RecursivePartOfLeven(str1, str2, &matr)
	matr.Matr[matr.M-1][matr.N-1] = ans

	return ans, matr
}

func DamerauLeven(str1, str2 []rune) int {
	m, n := len(str1), len(str2)
	var operation, transposition int

	if m == 0 && n == 0 {
		return 0
	} else if m == 0 && n > 0 {
		return n
	} else if m > 0 && n == 0 {
		return m
	} else {
		if str1[m-1] != str2[n-1] {
			operation = 1
		}
		moveLeft := DamerauLeven(str1, str2[:n-1]) + 1
		moveRight := DamerauLeven(str1[:m-1], str2) + 1
		moveDiag := DamerauLeven(str1[:m-1], str2[:n-1]) + operation
		if m >= 2 && n >= 2 && str1[m-1] == str2[n-2] && str1[m-2] == str2[n-1] {
			transposition = DamerauLeven(str1[:m-2], str2[:n-2]) + 1
			return GetMin(moveLeft, moveRight, moveDiag, transposition)
		}
		return GetMin(moveDiag, moveRight, moveLeft)
	}

	return 0
}
