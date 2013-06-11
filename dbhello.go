package main

import (
  "database/sql"
  "net/http"
  "fmt"
  "log"
)

import (
  _ "github.com/lib/pq"
)

func main() {
  http.HandleFunc("/", db_read)
  http.ListenAndServe(":8090", nil)
}

func db_read(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("postgres", "user=dev dbname=test_db password=dev")

  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()
  var msg string
  err = db.QueryRow("SELECT did || ' ' || name FROM distributors;").Scan(&msg)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Fprintf(w, msg)
}