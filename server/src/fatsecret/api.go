package fatsecret

import (
  "fmt"
  "io"
  "strconv"
  "strings"
  "regexp"
  "time"
  "io/ioutil"
  "net/http"
  "net/url"
  "crypto/hmac"
  "crypto/sha1"
  "encoding/base64"
  "encoding/xml"

  "github.com/dchest/uniuri"
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

func SearchFood(query string) (*Foods, error) {
  result, err := Query(query)
  if err != nil {
    return nil, err
  }

  foods, err := ParseFoods(result)
  if err != nil {
    return nil, err
  }

  return &foods, nil
}

func ParseFoods(b []byte) (Foods, error) {
  var q Foods
  xml.Unmarshal(b, &q)

  return q, nil
}

// This function fetch the content of a URL will return it as an
// array of bytes if retrieved successfully.
func Query(query string) ([]byte, error) {
  fatSecretUrl := "http://platform.fatsecret.com/rest/server.api"
  fatSecretConsumerKey := "62cc7c5caaf542668006fc70cbfdabae"
  // fatSecretAccessSecret := "de666f86e8634a77947c02fc39cf33cd"

  oauth_nonce := uniuri.New()
  oauth_timestamp := strconv.FormatInt(time.Now().Unix(), 10)

  reg, err := regexp.Compile("[^a-z]")
  if err != nil {
    return nil, err
  }

  oauth_nonce = reg.ReplaceAllString(oauth_nonce, "")

  apiValues := make(map[string]string)
  apiValues["method"] = "foods.search"
  apiValues["oauth_consumer_key"] = fatSecretConsumerKey
  apiValues["oauth_nonce"] = oauth_nonce
  apiValues["oauth_signature_method"] = "HMAC-SHA1"
  apiValues["oauth_timestamp"] = oauth_timestamp
  apiValues["oauth_version"] = "1.0"
  apiValues["search_expression"] = query

  paramStr := ""

  for k, v := range apiValues {
    paramStr = paramStr + "&" + k + "=" + url.QueryEscape(v)
  }

  paramStr = strings.TrimPrefix(paramStr, "&")

  sigBaseStr := "POST&" + url.QueryEscape(fatSecretUrl) + "&" + url.QueryEscape(paramStr)

  fmt.Println(sigBaseStr)

  sharedSecret := "de666f86e8634a77947c02fc39cf33cd&" // fatSecretConsumerKey + "&" + fatSecretAccessSecret

  hasher := hmac.New(sha1.New, []byte(sharedSecret))
  io.WriteString(hasher, sigBaseStr)

  oauth_signature := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

  // Build the request
  values := url.Values{"method": {"foods.search"},
    "oauth_consumer_key": {fatSecretConsumerKey},
    "oauth_nonce": {oauth_nonce},
    "oauth_signature_method": {"HMAC-SHA1"},
    "oauth_timestamp": {oauth_timestamp},
    "oauth_version": {apiValues["oauth_version"]},
    "search_expression": {apiValues["search_expression"]},
    "oauth_signature": {oauth_signature}}

  resp, err := http.PostForm(fatSecretUrl, values)
  if err != nil {
    return nil, err
  }
  // Defer the closing of the body
  defer resp.Body.Close()
  // Read the content into a byte array
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  // At this point we're done - simply return the bytes
  return body, nil
}
