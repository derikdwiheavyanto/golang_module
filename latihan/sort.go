package latihan

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (a UserSlice) Len() int           { return len(a) }
func (a UserSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserSlice) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Sort() {
	data := []User{
		{Name: "Derik", Age: 20},
		{Name: "Bagong", Age: 19},
		{Name: "Bima", Age: 21},
		{Name: "Danial", Age: 22},
	}

	fmt.Println(data)

	sort.Sort(UserSlice(data))

	fmt.Println(data)
}
