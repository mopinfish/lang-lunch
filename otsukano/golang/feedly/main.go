package main

import (
  "os"
  "fmt"
  "net/http"
  "github.com/m0a/easyjson"
  "math/rand"
  chatwork "github.com/yoppi/go-chatwork"
)

var api = "https://teratail.com/api/v1"

var token = os.Getenv("FEEDLY_TOKEN")
var user_id = os.Getenv("FEEDLY_USER_ID")

func get() ([]map[string]string, error) {
//  url := "http://cloud.feedly.com/v3/profile"
  var err error
  url := "https://cloud.feedly.com/v3/streams/contents?streamId=user/" + user_id + "/tag/global.saved&count=19"
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
//    url, url_err := jsonData.K("items").K(k).K("visual").K("url").AsString()
    url, url_err := jsonData.K("items").K(k).K("visual").K("url").AsString()
    if url_err != nil {
      err = fmt.Errorf("Invalid responses")
    }
    item := map[string]string{"title": title, "url": url}
    data = append(data, item)
  }

  fmt.Println(data)

  return data, err
}

func shuffle (list []map[string]string) []map[string]string {
  for i := len(list); i > 1; i-- {
    j := rand.Intn(i)          // 0～(i-1) の乱数発生
    list[i - 1], list[j] = list[j], list[i - 1]
  }
  return list
}

func makeBody (data []map[string]string) string {
  var body string
  for _, item := range data {
//    if k > 3 {
//      break
//    }
    body += item["title"] + "\n" + item["url"] + "\n\n"
  }
  return body
}

func send(message string) error {
  token := os.Getenv("CHATWORK_TOKEN")
  chatwork := chatwork.NewClient(token)
  endpoint := "/rooms/361427/messages"
  params :=  map[string]string{"body": message}
  chatwork.Post(endpoint, params)
  return nil
}

func main() {
  data, err := get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
  shuf_data := shuffle(data)
  body := makeBody(shuf_data)
  if err := send(body); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
}
