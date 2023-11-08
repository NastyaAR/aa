package algs

import (
	"lab02/inp_out"
	"lab02/matrix"
	"testing"
)

func BenchmarkGrapeNoOpt10(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(10)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeNoOpt50(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(50)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeNoOpt100(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(100)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeNoOpt200(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(200)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeNoOpt500(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(500)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeOpt1and2_10(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(10)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixOpt1and2(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeOpt1and2_50(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(50)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixOpt1and2(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeOpt1and2_100(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(100)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixOpt1and2(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeOpt1and2_200(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(200)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixOpt1and2(&mtr1, &mtr2)
	}
}

func BenchmarkGrapeOpt1and2_500(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(500)

	for i := 0; i < b.N; i++ {
		GrapeMulMatrixOpt1and2(&mtr1, &mtr2)
	}
}

func BenchmarkClassicNoOpt_10(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(10)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkClassicNoOpt_50(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(50)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkClassicNoOpt_100(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(100)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkClassicNoOpt_200(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(200)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkClassicNoOpt_500(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(500)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkClassicOpt_10(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(10)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixOpt1(&mtr1, &mtr2)
	}
}
func BenchmarkClassicOpt_50(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(50)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixOpt1(&mtr1, &mtr2)
	}
}
func BenchmarkClassicOpt_100(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(100)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixOpt1(&mtr1, &mtr2)
	}
}
func BenchmarkClassicOpt_200(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(200)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixOpt1(&mtr1, &mtr2)
	}
}
func BenchmarkClassicOpt_500(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(500)

	for i := 0; i < b.N; i++ {
		BaseMulMatrixOpt1(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenNoOpt_10(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(10)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenNoOpt_50(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(50)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenNoOpt_100(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(100)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenNoOpt_200(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(200)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenNoOpt_500(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(500)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixNoOpt(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenOpt_10(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(10)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixOpt1(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenOpt_50(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(50)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixOpt1(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenOpt_100(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(100)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixOpt1(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenOpt_200(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(200)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixOpt1(&mtr1, &mtr2)
	}
}

func BenchmarkStrassenOpt_500(b *testing.B) {
	mtr1, mtr2 := inp_out.GenerateSquareMatrixes(500)

	matrix.CompleteMatrixes(&mtr1, &mtr2)

	for i := 0; i < b.N; i++ {
		StrassenMulMatrixOpt1(&mtr1, &mtr2)
	}
}
