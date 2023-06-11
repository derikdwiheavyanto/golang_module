package latihan

import "fmt"

type Filter func(string) string

func PrintName(name string, filter Filter) {
	result := filter(name)

	fmt.Println("Nama : ", result)
}

func LFilter(name string) string {
	if name == "anjing" || name == "kontol" {
		return "***"
	} else {
		return name
	}
}

func Fibonaci2(n int32) int32 {

	if n <= 1 {
		return n
	}
	return Fibonaci2(n-1) + Fibonaci2(n-2)
}
