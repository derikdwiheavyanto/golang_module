package latihan

import (
	"fmt"
)

func CSlice() {
	// array := [3]string{"a", "b", "c"}
	// var name string = "eko kurniawan"
	// fmt.Println("Satu = ", name)
	// fmt.Println("Dua = ", 2)
	// fmt.Println("Tiga Koma Lima = ", 3.5)
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	slice := days[:5]
	slice2 := append(slice, "ndak roh")
	// fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(days)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice3 := make([]string, 2, 5)
	slice3[0] = "oke"
	slice3[1] = "ora"

	fmt.Println(slice3)

	slice4 := make([]string, len(slice3), cap(slice3))
	copy(slice4, slice3)
	fmt.Println("ini slice 4:	", slice4)

}

func Faktorial(angka int32) int32 {
	if angka == 0 {
		return 1
	}
	return angka * Faktorial(angka-1)
}

func FaktorialPangkat(angka int32, pangkat int32) int32 {
	if pangkat == 0 {
		return 1
	}
	return angka * FaktorialPangkat(angka, pangkat-1)
}

func LatihanRekursive1(n int32) int32 {
	if n == 1 {
		return 1
	}

	return n + LatihanRekursive1(n-1)
}

func LatihanRekursive2(input string, f byte, l byte) string {

	chars := []rune(input)
	if f >= l {
		return string(chars)
	}
	swap(chars, f, l)
	return LatihanRekursive2(string(chars), f+1, l-1)

}

func swap(v []rune, f byte, l byte) {
	temp := v[f]
	v[f] = v[l]
	v[l] = temp
}

func Fibonaci(n int) int {
	if n <= 1 {
		return n
	}
	fmt.Println(Fibonaci(n-1), "+", Fibonaci(n-2))
	return 1

}

func QuickSort(arr []int, low int) {
	lengthArr := (len(arr) - 1)

	if low <= lengthArr {
		indexPivot := partition(arr, low, lengthArr)
		fmt.Println("index", indexPivot)
		QuickSort(arr, indexPivot)
	}
}

func partition(arr []int, low int, lengthArr int) int {
	pivot := arr[low]
	swapMarker := low - 1
	// fmt.Println(pivot)

	for i := low; i <= lengthArr; i++ {
		if arr[i] < pivot {
			swapMarker++
			fmt.Println("\npivot: ", pivot)
			fmt.Println("current index: ", i)
			fmt.Println("swap marker: ", swapMarker)
			fmt.Println("value	: ", arr[i])
			fmt.Println("value swap	: ", arr[swapMarker])
			fmt.Println()

			swapSort(arr, swapMarker, i)
			PrintArray(arr)
		}
	}
	return low + 1
}

func swapSort(arr []int, low int, high int) {
	arr[low], arr[high] = arr[high], arr[low]
}

func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
