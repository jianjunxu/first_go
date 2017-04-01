package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	var P person
	P.name = "jayden"
	P.age = 20

	P1 := person{"Tom", 25}
	P2 := person{name:"David", age:30}

	fmt.Printf("P.name:%s, P.age:%d\r\n", P.name, P.age)
	fmt.Printf("P1.name:%s, P1.age:%d\r\n", P1.name, P1.age)
	fmt.Printf("P2.name:%s, P2.age:%d\r\n", P2.name, P2.age)

	// struct作为匿名字段
	type Human struct {
		name   string
		age    int
		weight int
	}
	type Student struct {
		Human
		specialty string
	}
	mark := Student{Human{"Tim", 20, 55}, "English"}
	fmt.Printf("Student mark: name:%s, age:%d, weight:%d, specialty:%s\r\n", mark.name, mark.age, mark.weight, mark.specialty)


}
