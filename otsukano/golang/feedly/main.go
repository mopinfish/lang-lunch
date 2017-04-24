package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/m0a/easyjson"
)

var api = "https://teratail.com/api/v1"

var token = os.Getenv("FEEDLY_TOKEN")

func get() error {
  url := "http://cloud.feedly.com/v3/profile"
  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", token)

  client := new(http.Client)
  response, err := client.Do(req)

  if err != nil {
    return fmt.Errorf("Invalid responses")
  }
  defer response.Body.Close()
  jsonData, err := easyjson.NewEasyJson(response.Body)
  if err != nil {
    return fmt.Errorf("Invalid responses")
  }
  fmt.Println(jsonData)
  return nil
}

func run() error {
  response, err := http.Get(api + "/questions")
  if err != nil {
    return fmt.Errorf("Failed to connect teratail.com")
  }
  defer response.Body.Close()

  jsonData, err := easyjson.NewEasyJson(response.Body)
  if err != nil {
    return fmt.Errorf("Invalid responses")
  }

  for _, v:=range jsonData.K("questions").RangeObjects() {
    fmt.Printf("%s\n", v.K("title"))
  }

  return nil
}

func main() {
  if err := get(); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
}
