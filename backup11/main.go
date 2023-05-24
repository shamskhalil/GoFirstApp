package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go makeServer()
	time.Sleep(2 * time.Second)
	doRequestTimeout()
	time.Sleep(1 * time.Hour)
}

//Background, TODO, ti

func doRequestTimeout() {
	fmt.Println("Making resquest ....")
	//tm := time.Duration(500 * int(time.Millisecond))
	//ctx, cancel := context.WithTimeout(context.Background(), tm)
	ctx, cancel := context.WithCancel(context.Background())
	go cancelRequest(cancel)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:3000/person", nil)
	checkErr(err)
	res, err := http.DefaultClient.Do(req)
	checkErr(err)
	jsonResponse, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	fmt.Printf("%s\n", jsonResponse)
	res.Body.Close()
}

func cancelRequest(c context.CancelFunc) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Cancelling request ...")
	c()
}

func doRequest() {
	fmt.Println("Making resquest ....")
	res, err := http.Get("http://localhost:3000/person")
	checkErr(err)
	jsonResponse, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	fmt.Printf("%s\n", jsonResponse)
	res.Body.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

func makeServer() {
	http.HandleFunc("/person", getPersonHandler)

	log.Print("Server listening on port 3000 !")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Error starting server! ", err)
	}
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("{\"person\":\"Khalil M. Shams\"}"))
}
