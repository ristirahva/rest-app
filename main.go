package main
import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"

    "github.com/ristirahva/rest-app/config"
    "github.com/ristirahva/rest-app/handlers"
)
                                                                                           
func main() {

    cfg, err := config.LoadConfig("config/config.json")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
  
    fmt.Printf("Server host: %s\n", cfg.Server.Host)
    fmt.Printf("Database name: %s\n", cfg.Database.Name)

    router := mux.NewRouter()
    api := router.PathPrefix("/api/v1/barrel-aging").Subrouter()
    
    api.HandleFunc("/barrel", handlers.GetBarrels).Methods("GET")
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
