package time_measure

import (
	"fmt"
	"lab02/algs"
	"lab02/inp_out"
	"lab02/matrix"

	"math"
	"syscall"
)

const N int = 500
const nAlgs int = 6

func GetCPU() int64 {
	usage := new(syscall.Rusage)
	syscall.Getrusage(syscall.RUSAGE_SELF, usage)
	return usage.Utime.Nano() + usage.Stime.Nano()
}

func grapeMulMatrixNoOptTimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixNoOpt(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func grapeMulMatrixOpt1TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixOpt1(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func grapeMulMatrixOpt2TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixOpt2(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func grapeMulMatrixOpt3TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixOpt3(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func grapeMulMatrixOpt1and2TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixOpt1and2(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func grapeMulMatrixOpt1and3TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixOpt1and3(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func grapeMulMatrixOpt1and3BuffTimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.GrapeMulMatrixOpt1and3WithBuffer(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func baseMulMatrixNoOptTimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.BaseMulMatrixNoOpt(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func baseMulMatrixOpt1TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.BaseMulMatrixOpt1(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func strassenMulMatrixNoOptTimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	matrix.CompleteMatrixes(matrix1, matrix2)

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.StrassenMulMatrixNoOpt(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func strassenMulMatrixOpt1TimeMeasurement(matrix1, matrix2 *matrix.Matrix) (float32, matrix.Matrix) {
	var sum float32

	var startTime, finishTime int64
	var mtr matrix.Matrix

	matrix.CompleteMatrixes(matrix1, matrix2)

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		mtr = algs.StrassenMulMatrixOpt1(matrix1, matrix2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	mtr.Matr[0][0] *= 1

	return (sum / float32(N)) / 1e+9, mtr
}

func MeasureTime_v2(begin, end int) ([][]float32, []int) {
	seconds := make([][]float32, 6)

	number := 0
	for n := end; n >= begin; n /= 2 {
		number += 1
	}

	sizes := make([]int, number)
	var matr1, matr2, res matrix.Matrix

	func_pointers := [6]func(*matrix.Matrix, *matrix.Matrix) (float32, matrix.Matrix){
		baseMulMatrixNoOptTimeMeasurement,
		baseMulMatrixOpt1TimeMeasurement,
		grapeMulMatrixNoOptTimeMeasurement,
		grapeMulMatrixOpt1and2TimeMeasurement,
		strassenMulMatrixNoOptTimeMeasurement,
		strassenMulMatrixOpt1TimeMeasurement,
	}

	for i := 0; i < 6; i++ {
		seconds[i] = make([]float32, number)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < number; j++ {
			matr1, matr2 = inp_out.GenerateSquareMatrixes(begin * int(math.Pow(2, float64(j))))
			seconds[i][j], res = func_pointers[i](&matr1, &matr2)
			sizes[j] = begin * int(math.Pow(2, float64(j)))
		}
	}

	res.Matr[0][0] *= 1

	return seconds, sizes
}

func MeasureTime(begin, end, step int) ([][]float32, []int) {
	seconds := make([][]float32, nAlgs)

	number := (end-begin)/step + 1

	sizes := make([]int, number)
	var matr1, matr2, res matrix.Matrix

	func_pointers := [nAlgs]func(*matrix.Matrix, *matrix.Matrix) (float32, matrix.Matrix){
		baseMulMatrixNoOptTimeMeasurement,
		baseMulMatrixOpt1TimeMeasurement,
		grapeMulMatrixNoOptTimeMeasurement,
		grapeMulMatrixOpt1and2TimeMeasurement,
		strassenMulMatrixNoOptTimeMeasurement,
		strassenMulMatrixOpt1TimeMeasurement,
	}

	for i := 0; i < nAlgs; i++ {
		seconds[i] = make([]float32, number)
	}

	for i := 0; i < nAlgs; i++ {
		for j := 0; j < number; j++ {
			matr1, matr2 = inp_out.GenerateSquareMatrixes(begin + step*j)
			seconds[i][j], res = func_pointers[i](&matr1, &matr2)
			fmt.Println(i)
			sizes[j] = begin + step*j
		}
	}

	res.Matr[0][0] *= 1

	return seconds, sizes
}
