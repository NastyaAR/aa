package algs

import (
	"lab02/matrix"
)

func BaseMulMatrixNoOpt(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			for k := 0; k < matrix1.N; k++ {
				result.Matr[i][j] = result.Matr[i][j] + matrix1.Matr[i][k]*matrix2.Matr[k][j]
			}
		}
	}

	return result
}

func BaseMulMatrixOpt1(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			for k := 0; k < matrix1.N; k++ {
				result.Matr[i][j] += matrix1.Matr[i][k] * matrix2.Matr[k][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixNoOpt(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)
	finish_cols := matrix1.N / 2
	finish_rows := matrix2.M / 2

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < finish_cols; j++ {
			rowFactor[i] = rowFactor[i] + matrix1.Matr[i][2*j]*matrix1.Matr[i][2*j+1]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 0; j < finish_rows; j++ {
			columnFactor[i] = columnFactor[i] + matrix2.Matr[2*j][i]*matrix2.Matr[2*j+1][i]
		}
	}

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			result.Matr[i][j] = -1*rowFactor[i] - columnFactor[j]
			for k := 0; k < finish_cols; k++ {
				result.Matr[i][j] = result.Matr[i][j] + (matrix1.Matr[i][2*k]+
					matrix2.Matr[2*k+1][j])*(matrix1.Matr[i][2*k+1]+matrix2.Matr[2*k][j])
			}
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] = result.Matr[i][j] + matrix1.Matr[i][matrix1.N-1]*matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt1(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)
	finish_cols := matrix1.N / 2
	finish_rows := matrix2.M / 2

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < finish_cols; j++ {
			rowFactor[i] += matrix1.Matr[i][2*j] * matrix1.Matr[i][2*j+1]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 0; j < finish_rows; j++ {
			columnFactor[i] += matrix2.Matr[2*j][i] * matrix2.Matr[2*j+1][i]
		}
	}

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			result.Matr[i][j] = -1*rowFactor[i] - columnFactor[j]
			for k := 0; k < finish_cols; k++ {
				result.Matr[i][j] += (matrix1.Matr[i][2*k] + matrix2.Matr[2*k+1][j]) * (matrix1.Matr[i][2*k+1] + matrix2.Matr[2*k][j])
			}
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] += matrix1.Matr[i][matrix1.N-1] * matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt2(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)
	var finish_cols int = matrix1.N / 2
	var finish_rows int = matrix2.M / 2

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < finish_cols; j++ {
			rowFactor[i] = rowFactor[i] + matrix1.Matr[i][j<<1]*matrix1.Matr[i][j<<1+1]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 0; j < finish_rows; j++ {
			columnFactor[i] = columnFactor[i] + matrix2.Matr[j<<1][i]*matrix2.Matr[j<<1+1][i]
		}
	}

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			result.Matr[i][j] = -1*rowFactor[i] - columnFactor[j]
			for k := 0; k < finish_cols; k++ {
				result.Matr[i][j] = result.Matr[i][j] + (matrix1.Matr[i][k<<1]+matrix2.Matr[k<<1+1][j])*(matrix1.Matr[i][k<<1+1]+matrix2.Matr[k<<1][j])
			}
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] = result.Matr[i][j] + matrix1.Matr[i][matrix1.N-1]*matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt3(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)
	finish_cols := matrix1.N / 2
	finish_rows := matrix2.M / 2

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < finish_cols; j++ {
			rowFactor[i] = rowFactor[i] + matrix1.Matr[i][2*j]*matrix1.Matr[i][2*j+1]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 0; j < finish_rows; j++ {
			columnFactor[i] = columnFactor[i] + matrix2.Matr[2*j][i]*matrix2.Matr[2*j+1][i]
		}
	}

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			result.Matr[i][j] = -1*rowFactor[i] - columnFactor[j]
			for k := 1; k < matrix1.N; k += 2 {
				result.Matr[i][j] = result.Matr[i][j] + (matrix1.Matr[i][k-1]+matrix2.Matr[k][j])*(matrix1.Matr[i][k]+matrix2.Matr[k-1][j])
			}
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] = result.Matr[i][j] + matrix1.Matr[i][matrix1.N-1]*matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt1and2(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)
	finish_cols := matrix1.N >> 1
	finish_rows := matrix2.M >> 1

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < finish_cols; j++ {
			rowFactor[i] += matrix1.Matr[i][j<<1] * matrix1.Matr[i][j<<1+1]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 0; j < finish_rows; j++ {
			columnFactor[i] += matrix2.Matr[j<<1][i] * matrix2.Matr[j<<1+1][i]
		}
	}

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			result.Matr[i][j] = -1*rowFactor[i] - columnFactor[j]
			for k := 0; k < finish_cols; k++ {
				result.Matr[i][j] += (matrix1.Matr[i][k<<1] + matrix2.Matr[k<<1+1][j]) *
					(matrix1.Matr[i][k<<1+1] + matrix2.Matr[k<<1][j])
			}
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] += matrix1.Matr[i][matrix1.N-1] *
					matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt1and3(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 1; j < matrix1.N; j += 2 {
			rowFactor[i] += matrix1.Matr[i][j-1] * matrix1.Matr[i][j]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 1; j < matrix2.M; j += 2 {
			columnFactor[i] += matrix2.Matr[j-1][i] * matrix2.Matr[j][i]
		}
	}

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			result.Matr[i][j] = -1*rowFactor[i] - columnFactor[j]
			for k := 1; k < matrix1.N; k += 2 {
				result.Matr[i][j] += (matrix1.Matr[i][k-1] + matrix2.Matr[k][j]) * (matrix1.Matr[i][k] + matrix2.Matr[k-1][j])
			}
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] += matrix1.Matr[i][matrix1.N-1] * matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt1and3WithBuffer(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 1; j < matrix1.N; j += 2 {
			rowFactor[i] += matrix1.Matr[i][j-1] * matrix1.Matr[i][j]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 1; j < matrix2.M; j += 2 {
			columnFactor[i] += matrix2.Matr[j-1][i] * matrix2.Matr[j][i]
		}
	}

	var buffer int

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			buffer = -1*rowFactor[i] - columnFactor[j]
			for k := 1; k < matrix1.N; k += 2 {
				buffer += (matrix1.Matr[i][k-1] + matrix2.Matr[k][j]) * (matrix1.Matr[i][k] + matrix2.Matr[k-1][j])
			}
			result.Matr[i][j] += buffer
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] += matrix1.Matr[i][matrix1.N-1] * matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func GrapeMulMatrixOpt4(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	result := matrix.CreateMatrix(matrix1.M, matrix2.N, false)

	finish_cols := matrix1.N / 2
	finish_rows := matrix2.M / 2

	rowFactor := make([]int, matrix1.M, matrix1.M)
	columnFactor := make([]int, matrix2.N, matrix2.N)

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < finish_cols; j++ {
			rowFactor[i] = rowFactor[i] + matrix1.Matr[i][2*j]*matrix1.Matr[i][2*j+1]
		}
	}

	for i := 0; i < matrix2.N; i++ {
		for j := 0; j < finish_rows; j++ {
			columnFactor[i] = columnFactor[i] + matrix2.Matr[2*j][i]*matrix2.Matr[2*j+1][i]
		}
	}

	var buffer int

	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			buffer = -1*rowFactor[i] - columnFactor[j]
			for k := 0; k < finish_cols; k++ {
				buffer = buffer + (matrix1.Matr[i][2*k]+matrix2.Matr[2*k+1][j])*(matrix1.Matr[i][2*k+1]+matrix2.Matr[2*k][j])
			}
			result.Matr[i][j] += buffer
		}
	}

	if matrix1.N%2 == 1 {
		for i := 0; i < matrix1.M; i++ {
			for j := 0; j < matrix2.N; j++ {
				result.Matr[i][j] += matrix1.Matr[i][matrix1.N-1] * matrix2.Matr[matrix1.N-1][j]
			}
		}
	}

	return result
}

func StrassenMulMatrixNoOpt(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	n := matrix1.M

	if n == 1 {
		result := matrix.CreateMatrix(1, 1, false)
		result.Matr[0][0] = matrix1.Matr[0][0] * matrix2.Matr[0][0]
		return result
	}

	a11 := matrix.CreateMatrix(n/2, n/2, false)
	a12 := matrix.CreateMatrix(n/2, n/2, false)
	a21 := matrix.CreateMatrix(n/2, n/2, false)
	a22 := matrix.CreateMatrix(n/2, n/2, false)
	b11 := matrix.CreateMatrix(n/2, n/2, false)
	b12 := matrix.CreateMatrix(n/2, n/2, false)
	b21 := matrix.CreateMatrix(n/2, n/2, false)
	b22 := matrix.CreateMatrix(n/2, n/2, false)

	for i := 0; i < n/2; i++ {
		a11.Matr[i] = matrix1.Matr[i][:n/2]
		a12.Matr[i] = matrix1.Matr[i][n/2:]
		a21.Matr[i] = matrix1.Matr[i+n/2][:n/2]
		a22.Matr[i] = matrix1.Matr[i+n/2][n/2:]
		b11.Matr[i] = matrix2.Matr[i][:n/2]
		b12.Matr[i] = matrix2.Matr[i][n/2:]
		b21.Matr[i] = matrix2.Matr[i+n/2][:n/2]
		b22.Matr[i] = matrix2.Matr[i+n/2][n/2:]
	}

	p1 := matrix.AddMatrixes(&a11, &a22)
	p2 := matrix.AddMatrixes(&b11, &b22)
	p3 := matrix.SubMatrixes(&a12, &a22)
	p4 := matrix.AddMatrixes(&b11, &b12)
	p5 := matrix.AddMatrixes(&a11, &a12)
	p6 := matrix.AddMatrixes(&a21, &a22)
	p7 := matrix.SubMatrixes(&b21, &b11)
	p8 := matrix.SubMatrixes(&b12, &b22)
	p9 := matrix.SubMatrixes(&a21, &a11)
	p10 := matrix.AddMatrixes(&b21, &b22)

	k1 := StrassenMulMatrixNoOpt(&p1, &p2)
	k2 := StrassenMulMatrixNoOpt(&p3, &p10)
	k3 := StrassenMulMatrixNoOpt(&p9, &p4)
	k7 := StrassenMulMatrixNoOpt(&p5, &b22)
	k4 := StrassenMulMatrixNoOpt(&p6, &b11)
	k5 := StrassenMulMatrixNoOpt(&a22, &p7)
	k6 := StrassenMulMatrixNoOpt(&a11, &p8)

	c11 := matrix.AddMatrixes(&k1, &k2)
	c12 := matrix.SubMatrixes(&k5, &k7)
	r11 := matrix.AddMatrixes(&c11, &c12)

	r12 := matrix.AddMatrixes(&k6, &k7)

	r21 := matrix.AddMatrixes(&k5, &k4)

	c22 := matrix.AddMatrixes(&k1, &k3)
	c23 := matrix.SubMatrixes(&k6, &k4)

	r22 := matrix.AddMatrixes(&c22, &c23)

	result := matrix.CreateMatrix(n, n, false)

	for i := 0; i < n/2; i++ {
		result.Matr[i] = append(r11.Matr[i], r12.Matr[i]...)
	}
	for i := n / 2; i < n; i++ {
		result.Matr[i] = append(r21.Matr[i-n/2], r22.Matr[i-n/2]...)
	}

	return result
}

func StrassenMulMatrixOpt1(matrix1, matrix2 *matrix.Matrix) matrix.Matrix {
	n := matrix1.M

	if n == 1 {
		result := matrix.CreateMatrix(1, 1, false)
		result.Matr[0][0] = matrix1.Matr[0][0] * matrix2.Matr[0][0]
		return result
	}

	a11 := matrix.CreateMatrix(n>>1, n>>1, false)
	a12 := matrix.CreateMatrix(n>>1, n>>1, false)
	a21 := matrix.CreateMatrix(n>>1, n>>1, false)
	a22 := matrix.CreateMatrix(n>>1, n>>1, false)
	b11 := matrix.CreateMatrix(n>>1, n>>1, false)
	b12 := matrix.CreateMatrix(n>>1, n>>1, false)
	b21 := matrix.CreateMatrix(n>>1, n>>1, false)
	b22 := matrix.CreateMatrix(n>>1, n>>1, false)

	for i := 0; i < n>>1; i++ {
		a11.Matr[i] = matrix1.Matr[i][:n>>1]
		a12.Matr[i] = matrix1.Matr[i][n>>1:]
		a21.Matr[i] = matrix1.Matr[i+n>>1][:n>>1]
		a22.Matr[i] = matrix1.Matr[i+n>>1][n>>1:]
		b11.Matr[i] = matrix2.Matr[i][:n>>1]
		b12.Matr[i] = matrix2.Matr[i][n>>1:]
		b21.Matr[i] = matrix2.Matr[i+n>>1][:n>>1]
		b22.Matr[i] = matrix2.Matr[i+n>>1][n>>1:]
	}

	s1 := matrix.AddMatrixes(&a21, &a22)
	s2 := matrix.SubMatrixes(&s1, &a11)
	s3 := matrix.SubMatrixes(&a11, &a21)
	s4 := matrix.SubMatrixes(&a12, &s2)
	s5 := matrix.SubMatrixes(&b12, &b11)
	s6 := matrix.SubMatrixes(&b22, &s5)
	s7 := matrix.SubMatrixes(&b22, &b12)
	s8 := matrix.SubMatrixes(&s6, &b21)

	p1 := StrassenMulMatrixOpt1(&s2, &s6)
	p2 := StrassenMulMatrixOpt1(&a11, &b11)
	p3 := StrassenMulMatrixOpt1(&a12, &b21)
	p4 := StrassenMulMatrixOpt1(&s3, &s7)
	p5 := StrassenMulMatrixOpt1(&s1, &s5)
	p6 := StrassenMulMatrixOpt1(&s4, &b22)
	p7 := StrassenMulMatrixOpt1(&a22, &s8)

	t1 := matrix.AddMatrixes(&p1, &p2)
	t2 := matrix.AddMatrixes(&t1, &p4)

	r11 := matrix.AddMatrixes(&p2, &p3)
	r121 := matrix.AddMatrixes(&t1, &p5)
	r12 := matrix.AddMatrixes(&r121, &p6)
	r21 := matrix.SubMatrixes(&t2, &p7)
	r22 := matrix.AddMatrixes(&t2, &p5)

	result := matrix.CreateMatrix(n, n, false)

	for i := 0; i < n>>1; i++ {
		result.Matr[i] = append(r11.Matr[i], r12.Matr[i]...)
	}
	for i := n >> 1; i < n; i++ {
		result.Matr[i] = append(r21.Matr[i-n>>1], r22.Matr[i-n>>1]...)
	}

	return result
}
