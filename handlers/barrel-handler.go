package handlers

import (
    "net/http"
    "fmt" 
    "log"
)

func GetBarrels(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    log.Println("вызван обработчик GetBarrels")
    fmt.Fprintf(w, `{"message":"список бочек!"}`)
}

func AddBarrel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    log.Println("вызван обработчик PostBarrels")
    fmt.Fprintf(w, `{"message":"Hello, добавить бочку!"}`)
}


func UpdateBarrel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    log.Println("вызван обработчик PutBarrels")
    fmt.Fprintf(w, `{"message":"Hello, изменить бочку!"}`)
}


func DeleteBarrel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
//    w.WriteHeader(http.StatusNoContent)
    w.WriteHeader(http.StatusOK)
    log.Println("вызван обработчик DeleteBarrels")
    fmt.Fprintf(w, `{"message":"Hello, удалить бочку!"}`)
}
