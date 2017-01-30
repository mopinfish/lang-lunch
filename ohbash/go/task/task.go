package task

import (
  "net/url"
  "log"
  "../database"
)

func GetAll() ([]string) {
  var tasks []string
  tasks = database.GetAllTasks()
  return tasks
}

func AddTask(reqParams url.Values) () {
  log.Print(reqParams["task"][0])
}
