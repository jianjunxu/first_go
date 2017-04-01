package main

import "fmt"

func main() {

}

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
	loan   float32
}
type Employee struct {
	Human
	company string
	money   float32
}

func (h *Human) sayHi() {
	fmt.Printf("Hi,i am %s you can call me on %s\n", h.name, h.phone)
}
func (h *Human) Sing(lyrics string) {
	fmt.Println("la la la la la ...", lyrics)
}
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle ...", beerStein)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}
func (e *Employee) SpendSalary(amount float32) {
	e.money += amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}
type ElderlyGent interface {
	SayHi()
	Sin(song string)
	SpendSalary(amount float32)
}