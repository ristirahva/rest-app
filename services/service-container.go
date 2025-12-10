package services

/*
import (
    "gorm.io/gorm"
    "github.com/ristirahva/rest-app/repositories"
)

// Контейнер сервисов

type ServiceContainer struct {
    WoodService           *WoodService
    BarrelService         *BarrelService
    DrinkService          *DrinkService
    BarrelManagementService *BarrelManagementService
}

func NewServiceContainer(db *gorm.DB) *ServiceContainer {
    // Инициализация репозиториев
    woodRepo := repositories.NewWoodRepository(db)
    barrelRepo := repositories.NewBarrelRepository(db)
    drinkRepo := repositories.NewDrinkRepository(db)
    dibRepo := repositories.NewDrinkInBarrelRepository(db)
    
    // Инициализация сервисов
    woodService := NewWoodService(woodRepo)
    barrelService := NewBarrelService(barrelRepo)
    drinkService := NewDrinkService(drinkRepo)
    barrelManagementService := NewBarrelManagementService(
        barrelRepo,
        drinkRepo,
        dibRepo,
        woodRepo,
    )
    
    return &ServiceContainer{
        WoodService:           woodService,
        BarrelService:         barrelService,
        DrinkService:          drinkService,
        BarrelManagementService: barrelManagementService,
    }
}
*/