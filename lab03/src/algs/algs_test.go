package algs

import (
	"lab03/inp_out"
	"testing"
)

func BenchmarkBlock100(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 100)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		BlockSort(copyArr, len(copyArr))
	}
}

func BenchmarkBlock200(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 200)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		BlockSort(copyArr, len(copyArr))
	}
}

func BenchmarkBlock500(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 500)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		BlockSort(copyArr, len(copyArr))
	}
}

func BenchmarkBlock700(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 700)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		BlockSort(copyArr, len(copyArr))
	}
}

func BenchmarkBlock1000(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 1000)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		BlockSort(copyArr, len(copyArr))
	}
}

func BenchmarkRadix100(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 100)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		RadixSort(copyArr)
	}
}

func BenchmarkRadix200(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 200)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		RadixSort(copyArr)
	}
}

func BenchmarkRadix500(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 500)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		RadixSort(copyArr)
	}
}

func BenchmarkRadix700(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 700)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		RadixSort(copyArr)
	}
}

func BenchmarkRadix1000(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 1000)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		RadixSort(copyArr)
	}
}

func BenchmarkMerge100(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 100)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		MergeSort(copyArr)
	}
}

func BenchmarkMerge200(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 200)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		MergeSort(copyArr)
	}
}

func BenchmarkMerge500(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 500)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		MergeSort(copyArr)
	}
}

func BenchmarkMerge700(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 700)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		MergeSort(copyArr)
	}
}

func BenchmarkMerge1000(b *testing.B) {
	arr := inp_out.GenerateArrayForTest(false, 1000)

	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		MergeSort(copyArr)
	}
}
