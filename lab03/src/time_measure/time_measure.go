package time_measure

import (
	"fmt"
	"lab03/algs"
	"lab03/inp_out"
	"syscall"
)

const N int = 10000
const nAlgs int = 3
const nTypesData int = 4

func GetCPU() int64 {
	usage := new(syscall.Rusage)
	syscall.Getrusage(syscall.RUSAGE_SELF, usage)
	return usage.Utime.Nano() + usage.Stime.Nano()
}

func blockSortTimeMeasurement(arr []int) (float32, []int) {
	var sum float32

	var startTime, finishTime int64
	var result []int

	for i := 0; i < N; i++ {
		result = make([]int, len(arr))
		copy(result, arr)
		startTime = GetCPU()
		result = algs.BlockSort(result, len(result))
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	result[0] += 1

	return (sum / float32(N)) / 1e+9, result
}

func radixSortTimeMeasurement(arr []int) (float32, []int) {
	var sum float32

	var startTime, finishTime int64
	var result []int

	for i := 0; i < N; i++ {
		result = make([]int, len(arr))
		copy(result, arr)
		startTime = GetCPU()
		result = algs.RadixSort(result)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	result[0] += 1

	return (sum / float32(N)) / 1e+9, result
}

func mergeSortTimeMeasurement(arr []int) (float32, []int) {
	var sum float32

	var startTime, finishTime int64
	var result []int

	for i := 0; i < N; i++ {
		result = make([]int, len(arr))
		copy(result, arr)
		startTime = GetCPU()
		result = algs.MergeSort(result)
		finishTime = GetCPU()
		sum += float32(finishTime - startTime)
	}

	result[0] += 1

	return (sum / float32(N)) / 1e+9, result
}

func MeasureTime(begin int, end int, step int) ([][][]float32, []int) {
	seconds := make([][][]float32, nTypesData, nTypesData)
	number := (end-begin)/step + 1
	for i := range seconds {
		seconds[i] = make([][]float32, nAlgs)
		for j := range seconds[i] {
			seconds[i][j] = make([]float32, number, number)
		}
	}

	sizes := make([]int, number)
	var arr, res []int

	func_pointers := [nAlgs]func([]int) (float32, []int){
		blockSortTimeMeasurement,
		radixSortTimeMeasurement,
		mergeSortTimeMeasurement,
	}

	for k := 0; k < nTypesData; k++ {
		for i := 0; i < nAlgs; i++ {
			for j := 0; j < number; j++ {
				arr = inp_out.GenerateArrayForTest(false, begin+j*step)
				seconds[k][i][j], res = func_pointers[i](arr)
				sizes[j] = begin + step*j
				fmt.Println(i, sizes[j])
			}
		}
	}

	res[0] *= 1

	return seconds, sizes
}
