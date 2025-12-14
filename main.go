package main
import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "gorm.io/gorm"

    "github.com/ristirahva/rest-app/config"
    "github.com/ristirahva/rest-app/handlers"
    "github.com/ristirahva/rest-app/services"
    "github.com/ristirahva/rest-app/repositories"
    "github.com/ristirahva/rest-app/db"
)
                                                                                           
func main() {

    cfg, err := config.LoadConfig("config/config.json")            
    if err != nil {
        log.Fatalf("Невозможно прочитать конфигурационные настройки: %v", err)
    }
  
    fmt.Printf("Server host: %s\n", cfg.Server.Host)
    fmt.Printf("Database name: %s\n", cfg.Database.Name)

    router := mux.NewRouter()
    api := router.PathPrefix("/api/v1/barrel-aging").Subrouter()

    dbConn, dbErr := db.DbConnect()
    if (dbErr != nil) {
        //sqlDB, err := db.DB()
        log.Fatal("База данных недоступна: ", dbErr)
        //defer sqlDB.Close()
        return
    } 

    woodRepository, barrelRepository, drinkInBarrelRepository := initRepositories(dbConn)

    barrelService := services.NewBarrelService(barrelRepository, drinkInBarrelRepository) 
    woodService := services.NewWoodService(woodRepository)

    barrelHandler := handlers.NewBarrelHandler(*barrelService)
    woodHandler := handlers.NewWoodHandler(*woodService)

    api.HandleFunc("/barrel", barrelHandler.GetBarrels).Methods("GET")
    api.HandleFunc("/barrel", handlers.AddBarrel).Methods("POST")
    api.HandleFunc("/barrel", handlers.UpdateBarrel).Methods("PUT")
    api.HandleFunc("/barrel", handlers.DeleteBarrel).Methods("DELETE")

    api.HandleFunc("/wood", woodHandler.GetWoods).Methods("GET")

    server := &http.Server{
        Addr: ":" + cfg.Server.Port,
        Handler: router,
    }
        
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal("Ошибка сервера: ", err)
        return;
    }
    fmt.Printf("Сервер запущен на http://localhost:%s\n", cfg.Server.Port)
}

func initRepositories(dbConn *gorm.DB) (*repositories.WoodRepository, *repositories.BarrelRepository, *repositories.DrinkInBarrelRepository) {
    woodRepository := repositories.NewWoodRepository(dbConn)
    barrelRepository := repositories.NewBarrelRepository(dbConn)
    drinkInBarrelRepository := repositories.NewDrinkInBarrelRepository(dbConn)
    return woodRepository, barrelRepository, drinkInBarrelRepository
}

