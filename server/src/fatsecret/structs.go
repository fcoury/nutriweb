package fatsecret

type Food struct {
  Id          int    `xml:"food_id"`
  Name        string `xml:"food_name"`
  Brand       string `xml:"brand_name"`
  Type        string `xml:"food_type"`
  Url         string `xml:"food_url"`
  Description string `xml:"food_description"`
}

type Foods struct {
  MaxResults   int    `xml:"max_results"`
  TotalResults int    `xml:"total_results"`
  PageNumber   int    `xml:"page_number"`
  FoodList     []Food `xml:"food"`
}

type Serving struct {
  Id                     int     `xml:"serving_id"`
  Description            string  `xml:"serving_description"`
  Url                    string  `xml:"serving_url"`
  MetricServingAmount    float32 `xml:"metric_serving_amount"`
  MetricServingUnit      string  `xml:"metric_serving_unit"`
  NumberOfUnits          float32 `xml:"number_of_units"`
  MeasurementDescription string  `xml:"measurement_description"`
  Calories               int     `xml:"calories"`
  Carbohydrate           float32 `xml:"carbohydrate,omitempty"`
  Protein                float32 `xml:"protein,omitempty"`
  Fat                    float32 `xml:"fat,omitempty"`
  SaturatedFat           float32 `xml:"saturated_fat,omitempty"`
  PolyunsaturatedFat     float32 `xml:"polyunsaturated_fat,omitempty"`
  MonounsaturatedFat     float32 `xml:"monounsaturated_fat,omitempty"`
  TransFat               float32 `xml:"trans_fat,omitempty"`
  Cholesterol            float32 `xml:"cholesterol,omitempty"`
  Sodium                 float32 `xml:"sodium,omitempty"`
  Potassium              float32 `xml:"potassium,omitempty"`
  Fiber                  float32 `xml:"fiber,omitempty"`
  Sugar                  float32 `xml:"sugar,omitempty"`
}

type Servings struct {
  Servings []Serving `xml:"serving"`
}

type FoodDetails struct {
  Id           int      `xml:"food_id"`
  Name         string   `xml:"food_name"`
  BrandName    string   `xml:"brand_name"`
  Type         string   `xml:"food_type"`
  Url          string   `xml:"food_url"`
  ServingsList Servings `xml:"servings"`
}

type Error struct {
  Code    int    `xml:"code"`
  Message string `xml:"message"`
}
