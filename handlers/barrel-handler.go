package handlers

import (
    "encoding/json"
    "net/http"
    "fmt" 
    "log"

    "github.com/ristirahva/rest-app/services"
)

type BarrelHandler struct {
    barrelService services.BarrelService
}

func NewBarrelHandler(barrelService services.BarrelService) *BarrelHandler {
    return &BarrelHandler{
        barrelService:        barrelService,
    }
}

func (h *BarrelHandler) GetBarrels(w http.ResponseWriter, r *http.Request) {
    log.Println("вызван обработчик GetBarrels")
    barrels, err := h.barrelService.GetAllBarrels()

    if (err != nil) { 
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(ApiResponse{
            Success: false,
            Error:   "Внутренняя ошибка сервера",
            })
            // В production лучше не показывать детали ошибки клиенту
        return
    }    

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    
    json.NewEncoder(w).Encode(barrels)
/*
    sendJSON(w, http.StatusOK, map[string]interface{}{
        "status": "success",
        "data": barrels,
        "count": len(barrels),
    })
*/
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
    w.WriteHeader(http.StatusNoContent)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    log.Println("вызван обработчик DeleteBarrels")
    fmt.Fprintf(w, `{"message":"Hello, удалить бочку!"}`)
}
