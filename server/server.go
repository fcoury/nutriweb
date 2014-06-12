package main

import (
  "fmt"
  "fatsecret"
  "encoding/json"
  "net/http"
  "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Search by: " + r.FormValue("q"))

  foods, error, err := fatsecret.SearchFood(r.FormValue("q"))
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

func main() {
  http.HandleFunc("/foods", handler)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("../www/" + r.URL.Path[1:])
    http.ServeFile(w, r, "../www/" + r.URL.Path[1:])
  })
  http.ListenAndServe(":8080", nil)
}

