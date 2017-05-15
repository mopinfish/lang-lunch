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
  var err error
//  url := "https://cloud.feedly.com/v3/streams/contents?streamId=user/" + user_id + "/tag/global.saved&count=9"
  url := "https://cloud.feedly.com/v3/streams/contents?streamId=user/" + user_id + "/tag/Read%20For%20Lator&count=90"
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
    url, url_err := jsonData.K("items").K(k).K("alternate").RangeObjects()[0].K("href").AsString()
fmt.Println(url)
    if url_err != nil {
      err = fmt.Errorf("Invalid responses")
    }
    item := map[string]string{"title": title, "url": url}
    data = append(data, item)
  }

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
  for k, item := range data {
    if k > 3 {
      break
    }
    body += item["title"] + "\n" + item["url"] + "\n\n"
  }
  return body
}

func send(message string) error {
  token := os.Getenv("CHATWORK_TOKEN")
  chatwork := chatwork.NewClient(token)
  endpoint := "/rooms/74123818/messages"
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
