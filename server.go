package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
)

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/object", object)
    http.ListenAndServe(":8080", nil)
}

func object(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.Method, r.URL, r.Proto)

  type Obj struct {
    Text string
    Centstatus int
    Id string
    Level int
  }
  m := Obj{"Подмостки (д)", 1, "j7d39dc2-cd93-466c-8389-4d6b56f017cb", 6}
  b, err := json.Marshal(m)
  if err != nil {log.Fatal(err)}
  fmt.Fprintf(w, "%s", b)
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.Method, r.URL, r.Proto)
  fmt.Fprintf(w, "<h1>Hello, world</h1><br/><img src='http://golang.org/doc/gopher/frontpage.png'><img src='http://golang.org/doc/gopher/frontpage.png'>")
}
