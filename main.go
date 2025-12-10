package main
import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"

    "github.com/ristirahva/rest-app/config"
    "github.com/ristirahva/rest-app/handlers"
    "github.com/ristirahva/rest-app/services"
    "github.com/ristirahva/rest-app/repositories"
    "github.com/ristirahva/rest-app/db"
)
                                                                                           
func main() {

    cfg, err := config.LoadConfig("config/config.json")            
    if err != nil {
        log.Fatalf("Невозможно прочитать конфинурационные настройки: %v", err)
    }
  
    fmt.Printf("Server host: %s\n", cfg.Server.Host)
    fmt.Printf("Database name: %s\n", cfg.Database.Name)

    router := mux.NewRouter()
    api := router.PathPrefix("/api/v1/barrel-aging").Subrouter()

    dbConn, dbErr := db.DbConnect()
    if (dbErr != nil) {
        log.Fatal("База данных недоступна: ", dbErr)
        return
    } 
    barrelRepository := repositories.NewBarrelRepository(dbConn)
    drinkInBarrelRepository := repositories.NewDrinkInBarrelRepository(dbConn)
    barrelService := services.NewBarrelService(barrelRepository, drinkInBarrelRepository) 
    barrelHandler := handlers.NewBarrelHandler(*barrelService)
    
    api.HandleFunc("/barrel", barrelHandler.GetBarrels).Methods("GET")
    api.HandleFunc("/barrel", handlers.AddBarrel).Methods("POST")
    api.HandleFunc("/barrel", handlers.UpdateBarrel).Methods("PUT")
    api.HandleFunc("/barrel", handlers.DeleteBarrel).Methods("DELETE")

    server := &http.Server{
        Addr: ":" + cfg.Server.Port,
        Handler: router,
    }
    
    fmt.Printf("Server is running on http://localhost:%s\n", cfg.Server.Port)
    
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal("Ошибка сервера: ", err)
    }
}
