package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type TestStruct struct {
	Name string
	ID   int32
}

func main() {
	fmt.Println("start")
	data, err := bson.Marshal(&TestStruct{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Println("%q", data)

	value := TestStruct{}
	err2 := bson.Unmarshal(data, &value)
	if err2 != nil {
		panic(err)
	}
	fmt.Println("value:", value)

	mmap := bson.M{}
	err3 := bson.Unmarshal(data, mmap)
	if err3 != nil {
		panic(err)
	}
	fmt.Println("mmap:", mmap)

}