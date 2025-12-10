package db

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

    "github.com/ristirahva/rest-app/config"
)

func DbConnect() (*gorm.DB, error) {
    cfg, err := config.LoadConfig("config/config.json")            
    if err != nil {
        log.Fatalf("Невозможно прочитать конфигурационные настройки: %v", err)
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.Database.Host,
        cfg.Database.Username,
        cfg.Database.Password,
        cfg.Database.Name,
        cfg.Database.Port,
    )
//    var err error
    db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if dbErr != nil {
        log.Fatal("Невозможно подключиться к БД:", err)
        return nil, dbErr
    }
    // Автоматическая миграция схемы
/*
    if err := DB.AutoMigrate(&models.Item{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
*/
    
    log.Println("Подключение к БД установлено")
    return db, nil
}