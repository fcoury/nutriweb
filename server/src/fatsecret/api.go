package fatsecret

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"utils"
)

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

func SearchFood(query string) (*Foods, *Error, error) {
	result, err := SearchFoodQuery(query)
	if err != nil {
		return nil, nil, err
	}

	if strings.Contains(string(result), "<error") {
		fmt.Println("ERROR DETECTED")
		error, err := ParseError(result)
		if err != nil {
			return nil, nil, err
		}
		return nil, &error, nil
	} else {
		foods, err := ParseFoods(result)
		if err != nil {
			return nil, nil, err
		}
		return &foods, nil, nil
	}
}

func GetFood(id string) (*FoodDetails, *Error, error) {
	result, err := GetFoodQuery(id)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(string(result))

	if strings.Contains(string(result), "<error") {
		return nil, nil, err
	} else {
		food_detail, err := ParseFoodDetails(result)
		if err != nil {
			return nil, nil, err
		}
		return &food_detail, nil, nil
	}
}

func GetFoodQuery(id string) ([]byte, error) {
	params := map[string]string{"method": "food.get", "food_id": id}

	return SendQuery(params)
}

func SearchFoodQuery(query string) ([]byte, error) {
	params := make(map[string]string)
	params["method"] = "foods.search"
	params["search_expression"] = strings.Replace(query, " ", "+", -1)

	body, err := SendQuery(params)
	return body, err
}

func SendQuery(params map[string]string) ([]byte, error) {
	fatSecretUrl := "http://platform.fatsecret.com/rest/server.api"
	fatSecretConsumerKey := "62cc7c5caaf542668006fc70cbfdabae"
	fatSecretAccessSecret := "de666f86e8634a77947c02fc39cf33cd"

	oauth_timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	oauth_nonce := strconv.FormatInt(rand.Int63(), 16) // + strconv.FormatInt(rand.Int63(), 16)

	params["oauth_consumer_key"] = fatSecretConsumerKey
	params["oauth_nonce"] = oauth_nonce
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_timestamp"] = oauth_timestamp
	params["oauth_version"] = "1.0"

	params = utils.SortMap(params)

	paramsStr := ""
	for k, v := range params {
		paramsStr += k + "=" + url.QueryEscape(v) + "&"
	}

	paramsStr = strings.TrimSuffix(paramsStr, "&")

	sigBaseStr := "GET&" + url.QueryEscape(fatSecretUrl) + "&" + url.QueryEscape(paramsStr)
	sharedSecret := fatSecretAccessSecret + "&"

	sig := url.QueryEscape(utils.Sign(sigBaseStr, sharedSecret))

	paramsStr += "&oauth_signature=" + sig

	fmt.Println(fatSecretUrl + "?" + paramsStr)

	resp, err := http.Get(fatSecretUrl + "?" + paramsStr)

	fmt.Println(fmt.Sprintf("Response code: %d", resp.StatusCode))

	// Defer the closing of the body
	defer resp.Body.Close()

	// Read the content into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s Food) String() string {
	return fmt.Sprintf("%d - %s - %s - %s\n%s\n%s\n", s.Id, s.Name, s.Brand, s.Type, s.Url, s.Description)
}

func (s Foods) String() string {
	return fmt.Sprintf("%d, %d", s.MaxResults, s.TotalResults)
}

func ParseFoodDetails(b []byte) (FoodDetails, error) {
	var q FoodDetails
	xml.Unmarshal(b, &q)

	return q, nil
}

func ParseFoods(b []byte) (Foods, error) {
	var q Foods
	xml.Unmarshal(b, &q)

	return q, nil
}

func ParseError(b []byte) (Error, error) {
	var q Error
	xml.Unmarshal(b, &q)

	return q, nil
}
