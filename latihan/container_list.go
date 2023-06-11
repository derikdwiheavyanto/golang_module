package latihan

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func ContainerList() {
	data := list.New()

	data.PushBack("Derik")
	data.PushFront("Dwi")
	data.PushBack("Heavyanto")

	for i := data.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

}

func ContainerRing() {
	data := ring.New(3)

	for i := 0; i < data.Len(); i++ {
		data.Value = "data-" + strconv.FormatInt(int64(i), 10)
		data = data.Next()

	}

	data.Do(func(a any) {
		fmt.Println(a)
	})
}
