package fatsecret

import (
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

func SearchFood(query string) (*Foods, *Error, error) {
  result, err := SearchFoodQuery(query)
  if err != nil {
    return nil, nil, err
  }

  error, err := CheckError(result)
  if err != nil {
  	return nil, nil, err
  }

  if error != nil {
  	return nil, error, nil
  }

  foods, err := ParseFoods(result)
  if err != nil {
    return nil, nil, err
  }
  return &foods, nil, nil
}

func GetFood(id string) (*FoodDetails, *Error, error) {
	result, err := GetFoodQuery(id)
	if err != nil {
		return nil, nil, err
	}

  error, err := CheckError(result)
  if err != nil {
  	return nil, nil, err
  }

  if error != nil {
  	return nil, error, nil
  }

	food_detail, err := ParseFoodDetails(result)
	if err != nil {
		return nil, nil, err
	}
	return &food_detail, nil, nil
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
