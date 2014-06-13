package main

import (
	"encoding/json"
	"fatsecret"
	"fmt"
	"net/http"
	"strconv"
)

func foodsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Search by: " + r.FormValue("q"))

	foods, error, err := fatsecret.SearchFood(r.FormValue("q"), r.FormValue("page_size"), r.FormValue("page"))
	if err != nil {
		fmt.Fprintf(w, "Error 1")
		return
	}

	if error != nil {
		fmt.Println("Error: " + error.Message)
		js, err := json.Marshal(error)
		if err != nil {
			fmt.Fprintf(w, "Error 3")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(js)
		return
	}

	fmt.Println("Found: " + strconv.Itoa(foods.TotalResults))
	js, err := json.Marshal(foods)
	if err != nil {
		fmt.Fprintf(w, "Error 2")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func foodHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println("Food id: " + id)
	food_details, _, err := fatsecret.GetFood(id)

	if err != nil {
		fmt.Fprintf(w, "Error 4")
		return
	}

	js, err := json.Marshal(food_details)
	if err != nil {
		fmt.Fprintf(w, "Error 5")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/foods", foodsHandler)
	http.HandleFunc("/food", foodHandler)

	// static
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("../www/" + r.URL.Path[1:])
		http.ServeFile(w, r, "../www/"+r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}
