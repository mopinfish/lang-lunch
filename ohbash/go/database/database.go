package database

import (
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
)

type MyDb struct {
  db *sql.DB
}

func Connect(this *MyDb) {
  this.db, err = sql.Open("mysql", "root@/gosample")

  if err != nil {
    panic(err.Error())
  }
  defer this.db.Close()
}

func GetAllTasks() ([]string) {
  //db, err := sql.Open("mysql", "root@/gosample")

  //if err != nil {
  //  panic(err.Error())
  //}
  //defer db.Close()

  rows, err := this.db.Query("SELECT * FROM tasks")
  defer rows.Close()

  if err != nil {
    panic(err.Error())
  }
  var tasks []string

  for rows.Next() {
    var id int
    var name string
    var status int
    if err := rows.Scan(&id, &name, &status); err != nil {
      panic(err.Error())
    }
    tasks = append(tasks, name)
  }

  return tasks
}
