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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// your bot respons should be json object
		w.Header().Set("content-type", "application/json")

		data, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var req Request
		_ = json.Unmarshal(data, &req)

		// available actions and directions
		var actions = []string{"stay", "move", "eat", "take", "put"}
		var directions = []string{"up", "down", "right", "left"}
		response := Response{
			Orders: make([]Order, 0),
		}

		// loop through ants and give orders
		for _, ant := range req.Ants {
			order := Order{
				AntId:     ant.Id,
				Action:    actions[rand.Intn(4)],    // pick random action from array on line 54
				Direction: directions[rand.Intn(3)], // pick random direction from array on line 55
			}

			// add order to your response object from line 20
			response.Orders = append(response.Orders, order)
		}

		bytes, _ := json.Marshal(response)
		w.Write(bytes)

		// fmt.Println(string(bytes))
		// {"orders": [
		//	 {"antId":1,"act":"move","dir":"down"},
		//	 {"antId":17,"act":"load","dir":"up"}
		//	]}
	})

	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// this code available at https://github.com/anthive/go
// to test it localy, submit post request with payload.json using postman or curl
// curl -X 'POST' -d @payload.json http://localhost:7070

// have fun!
