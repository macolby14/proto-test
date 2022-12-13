package main

import (
	"fmt"

	"github.com/macolby14/proto_test/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
)

func main(){
	fmt.Println("Test")
	
	p := tutorialpb.Person{
		Id: 1234,
		Name: "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorialpb.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorialpb.Person_HOME},
		},
	}

	fmt.Println(p)
}