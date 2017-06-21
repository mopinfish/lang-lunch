package main

import (
  "os"
  "fmt"
  "strconv"
  "net/http"
  "github.com/m0a/easyjson"
  chatwork "github.com/griffin-stewie/go-chatwork"
)

func get() (string, error) {
  var err error
  url := os.Getenv("RA_API_URL")
  req, _ := http.NewRequest("GET", url, nil)

  // setup ra api client
  client := new(http.Client)
  response, res_err := client.Do(req)
  if res_err != nil {
    fmt.Println(res_err)
    err = fmt.Errorf("Invalid responses")
  }

  defer response.Body.Close()
  jsonData, err := easyjson.NewEasyJson(response.Body)

  ra_count, count_err := jsonData.K("result").K("result_set").K("total_hits").AsInt()
  if count_err != nil {
    fmt.Println(count_err)
    err = fmt.Errorf("Invalid json data")
  }

  return strconv.Itoa(ra_count), err
}

func makeBody (count string) string {
  var body string = "[info][title]投資反響レポート from GoogleAnalytics[/title]※ 本メール大塚のPCより、毎日自動で送信されます。\n\n本日の投資掲載物件数は【" + count + "】件でした。\n明日も１日頑張りましょう!!![/info]"
  return body
}

func send(message string) error {
  token := os.Getenv("CHATWORK_TOKEN")
  chatwork := chatwork.NewClient(token)
//  endpoint := "/rooms/74123818/messages"
  endpoint := "/rooms/18613296/messages"
  params :=  map[string]string{"body": message}
  chatwork.Post(endpoint, params)
  return nil
}

func main() {
  count, err := get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
  body := makeBody(count)
  if err := send(body); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
}
