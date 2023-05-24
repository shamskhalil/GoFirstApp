package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

var db = []Person{
	{Id: 0, FName: "Khalil", LName: "Shams"},
	{Id: 1, FName: "Aminu", LName: "Salis"},
	{Id: 2, FName: "Kabiru", LName: "Garba"},
}

func main() {
	url := "http://localhost:3000/person"
	go makeServer()

	time.Sleep(5 * time.Second)
	log.Println("Making get request ...")
	res, err := http.Get(url)
	checkErr(err)
	js, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	personObjects := []Person{}
	json.Unmarshal(js, &personObjects)
	fmt.Printf("%+v\n", personObjects)
	res.Body.Close()
	time.Sleep(1 * time.Hour)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

func makeServer() {
	http.HandleFunc("/person", getPersonHandler)
	http.HandleFunc("/person/", getOnePersonHandler)
	log.Print("Server listening on port 3000 !")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Error starting server! ", err)
	}
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("GET REQ <<<")
		fmt.Println(r.URL.Path)

		response, err := json.Marshal(db)
		checkErr(err)
		w.Header().Add("Content-Type", "application/json")
		w.Write(response)
	} else {
		//POST
		log.Println("POST REQ <<<")
		body, err := ioutil.ReadAll(r.Body)
		checkErr(err)
		defer r.Body.Close()
		person := Person{}
		newPerson := person.fromJson(body)
		db = append(db, *newPerson)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"msg\":\"Person created successfully\"}"))
	}

}

func getOnePersonHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET REQ <<<")
	fmt.Println(r.URL.Path)
	chunks := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(chunks[len(chunks)-1])
	checkErr(err)

	if r.Method == "GET" {
		response, err := json.Marshal(db[id])
		checkErr(err)
		w.Header().Add("Content-Type", "application/json")
		w.Write(response)
	} else if r.Method == "DELETE" {
		//delete id
		db = append(db[:id], db[id+1:]...)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"msg\":\"Person deleted successfully\"}"))
	}

}
