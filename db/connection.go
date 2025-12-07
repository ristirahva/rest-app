package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

    "github.com/ristirahva/rest-app/config"
)
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.Database.Host,
        cfg.Database.Username,
        cfg.Database.Password,
        cfg.Database.Name,
        cfg.Database.Port,
    )
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    // Автоматическая миграция схемы
    if err := DB.AutoMigrate(&models.Item{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    log.Println("Database connected successfully")