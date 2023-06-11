package latihan

import "fmt"

type Address struct {
	Name string
	No   int
}

type Person struct {
	Name    string
	Address *Address
}

func CPointer() {
	address := Address{Name: "Besuki", No: 58}

	person := Person{Name: "Derik", Address: &address}

	address.changeAddress()
	person.changeName()

	fmt.Println("Nama: ", person.Name, ",Alamat: ", address.Name)

	fmt.Println("Nama: ", person.Name, ",Alamat: ", person.Address.Name)

}

func (address *Address) changeAddress() {
	*address = Address{Name: "bandung", No: 58}

}

func (person *Person) changeName() {
	*person.Address = Address{Name: "Pakel", No: 59}
	address := person.Address
	*address = Address{Name: "Campur", No: 57}
	person.Name = "Khulun"
}
