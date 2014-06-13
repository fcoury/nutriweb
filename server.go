package main

import (
	"encoding/json"
	"fatsecret"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port to listen")
	flag.Parse()

	http.HandleFunc("/foods", foodsHandler)
	http.HandleFunc("/food", foodHandler)

	// static
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving: www/%s\n", r.URL.Path[1:])
		http.ServeFile(w, r, "www/"+r.URL.Path[1:])
	})

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf(":%d", *port)
	err := http.ListenAndServe(addr, nil)
	log.Printf("%s\n", err.Error())
}

func foodsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Search by: " + r.FormValue("q"))

	foods, error, err := fatsecret.SearchFood(r.FormValue("q"), r.FormValue("page_size"), r.FormValue("page"))
	if err != nil {
		fmt.Fprintf(w, "Error 1")
		return
	}

	if error != nil {
		log.Println("Error: " + error.Message)
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

	log.Println("Found: " + strconv.Itoa(foods.TotalResults) + " results")
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
	log.Println("Food id: " + id)
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
