package time_measure

import (
	// "fmt"
	"math/rand"
	"syscall"

	"lab1.com/algs1"
	"lab1.com/matrix"
)

const N int = 100
const symbols string = "abcdefghijklmnopqrstuvwxyz"
const MaxInd int = 26

func GetCPU() int64 {
	usage := new(syscall.Rusage)
	syscall.Getrusage(syscall.RUSAGE_SELF, usage)
	return usage.Utime.Nano() + usage.Stime.Nano()
}

func GetRandomString(len int) []rune {
	rstr := make([]rune, len)
	for i := 0; i < len; i++ {
		symb := rand.Intn(MaxInd)
		rstr[i] = rune(symbols[symb])
	}
	return rstr
}

func RecLevenTimeMeasurement(str1 []rune, str2 []rune) (float32, int) {
	var sum float32

	var startTime, finishTime int64
	var ans int

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		ans = algs1.ResursiveLeven(str1, str2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	return (sum / float32(N)) / 1e+9, ans
}

func DamerauLevenTimeMeasurement(str1 []rune, str2 []rune) (float32, int) {
	var sum float32

	var startTime, finishTime int64
	var ans int

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		ans = algs1.DamerauLeven(str1, str2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	return (sum / float32(N)) / 1e+9, ans
}

func MatrixDamerauLevenTimeMeasurement(str1 []rune, str2 []rune) (float32, int) {
	var sum float32

	var startTime, finishTime int64
	var ans int
	var matr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		ans, matr = algs1.MatrixDamerauLeven(str1, str2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	matr.Matr[0][0] += 1

	return (sum / float32(N)) / 1e+9, ans
}

func RecursiveDamerauLevenWithCacheTimeMeasurement(str1 []rune, str2 []rune) (float32, int) {
	var sum float32

	var startTime, finishTime int64
	var ans int
	var matr matrix.Matrix

	for i := 0; i < N; i++ {
		startTime = GetCPU()
		ans, matr = algs1.RecursiveDamerauLevenWithCache(str1, str2)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	matr.Matr[0][0] += 1

	return (sum / float32(N)) / 1e+9, ans
}

func MeasureTime(rstr1 []rune, rstr2 []rune) [4]float32 {
	var seconds [4]float32
	var ans int

	seconds[0], ans = RecLevenTimeMeasurement(rstr1, rstr2)
	seconds[1], ans = DamerauLevenTimeMeasurement(rstr1, rstr2)
	seconds[2], ans = MatrixDamerauLevenTimeMeasurement(rstr1, rstr2)
	seconds[3], ans = RecursiveDamerauLevenWithCacheTimeMeasurement(rstr1, rstr2)

	ans += 1

	return seconds
}
