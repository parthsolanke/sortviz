package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/parthsolanke/SortViz/internal/sorter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	tmpl.Execute(w, nil)
}

func SortHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Array     []int  `json:"array"`
		Algorithm string `json:"algorithm"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var steps []sorter.SortStep

	switch data.Algorithm {
	case "bubble":
		steps = sorter.BubbleSort(data.Array)
	case "insertion":
		steps = sorter.InsertionSort(data.Array)
	case "selection":
		steps = sorter.SelectionSort(data.Array)
	case "merge":
		steps = sorter.MergeSort(data.Array)
	case "quick":
		steps = sorter.QuickSort(data.Array)
	case "heap":
		steps = sorter.HeapSort(data.Array)
	default:
		http.Error(w, "Unknown sorting algorithm", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}
