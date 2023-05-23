package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id    int    `json:"id"`
	FName string `json:"fName"`
	LName string `json:"lName"`
}

func (p *Person) toJson() []byte {
	js, _ := json.Marshal(p)
	return js
}

func (p *Person) fromJson(personStr []byte) *Person {
	newPerson := &Person{}
	json.Unmarshal(personStr, newPerson)
	return newPerson
}

func main() {
	p1 := Person{Id: 1, FName: "Khalil", LName: "Shams"}
	p1j := p1.toJson()
	fmt.Printf("%s\n", p1j)
	postReq := []byte("{\"id\":2,\"fName\":\"Aminu\",\"lName\":\"Garba\"}")
	p2 := Person{}
	fmt.Printf("%+v\n", p2.fromJson(postReq))
}
