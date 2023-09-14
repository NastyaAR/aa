package algs1

import (
	"testing"
)

func BenchmarkRecursiveLeven5(b *testing.B) {
	s1 := []rune("point")
	s2 := []rune("lines")

	for i := 0; i < b.N; i++ {
		ResursiveLeven(s1, s2)
	}
}

func BenchmarkRecursiveLen10(b *testing.B) {
	s1 := []rune("pointlucky")
	s2 := []rune("linesviews")

	for i := 0; i < b.N; i++ {
		ResursiveLeven(s1, s2)
	}
}

func BenchmarkRecursiveLen15(b *testing.B) {
	s1 := []rune("pointluckycrazy")
	s2 := []rune("linesviewslikes")

	for i := 0; i < b.N; i++ {
		ResursiveLeven(s1, s2)
	}
}

func BenchmarkDamerauLeven5(b *testing.B) {
	s1 := []rune("point")
	s2 := []rune("lines")

	for i := 0; i < b.N; i++ {
		DamerauLeven(s1, s2)
	}
}

func BenchmarkDamerauLen10(b *testing.B) {
	s1 := []rune("pointlucky")
	s2 := []rune("linesviews")

	for i := 0; i < b.N; i++ {
		DamerauLeven(s1, s2)
	}
}

func BenchmarkDamerauLevenLen15(b *testing.B) {
	s1 := []rune("pointluckycrazy")
	s2 := []rune("linesviewslikes")

	for i := 0; i < b.N; i++ {
		DamerauLeven(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenLen5(b *testing.B) {
	s1 := []rune("point")
	s2 := []rune("lines")

	for i := 0; i < b.N; i++ {
		MatrixDamerauLeven(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenLen10(b *testing.B) {
	s1 := []rune("pointlucky")
	s2 := []rune("linesviews")

	for i := 0; i < b.N; i++ {
		MatrixDamerauLeven(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenLen15(b *testing.B) {
	s1 := []rune("pointluckycrazy")
	s2 := []rune("linesviewslikes")

	for i := 0; i < b.N; i++ {
		MatrixDamerauLeven(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenLen50(b *testing.B) {
	s1 := []rune("wpyibmiiniednngwziyqgeykknlfourltggbkwquqctwphduhm")
	s2 := []rune("rplsaehpegyzrzdmvzrhyyjfyzryztocdvrcisnfdmnewevncw")

	for i := 0; i < b.N; i++ {
		MatrixDamerauLeven(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenLen100(b *testing.B) {
	s1 := []rune("sptjdsihnhmnazevbjigfregecmdkiqlysbmcqekymzbufrqcrkrawofpbjuvskkrbkqrhiokxqzqwjjseppbhreddsnikcwsuyq")
	s2 := []rune("uuxahkjlvwefmbhrtgbwozechtqkdyxcoqsldeuxxjkrkzwjcpuavsmliwfiznnkywmfcilkyugztnksajmippfilenxcutjbhqp")

	for i := 0; i < b.N; i++ {
		MatrixDamerauLeven(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenCacheLen5(b *testing.B) {
	s1 := []rune("point")
	s2 := []rune("lines")

	for i := 0; i < b.N; i++ {
		RecursiveDamerauLevenWithCache(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenCacheLen10(b *testing.B) {
	s1 := []rune("pointlucky")
	s2 := []rune("linesviews")

	for i := 0; i < b.N; i++ {
		RecursiveDamerauLevenWithCache(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenCacheLen15(b *testing.B) {
	s1 := []rune("pointluckycrazy")
	s2 := []rune("linesviewslikes")

	for i := 0; i < b.N; i++ {
		RecursiveDamerauLevenWithCache(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenCacheLen50(b *testing.B) {
	s1 := []rune("wpyibmiiniednngwziyqgeykknlfourltggbkwquqctwphduhm")
	s2 := []rune("rplsaehpegyzrzdmvzrhyyjfyzryztocdvrcisnfdmnewevncw")

	for i := 0; i < b.N; i++ {
		RecursiveDamerauLevenWithCache(s1, s2)
	}
}

func BenchmarkMatrixDamerauLevenCacheLen100(b *testing.B) {
	s1 := []rune("sptjdsihnhmnazevbjigfregecmdkiqlysbmcqekymzbufrqcrkrawofpbjuvskkrbkqrhiokxqzqwjjseppbhreddsnikcwsuyq")
	s2 := []rune("uuxahkjlvwefmbhrtgbwozechtqkdyxcoqsldeuxxjkrkzwjcpuavsmliwfiznnkywmfcilkyugztnksajmippfilenxcutjbhqp")

	for i := 0; i < b.N; i++ {
		RecursiveDamerauLevenWithCache(s1, s2)
	}
}
