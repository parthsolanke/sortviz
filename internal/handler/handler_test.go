package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/parthsolanke/SortViz/internal/sorter"
)

func TestSortHandler(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		algorithm      string
		expectedResult []int
	}{
		{"BubbleSort", []int{5, 3, 8, 4, 2}, "bubble", []int{2, 3, 4, 5, 8}},
		{"InsertionSort", []int{5, 3, 8, 4, 2}, "insertion", []int{2, 3, 4, 5, 8}},
		{"SelectionSort", []int{5, 3, 8, 4, 2}, "selection", []int{2, 3, 4, 5, 8}},
		{"MergeSort", []int{5, 3, 8, 4, 2}, "merge", []int{2, 3, 4, 5, 8}},
		{"HeapSort", []int{5, 3, 8, 4, 2}, "heap", []int{2, 3, 4, 5, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := struct {
				Array     []int  `json:"array"`
				Algorithm string `json:"algorithm"`
			}{
				Array:     tt.input,
				Algorithm: tt.algorithm,
			}

			body, err := json.Marshal(payload)
			if err != nil {
				t.Fatalf("could not marshal payload: %v", err)
			}

			req, err := http.NewRequest("POST", "/sort", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(SortHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			var response []sorter.SortStep
			if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
				t.Errorf("could not decode response: %v", err)
			}

			finalArray := response[len(response)-1].Array
			if !equal(finalArray, tt.expectedResult) {
				t.Errorf("unexpected result: got %v want %v", finalArray, tt.expectedResult)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
