package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/m0a/easyjson"
  chatwork "github.com/yoppi/go-chatwork"
)

var api = "https://teratail.com/api/v1"

var token = os.Getenv("FEEDLY_TOKEN")
var user_id = os.Getenv("FEEDLY_USER_ID")

func get() ([]map[string]string, error) {
//  url := "http://cloud.feedly.com/v3/profile"
  var err error
  url := "https://cloud.feedly.com/v3/streams/contents?streamId=user/" + user_id + "/tag/global.saved&count=1"
  req, _ := http.NewRequest("GET", url, nil)
  req.Header.Set("Authorization", token)

  client := new(http.Client)
  response, res_err := client.Do(req)

  if res_err != nil {
    err = fmt.Errorf("Invalid responses")
  }
  defer response.Body.Close()
  jsonData, err := easyjson.NewEasyJson(response.Body)
  if res_err != nil {
    err = fmt.Errorf("Invalid responses")
  }


  var data []map[string]string

  for k,_ := range jsonData.K("items").RangeObjects() {
    title, title_err := jsonData.K("items").K(k).K("title").AsString()
    if title_err != nil {
      err = fmt.Errorf("Invalid responses")
    }
    url, url_err := jsonData.K("items").K(k).K("visual", "url").AsString()
    if url_err != nil {
      err = fmt.Errorf("Invalid responses")
    }
    item := map[string]string{"title": title, "url": url}
    data = append(data, item)
  }

  fmt.Println(data)
//  //string value
//  copyrights,err:=json.K("routes").K(0).K("copyrights").AsString()
//  if err!=nil {
//    panic("AsString err")
//  }



  return data, err
}

func send(message string) error {
  token := os.Getenv("CHATWORK_TOKEN")
  chatwork := chatwork.NewClient(token)
  endpoint := "/rooms/361427/messages"
  params :=  map[string]string{"body": message}
  chatwork.Post(endpoint, params)
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
  data, err := get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
  fmt.Println(data)
//  if err := send("hogehoge"); err != nil {
//    fmt.Fprintf(os.Stderr, "%s\n", err)
//    os.Exit(1)
//  }
}
