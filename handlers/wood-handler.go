package handlers

import (
    "encoding/json"
    "net/http"
    "fmt" 
    "log"

    "github.com/ristirahva/rest-app/services"
)

type WoodHandler struct {
    woodService services.WoodService
}

func NewWoodHandler(woodService services.WoodService) *WoodHandler {
    return &WoodHandler{
        woodService:        woodService,
    }
}

func (h *WoodHandler) GetWoods(w http.ResponseWriter, r *http.Request) {
    log.Println("вызван обработчик GetWoods")
    woods, err := h.woodService.GetAllWoods()

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
    
    json.NewEncoder(w).Encode(woods)
/*
    sendJSON(w, http.StatusOK, map[string]interface{}{
        "status": "success",
        "data": woods,
        "count": len(woods),
    })
*/
}

func AddWood(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    log.Println("вызван обработчик PostWoods")
    fmt.Fprintf(w, `{"message":"Hello, добавить материал!"}`)
}


func UpdateWood(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    log.Println("вызван обработчик PutWoods")
    fmt.Fprintf(w, `{"message":"Hello, изменить материал!"}`)
}


func DeleteWood(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNoContent)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    log.Println("вызван обработчик DeleteWoods")
    fmt.Fprintf(w, `{"message":"Hello, удалить материал!"}`)
}
