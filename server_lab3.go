package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	//"net/url"
)

var counter int

type Pair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

type ResponseStructure struct {
	Pairs []Pair `json:"response"`
}

var Response3000 ResponseStructure
var Response3001 ResponseStructure
var Response3002 ResponseStructure

func getAllHandler3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "GET ALL end point is called")
	uj, _ := json.MarshalIndent(Response3000, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "locahost:3000 gives JSON response  --> %s", uj)
}

func getAllHandler3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "GET ALL end point is called")
	uj, _ := json.MarshalIndent(Response3001, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "locahost:3001 gives JSON response  --> %s", uj)
}

func getAllHandler3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "GET ALL end point is called")
	uj, _ := json.MarshalIndent(Response3002, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "locahost:3002 gives JSON response  --> %s", uj)
}
func getHandler3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "GET end point is called")
	id, _ := strconv.Atoi(p.ByName("id"))
	inputkeyvaluepair := Pair{}
	i := 0
	for i = 0; i < len(Response3000.Pairs); i++ {
		if Response3000.Pairs[i].Key == id {

			inputkeyvaluepair = Response3000.Pairs[i]

			break
		}
	}
	var getResponse ResponseStructure
	if i != len(Response3000.Pairs) {

		getResponse.Pairs = append(getResponse.Pairs, inputkeyvaluepair)

	}
	//fmt.Println(getResponse)
	uj, _ := json.MarshalIndent(getResponse, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "locahost:3000 gives JSON response  --> %s", uj)
}
func getHandler3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "GET end point is called")
	id, _ := strconv.Atoi(p.ByName("id"))
	inputkeyvaluepair := Pair{}
	i := 0
	for i = 0; i < len(Response3001.Pairs); i++ {
		if Response3001.Pairs[i].Key == id {

			inputkeyvaluepair = Response3001.Pairs[i]

			break
		}
	}

	var getResponse ResponseStructure
	if i != len(Response3001.Pairs) {

		getResponse.Pairs = append(getResponse.Pairs, inputkeyvaluepair)

	}
	//fmt.Println(getResponse)
	uj, _ := json.MarshalIndent(getResponse, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "locahost:3001 gives JSON response  --> %s", uj)
}

func getHandler3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "GET end point is called")
	id, _ := strconv.Atoi(p.ByName("id"))
	inputkeyvaluepair := Pair{}
	i := 0
	for i = 0; i < len(Response3002.Pairs); i++ {
		if Response3002.Pairs[i].Key == id {

			inputkeyvaluepair = Response3002.Pairs[i]

			break
		}
	}
	var getResponse ResponseStructure
	if i != len(Response3002.Pairs) {

		getResponse.Pairs = append(getResponse.Pairs, inputkeyvaluepair)

	}
	//fmt.Println(getResponse)
	uj, _ := json.MarshalIndent(getResponse, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "locahost:3002 gives JSON response  --> %s", uj)
}

func putHandler3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "PUT end point is called")
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	temp := Pair{}
	temp.Key = Id
	temp.Value = value
	Response3000.Pairs = append(Response3000.Pairs, temp)
	//fmt.Println("Response3000", Response3000)
	uj, _ := json.MarshalIndent(Response3000, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response --> %s", uj)

}

func putHandler3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "PUT end point is called")
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	temp := Pair{}
	temp.Key = Id
	temp.Value = value
	Response3001.Pairs = append(Response3001.Pairs, temp)
	//fmt.Println("Response3001", Response3001)
	uj, _ := json.MarshalIndent(Response3001, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response --> %s", uj)

}
func putHandler3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(r.Host, "PUT end point is called")
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")

	temp := Pair{}
	temp.Key = Id
	temp.Value = value
	Response3002.Pairs = append(Response3002.Pairs, temp)
	//	fmt.Println("Response3002", Response3002)
	uj, _ := json.MarshalIndent(Response3002, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response --> %s", uj)

}

func main() {
	fmt.Println("-----------------SERVER LOGS ARE ENABLED-------------------------- func main()-------")

	counter = 1

	r3000 := httprouter.New()
	r3001 := httprouter.New()
	r3002 := httprouter.New()
	r3000.GET("/keys", getAllHandler3000)
	r3001.GET("/keys", getAllHandler3001)
	r3002.GET("/keys", getAllHandler3002)

	r3000.GET("/keys/:id", getHandler3000)
	r3001.GET("/keys/:id", getHandler3001)
	r3002.GET("/keys/:id", getHandler3002)

	r3000.PUT("/keys/:id/:value", putHandler3000)
	r3001.PUT("/keys/:id/:value", putHandler3001)
	r3002.PUT("/keys/:id/:value", putHandler3002)

	go func() {
		server3000 := http.Server{
			Addr:    "0.0.0.0:3000",
			Handler: r3000,
		}
		server3000.ListenAndServe()
	}()

	go func() {
		server3001 := http.Server{
			Addr:    "0.0.0.0:3001",
			Handler: r3001,
		}
		server3001.ListenAndServe()
	}()
	server3002 := http.Server{
		Addr:    "0.0.0.0:3002",
		Handler: r3002,
	}
	server3002.ListenAndServe()
	fmt.Println("-----------------SERVER LOGS ARE ENABLED-------------------------- func main()-------")
}
