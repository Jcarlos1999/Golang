package structs

import "fmt"

type Person struct {
	name string
	age  uint8
}

var (
	p1 = Person{"Joe", 34}
	p2 = Person{name: "Joe"}
	p3 = Person{}
	k  = &Person{"Foo", 53}
)

func main() {

	fmt.Println(Person{"joao", 23})

	person := Person{name: "Foo guy", age: 100}

	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Age: %d\n", person.age)

	v := &person
	v.age = 18
	fmt.Println("New age:", person)

	fmt.Println(p1, &k, p2, p3)
}
