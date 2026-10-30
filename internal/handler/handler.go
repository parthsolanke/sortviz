package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/parthsolanke/SortViz/internal/sorter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "web/templates/index.html")
}

func AlgorithmsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "web/templates/algorithms.html")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "web/templates/about.html")
}

func renderTemplate(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Could not load page", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Could not render page", http.StatusInternalServerError)
	}
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
	case "shell":
		steps = sorter.ShellSort(data.Array)
	case "cocktail":
		steps = sorter.CocktailSort(data.Array)
	default:
		http.Error(w, "Unknown sorting algorithm", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(steps)
}
