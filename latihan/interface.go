package latihan

import "fmt"

type Human interface {
	GetName()
}

// type Person struct {
// 	Name, Address string
// 	Age           int
// }

func (person Person) GetName() {
	fmt.Println("Hello my name is", person.Name)

}

func SayHello(human Human) {
	human.GetName()
	// fmt.Println("Hello my name is", person.Name)
}

func Apapun(pertama any, kedua any) {
	fmt.Println(pertama, kedua)
}
