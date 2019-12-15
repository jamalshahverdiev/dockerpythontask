package main

import (
    "fmt"
    "log"
    "net/http"
)

func rootPage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Success!")
    fmt.Println("Endpoint Hit: rootPage")
}

func pingPage(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Ok")
     fmt.Println("Endpoint Hit: pingPage")
}


func handleRequests() {
    http.HandleFunc("/", rootPage)
    http.HandleFunc("/ping", pingPage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
