package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Request struct {
	Id   string `json:"id"`
	Tick int    `json:"tick"`
	Ants []Ant  `json:"ants"`
}

type Response struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	AntId     int    `json:"antId"`
	Action    string `json:"act"`
	Direction string `json:"dir"`
}

type Ant struct {
	Id      int    `json:"id"`
	Event   string `json:"event"`
	Errors  uint   `json:"errors"`
	Age     int    `json:"age"`
	Health  int    `json:"health"`
	Payload int    `json:"payload"`
	Point   Point  `json:"point"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	// starting listen for http calls on port :7070
	http.HandleFunc("/", handleAllRequests)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleAllRequests(w http.ResponseWriter, r *http.Request) {
	// your bot response should be json object
	w.Header().Set("content-type", "application/json")

	// unmarshal request payload into objects
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var req Request
	_ = json.Unmarshal(data, &req)

	// available actions and directions
	// var actions = []string{"stay", "move", "eat", "take", "put"}
	var directions = []string{"up", "down", "right", "left"}

	orders := []Order{}
	// loop through ants and give random move order
	for _, ant := range req.Ants {
		order := Order{
			AntId:  ant.Id,
			Action: "move",
			// pick random direction from slice on line 62
			Direction: directions[rand.Intn(3)],
		}
		orders = append(orders, order)
	}

	// prepare response object
	response := Response{Orders: orders}

	bytes, _ := json.Marshal(response)
	_, err := w.Write(bytes)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// fmt.Println(string(bytes)) // your json response should like this
	// {"orders": [
	//	 {"antId":1,"act":"move","dir":"down"},
	//	 {"antId":17,"act":"load","dir":"up"}
	//	]}
}

// this code available at https://github.com/anthive/go
// to test it localy, submit post request with payload.json using postman or curl
// curl -X 'POST' -d @payload.json http://localhost:7070

// have fun!
