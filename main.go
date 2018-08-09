package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Action int

const (
	Nothing Action = iota * 10
	Move
	Load
	Unload
	Eat
)

type Direction int

const (
	Up Direction = iota
	UpRight
	Right
	DownRight
	Down
	DownLeftf
	Left
	UpLeft
)

type Hive struct {
	Id   string
	Skin uint8
	Ox   uint8
	Oy   uint8
	Ants map[int]*Ant
	Map  *Map
}

type Map struct {
	Width     uint8
	Height    uint8
	HiveLimit uint8
	Cells     [][]*Cell
}

type Cell struct {
	Food uint8
	CellType
}

type CellType uint8

const (
	Dark CellType = iota
	Grass
	HillBump
	HillConer
	HillHalf
	HillThird
)

type Ant struct {
	Wasted  int
	Age     uint8
	Health  uint8
	Payload uint8
	X       uint8
	Y       uint8
	Event
}

type Event uint8

const (
	Good Event = iota
	Birth
	NoAction
	Slow
	BadMove
	BadLoad
	BadUnload
	BadEat
	Collision
	Error
	Death
)

func main() {
	fmt.Println("let's ant")
	http.HandleFunc("/", randomMoves)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func randomMoves(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	if len(data) == 0 {
		http.Error(w, err.Error(), 502)
		return
	}

	var hive Hive
	err = json.Unmarshal(data, &hive)
	if err != nil {
		fmt.Println("Fail to convrt json to object", err)
		http.Error(w, err.Error(), 503)
		return
	}

	rand.Seed(time.Now().UnixNano())
	ants := make(map[int]int)
	for id := range hive.Ants {
		ants[id] = rand.Intn(5)*10 + rand.Intn(4)
	}

	fmt.Println(ants)

	output, err := json.Marshal(ants)
	if err != nil {
		http.Error(w, err.Error(), 504)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
