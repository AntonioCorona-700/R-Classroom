package main

import (
   "fmt"
   "log"
   "net/http"
   "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("I got a request!")
   fmt.Fprintf(w, "Hi there :)")
}
func byeHandler(w http.ResponseWriter, r *http.Request) {
   parts := strings.Split(r.URL.Path, "/")
   name := "human"
   if len(parts) > 2 {
       name = parts[2]
   }
   fmt.Fprintf(w, "Bye "+name)
}

func main() {
   http.HandleFunc("/hello", handler)
   http.HandleFunc("/bye", byeHandler)
   log.Fatal(http.ListenAndServe(":8080", nil))
   

}

