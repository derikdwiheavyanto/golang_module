package latihan

import "fmt"

type person struct {
	Name, Address string
	Age           int
}

func (person Person) getName() {
	fmt.Println("Hello my name is", person.Name)

}
