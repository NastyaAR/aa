package algs

import (
	"math"
)

const threshold int = 32
const base int = 10

func SelectionSort(arr []int, dir int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minInd := i
		for j := i + 1; j < len(arr); j++ {
			if dir*arr[j] < dir*arr[minInd] {
				minInd = j
			}
		}
		arr[i], arr[minInd] = arr[minInd], arr[i]
	}

	return arr
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		el := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > el; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = el
	}

	return arr
}

func findMinAndMax(arr []int) (min int, max int) {
	min = arr[0]
	max = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func BlockSort(arr []int, nBuckets int) []int {
	buckets := make([][]int, nBuckets)
	min, max := findMinAndMax(arr)

	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0, len(arr))
	}

	for i := range arr {
		ind := nBuckets * (arr[i] - min) / (max - min + 1)
		buckets[ind] = append(buckets[ind], arr[i])
	}

	for i := range buckets {
		if len(buckets[i]) < threshold && len(buckets[i]) > 0 {
			buckets[i] = InsertionSort(buckets[i])
		} else if len(buckets[i]) >= threshold {
			buckets[i] = BlockSort(buckets[i], len(buckets[i]))
		}
	}

	pos := 0
	for i := range buckets {
		if len(buckets[i]) != 0 {
			for j := 0; j < len(buckets[i]); j++ {
				arr[pos] = buckets[i][j]
				pos++
			}
		}
	}

	return arr
}

func getPlace(num int) int {
	places := 1
	num /= 10

	if num < 0 {
		num *= -1
	}

	for ; num > 0; num /= 10 {
		places++
	}
	return places
}

func getMaxPlace(arr []int) int {
	maxPlace := getPlace(arr[0])

	for i := 1; i < len(arr); i++ {
		curPlace := getPlace(arr[i])
		if curPlace > maxPlace {
			maxPlace = curPlace
		}
	}

	return maxPlace
}

func RadixSort(arr []int) []int {
	var ind int
	maxPlace := getMaxPlace(arr)

	buckets := make([][]int, 2*base)
	for i := 0; i < 2*base; i++ {
		buckets[i] = make([]int, 0, len(arr))
	}

	for place := 0; place < maxPlace; place++ {
		for j := range arr {
			if arr[j] < 0 {
				ind = base - (-arr[j]/int(math.Pow10(place)))%base
			} else {
				ind = base + (arr[j]/int(math.Pow10(place)))%base
			}
			buckets[ind] = append(buckets[ind], arr[j])
		}
		pos := 0
		for m := range buckets {
			if len(buckets[m]) != 0 {
				for n := range buckets[m] {
					arr[pos] = buckets[m][n]
					pos++
				}
			}
			buckets[m] = make([]int, 0, len(arr))
		}
	}

	return arr
}

func merge(left, right []int) []int {
	arr := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			arr = append(arr, left[0])
			left = left[1:]
		} else {
			arr = append(arr, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		arr = append(arr, left...)
	}
	if len(right) > 0 {
		arr = append(arr, right...)
	}

	return arr
}

func MergeSort(arr []int) []int {
	if len(arr) >= 2 {
		middle := len(arr) / 2
		arr = merge(MergeSort(arr[:middle]), MergeSort(arr[middle:]))
	}

	return arr
}
