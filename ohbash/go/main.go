package main

import (
  "net/http"
  "text/template"
  "log"
  "./task"
  "./database"
)

type Page struct {
  Title string
  Tasks []string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  tasks := task.GetAll()
  page := Page{"ToDo Application", tasks}
  tmpl, err := template.ParseFiles("layout.html")
  if err != nil {
    panic(err)
  }

  err = tmpl.Execute(w, page)
  if err != nil{
    panic(err)
  }
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
  task.AddTask(r.URL.Query())
  tasks := task.GetAll()
  page := Page{"Add Task", tasks}
  tmpl, err := template.ParseFiles("layout.html")
  if err != nil {
    panic(err)
  }

  err = tmpl.Execute(w, page)
  if err != nil{
    panic(err)
  }
}

func main () {
  log.Print("start")
  db := database.myDb {}
  db.Connect()
  defer db.Close()

  http.HandleFunc("/", viewHandler)
  http.HandleFunc("/addTask", addTaskHandler)
  http.ListenAndServe("localhost:8080", nil)

}
