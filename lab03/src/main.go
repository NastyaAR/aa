package main

import (
	"fmt"
	"lab03/algs"
	"lab03/inp_out"
	"lab03/time_measure"
)

func main() {
	var arr, sortedArr, copyArr []int
	var err error

	choice := 1

	inp_out.OutputMenu()

	for choice != 0 {
		choice, err = inp_out.ChooseOption()

		if err != nil || choice < 0 || choice > 11 {
			fmt.Print(err)
			continue
		}

		switch choice {
		case 1:
			arr, err = inp_out.GenerateArray(false)
			if err != nil {
				fmt.Print(err)
				continue
			}
		case 2:
			arr, err = inp_out.GenerateArray(true)
			if err != nil {
				fmt.Print(err)
				continue
			}
		case 3:
			arr, err = inp_out.GenerateArray(false)
			if err != nil {
				fmt.Print(err)
				continue
			}
			arr = algs.SelectionSort(arr, 1)
		case 4:
			arr, err = inp_out.GenerateArray(false)
			if err != nil {
				fmt.Print(err)
				continue
			}
			arr = algs.SelectionSort(arr, -1)
		case 5:
			arr, err = inp_out.InputArray()
			if err != nil {
				fmt.Print(err)
				continue
			}
		case 6:
			fmt.Println("\nМассив:")
			fmt.Println(arr)
		case 7:
			copyArr = make([]int, len(arr))
			copy(copyArr, arr)
			sortedArr = algs.BlockSort(copyArr, len(copyArr))
			fmt.Println("\nОтсортированный массив:")
			fmt.Println(sortedArr)
		case 8:
			copyArr = make([]int, len(arr))
			copy(copyArr, arr)
			sortedArr = algs.RadixSort(copyArr)
			fmt.Println("\nОтсортированный массив:")
			fmt.Println(sortedArr)
		case 9:
			copyArr = make([]int, len(arr))
			copy(copyArr, arr)
			sortedArr = algs.MergeSort(copyArr)
			fmt.Println("\nОтсортированный массив:")
			fmt.Println(sortedArr)
		case 10:
			seconds, sizes := time_measure.MeasureTime(10, 1000, 50)
			inp_out.GetFullCsvs(seconds, sizes)
			inp_out.OutputTimeIntoFiles(seconds, sizes)
		case 11:
			inp_out.OutputMenu()
		}
	}
}
