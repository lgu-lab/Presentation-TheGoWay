package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//HistoryItem is the one that will be serialized
type HistoryItem struct {
	N int `json:"Number"`
	M int `json:"Multiples"`
}

//HistoryDb is where the History will be cached, potentially a database
type HistoryDb struct {
	History []HistoryItem `json:"History"`
}

var historyDb = HistoryDb{History: []HistoryItem{}}

func (db *HistoryDb) append(n, m int) {

	item := HistoryItem{N: n, M: m}
	db.History = append(db.History, item)
	histCount := len(db.History)
	if histCount > 5 {
		histCount = 5
	}
	db.History = db.History[len(db.History)-histCount : len(db.History)]
}

func httpHistory(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(historyDb)
}

// TableValues keeps the computed values from multiplication
type TableValues struct {
	Number int
	Size   int
	Values []int
}

func httpCompute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, _ := strconv.Atoi(vars["n"])
	size, _ := strconv.Atoi(vars["m"])

	t := TableValues{Number: number, Size: size}
	t.Values = make([]int, size)
	for multiple := 0; multiple < size; multiple++ {
		t.Values[multiple] = number * (multiple + 1)
	}

	json.NewEncoder(w).Encode(t)
	historyDb.append(number, size)
}
