package main

import (
	"fmt"
	"hash/crc32"
	"net/http"

	"bytes"
	_ "crypto/md5"
	_ "encoding/binary"
	_ "encoding/json"
	"io/ioutil"
	_ "sort"
	"strconv"
)

type Pair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var pairs []Pair
var DATASTORE3000 = "3000"
var DATASTORE3001 = "3001"
var DATASTORE3002 = "3002"
var NUMBEROFSERVER = 3

type ConsistentHashing struct {
	circle []Node
}

type Node struct {
	hashValue uint32
	nodeName  string
}

func (c *ConsistentHashing) add(pair Pair) {

	for i := 0; i < NUMBEROFSERVER; i++ {

		if c.circle[i].hashValue > hashing(strconv.Itoa(pair.Key)) {
			//fmt.Println("Server Found", c.circle[i])
			put(c.circle[i].nodeName, pair)
			return
		}
	}
	put(c.circle[0].nodeName, pair)
}

func (c *ConsistentHashing) get(pair Pair) {

	for i := 0; i < NUMBEROFSERVER; i++ {

		if c.circle[i].hashValue > hashing(strconv.Itoa(pair.Key)) {
			//	fmt.Println("Server Found", c.circle[i])
			get(c.circle[i].nodeName, pair)
			return
		}
	}
	get(c.circle[0].nodeName, pair)

}

func (c *ConsistentHashing) addServer(SERVERNAME string) {

	var node Node
	node.hashValue = hashing(SERVERNAME)
	node.nodeName = SERVERNAME
	c.circle = append(c.circle, node)

}

func (c *ConsistentHashing) display() {
	fmt.Println(c.circle)
}

func prepareInput() {
	//var temp string
	for i := 0; i < 10; i++ {
		pairs = append(pairs, Pair{})
		pairs[i].Key = i + 1
		//pairs[i].Value = temp
	}
	pairs[0].Value = "a"
	pairs[1].Value = "b"
	pairs[2].Value = "c"
	pairs[3].Value = "d"
	pairs[4].Value = "e"
	pairs[5].Value = "f"
	pairs[6].Value = "g"
	pairs[7].Value = "h"
	pairs[8].Value = "i"
	pairs[9].Value = "j"

	fmt.Println("The pairs are---", pairs)
}

func hashing(toBeHashed string) uint32 {
	hash := crc32.ChecksumIEEE
	hashing := hash([]byte(toBeHashed))
	return hashing
}

func put(SERVERNAME string, pair Pair) {
	client := &http.Client{}

	//hashValue := hashing(strconv.Itoa(pair.Key))
	urlStr := "http://localhost:" + SERVERNAME + "/keys/" + strconv.Itoa(pair.Key) + "/" + pair.Value
	fmt.Println("PUT", urlStr)
	jsonStr := []byte(`{}`)
	r1, _ := http.NewRequest("PUT", urlStr, bytes.NewBuffer(jsonStr))

	r1.Header.Set("Content-Type", "application/json")
	r1.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp1, err := client.Do(r1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response:", resp1.StatusCode)
	fmt.Println("Key Value Pair ", pairs[0].Key, "=>", pairs[0].Value, "is going to http://localhost:", SERVERNAME)
	defer resp1.Body.Close()

}

func get(SERVERNAME string, pair Pair) {

	urlStr := "http://localhost:" + SERVERNAME + "/keys/" + strconv.Itoa(pair.Key)
	fmt.Println("GET", urlStr)
	res, err := http.Get(urlStr)
	//fmt.Println(Url.String())
	if err != nil {
		panic("Error Panic")
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s\n", string(contents))

}

func main() {
	prepareInput()
	var consistentHashing ConsistentHashing
	consistentHashing.addServer(DATASTORE3000)
	consistentHashing.addServer(DATASTORE3001)
	consistentHashing.addServer(DATASTORE3002)

	fmt.Println("-----------------CLIENT LOGS ARE ENABLED-------------------------- func main()-------")
	fmt.Println("OUTPUT OF ALL PUT OPERATIONS \n\n")
	consistentHashing.add(pairs[0]) // PUT 1 => a
	consistentHashing.add(pairs[1]) // PUT 2 => b
	consistentHashing.add(pairs[2]) // PUT 3 => c
	consistentHashing.add(pairs[3]) // PUT 4 => d
	consistentHashing.add(pairs[4]) // PUT 5 => e
	consistentHashing.add(pairs[5]) // PUT 6 => f
	consistentHashing.add(pairs[6]) // PUT 7 => g
	consistentHashing.add(pairs[7]) // PUT 8 => h
	consistentHashing.add(pairs[8]) // PUT 9 => i
	consistentHashing.add(pairs[9]) // PUT 10 => j

	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("OUTPUT OF ALL GET OPERATIONS \n\n")
	consistentHashing.get(pairs[0]) // GET /Keys/1
	consistentHashing.get(pairs[1]) // GET /Keys/2
	consistentHashing.get(pairs[2]) // GET /Keys/3
	consistentHashing.get(pairs[3]) // GET /Keys/4
	consistentHashing.get(pairs[4]) // GET /Keys/5
	consistentHashing.get(pairs[5]) // GET /Keys/6
	consistentHashing.get(pairs[6]) // GET /Keys/7
	consistentHashing.get(pairs[7]) // GET /Keys/8
	consistentHashing.get(pairs[8]) // GET /Keys/9
	consistentHashing.get(pairs[9]) // GET /Keys/10

	consistentHashing.display()

	fmt.Println("-----------------CLIENT LOGS ARE ENABLED--------------------------func main()-------")

}

//1831254863
//438406105 http://localhost:3001} {2200492643 http://localhost:3002}]

//[{1831254863 http://localhost:3000} {438406105 http://localhost:3001} {2200492643 http://localhost:3002}]

//506358684
//1764317962
//4028631728
