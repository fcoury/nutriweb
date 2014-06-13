package fatsecret

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
  "math/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	// "regexp"
	"strconv"
	"strings"
	"time"
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

type Error struct {
	Code    int    `xml:"code"`
	Message string `xml:"message"`
}

func (s Food) String() string {
	return fmt.Sprintf("%d - %s - %s - %s\n%s\n%s\n", s.Id, s.Name, s.Brand, s.Type, s.Url, s.Description)
}

func (s Foods) String() string {
	return fmt.Sprintf("%d, %d", s.MaxResults, s.TotalResults)
}

func SearchFood(query string) (*Foods, *Error, error) {
	result, err := Query(query)
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

func Sign(base string, secret string) (string) {
  hasher := hmac.New(sha1.New, []byte(secret))
  io.WriteString(hasher, base)

  return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func Get(getUrl string, params map[string]string) ([]byte, error) {
	fatSecretAccessSecret := "de666f86e8634a77947c02fc39cf33cd"
	// getUrl = getUrl + "?"

  paramsStr := ""
	for k, v := range params {
		paramsStr += k + "=" + url.QueryEscape(v) + "&"
	}

	paramsStr = strings.TrimSuffix(paramsStr, "&")

  fmt.Println(getUrl)
	fmt.Println(paramsStr)

  sigBaseStr := "GET&" + url.QueryEscape(getUrl) + "&" + url.QueryEscape(paramsStr)
  sharedSecret := fatSecretAccessSecret + "&"

  sig := url.QueryEscape(Sign(sigBaseStr, sharedSecret))
  fmt.Println("Sig: " + sig)

	paramsStr += "&oauth_signature=" + sig

	fmt.Println(getUrl + "?" + paramsStr)

	resp, err := http.Get(getUrl + "?" + paramsStr)

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

func Post(postUrl string, params map[string]string) ([]byte, error) {
  fatSecretAccessSecret := "de666f86e8634a77947c02fc39cf33cd"
	values := url.Values{}

	for k, v := range params {
		values.Add(k, v)
	}

  paramStr := ""

  for k, v := range values {
    paramStr = paramStr + "&" + k + "=" + url.QueryEscape(v[0])
  }

  paramStr = strings.TrimPrefix(paramStr, "&")

  sigBaseStr := "POST&" + url.QueryEscape(postUrl) + "&" + url.QueryEscape(paramStr)
  sharedSecret := fatSecretAccessSecret + "&"

  fmt.Println("sigBaseStr: " + sigBaseStr)

  oauth_signature := Sign(sigBaseStr, sharedSecret)

  values.Add("oauth_signature", oauth_signature)

	resp, err := http.PostForm(postUrl, values)
	if err != nil {
		return nil, err
	}

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

// This function fetch the content of a URL will return it as an
// array of bytes if retrieved successfully.
func Query(query string) ([]byte, error) {
	fatSecretUrl := "http://platform.fatsecret.com/rest/server.api"
	fatSecretConsumerKey := "62cc7c5caaf542668006fc70cbfdabae"
	fatSecretAccessSecret := "de666f86e8634a77947c02fc39cf33cd"

	oauth_timestamp := strconv.FormatInt(time.Now().Unix(), 10)
  oauth_nonce := strconv.FormatInt(rand.Int63(), 16) // + strconv.FormatInt(rand.Int63(), 16)

	apiValues := make(map[string]string)
	apiValues["method"] = "foods.search"
	apiValues["oauth_consumer_key"] = fatSecretConsumerKey
	apiValues["oauth_nonce"] = oauth_nonce
	apiValues["oauth_signature_method"] = "HMAC-SHA1"
	apiValues["oauth_timestamp"] = oauth_timestamp
	apiValues["oauth_version"] = "1.0"
	apiValues["search_expression"] = strings.Replace(query, " ", "+", -1)

	paramStr := ""

	for k, v := range apiValues {
		paramStr = paramStr + "&" + k + "=" + url.QueryEscape(v)
	}

	paramStr = strings.TrimPrefix(paramStr, "&")

	sigBaseStr := "POST&" + url.QueryEscape(fatSecretUrl) + "&" + url.QueryEscape(paramStr)
	sharedSecret := fatSecretAccessSecret + "&"

	oauth_signature := Sign(sigBaseStr, sharedSecret)

	fmt.Println("Signature: " + oauth_signature)

	body, err := Get(fatSecretUrl, apiValues)
	return body, err
}
