package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/xml"
)

type Food struct {
  Id int `xml:"food_id"`
  Name string `xml:"food_name"`
  Brand string `xml:"brand_name"`
  Type string `xml:"food_type"`
  Url string `xml:"food_url"`
  Description string `xml:"food_description"`
}

type Foods struct {
  MaxResults int `xml:"max_results"`
  TotalResults int `xml:"total_results"`
  PageNumber int `xml:"page_number"`
  FoodList []Food `xml:"food"`
}

func (s Food) String() string {
  return fmt.Sprintf("%d - %s - %s - %s\n%s\n%s\n", s.Id, s.Name, s.Brand, s.Type, s.Url, s.Description)
}

func (s Foods) String() string {
  return fmt.Sprintf("%d, %d", s.MaxResults, s.TotalResults)
}

func parseFoods(b []byte) (Foods, error) {
  var q Foods
  xml.Unmarshal(b, &q)

  for _, food := range q.FoodList {
    fmt.Printf("%s\n", food)
  }

  return q, nil
}

func main() {
  xmlFile, err := os.Open("sample.xml")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer xmlFile.Close()

  b, _ := ioutil.ReadAll(xmlFile)

  fmt.Println(string(b))

  foods, _ := parseFoods(b)
  for _, food := range foods.FoodList {
    fmt.Printf("%s\n", food)
  }

}
