package services

import (
    "context"
    "fmt"
    "time"
    "github.com/ristirahva/rest-app/db"
    "github.com/ristirahva/rest-app/repositories"
)

// Сервис работы с бочками

type BarrelManagementService struct {
    barrelRepo *repositories.BarrelRepository
    drinkRepo  *repositories.DrinkRepository
    dibRepo    *repositories.DrinkInBarrelRepository
    woodRepo   *repositories.WoodRepository
}

func NewBarrelManagementService(
    barrelRepo *repositories.BarrelRepository,
    drinkRepo *repositories.DrinkRepository,
    dibRepo *repositories.DrinkInBarrelRepository,
    woodRepo *repositories.WoodRepository,
) *BarrelManagementService {
    return &BarrelManagementService{
        barrelRepo: barrelRepo,
        drinkRepo:  drinkRepo,
        dibRepo:    dibRepo,
        woodRepo:   woodRepo,
    }
}

// TransferDrink переливает напиток из одной бочки в другую
func (s *BarrelManagementService) TransferDrink(
    ctx context.Context,
    fromBarrelID, toBarrelID uint,
    description string,
) error {
    // Проверяем, что исходная бочка занята
    activeBarrels, err := s.barrelRepo.FindNonEmptyBarrels()
    if err != nil {
        return fmt.Errorf("failed to check barrels: %w", err)
    }
    
    var currentDrinkID uint
    fromBarrelFound := false
    
    for _, b := range activeBarrels {
        if b.ID == fromBarrelID {
            if len(b.Drinks) > 0 {
                currentDrinkID = b.Drinks[0].ID
                fromBarrelFound = true
                break
            }
        }
    }
    
    if !fromBarrelFound {
        return fmt.Errorf("Исходная бочка пуста")
    }
    
    // Проверяем, что целевая бочка свободна
    for _, b := range activeBarrels {
        if b.ID == toBarrelID {
            return fmt.Errorf("Целевая бочка уже залита")
        }
    }
    
    // Освобождаем исходную бочку
    now := time.Now()
    if err := s.barrelRepo.RemoveDrinkFromBarrel(fromBarrelID, currentDrinkID, now, nil); err != nil {
        return fmt.Errorf("Невозможно освободить исходную бочку: %w", err)
    }
    
    // Заполняем целевую бочку
    if err := s.barrelRepo.AddDrinkToBarrel(toBarrelID, currentDrinkID, now, nil, description); err != nil {
        return fmt.Errorf("Невозможно залить целевую бочку: %w", err)
    }
    
    return nil
}

// GetBarrelStatusReport возвращает отчет по статусу всех бочек
func (s *BarrelManagementService) GetBarrelStatusReport(ctx context.Context) (map[string]interface{}, error) {
    // Получаем все бочки
    allBarrels, err := s.barrelRepo.FindAll()
    if err != nil {
        return nil, fmt.Errorf("Невозможно получить отчёт по бочкам: %w", err)
    }
    
    // Получаем занятые бочки
    inUseBarrels, err := s.barrelRepo.FindBarrelsWithCurrentDrink()
    if err != nil {
        return nil, fmt.Errorf("Невозможно получить список занятых бочек: %w", err)
    }
    
    // Получаем свободные бочки
    emptyBarrels, err := s.barrelRepo.FindEmptyBarrels()
    if err != nil {
        return nil, fmt.Errorf("Невозможно получить список свободных бочек: %w", err)
    }
    
    // Считаем общий объем
    totalVolume := 0
    usedVolume := 0
    for _, b := range allBarrels {
        totalVolume += b.Volume
    }
    for _, b := range inUseBarrels {
        usedVolume += b.Volume
    }
    
    return map[string]interface{}{
        "total_barrels":    len(allBarrels),
        "in_use_barrels":   len(inUseBarrels),
        "empty_barrels":    len(emptyBarrels),
        "total_volume":     totalVolume,
        "used_volume":      usedVolume,
        "available_volume": totalVolume - usedVolume,
        "utilization_rate": float64(usedVolume) / float64(totalVolume) * 100,
    }, nil
}

// Отчет по использованию пород дерева
func (s *BarrelManagementService) GetWoodUtilizationReport(ctx context.Context) ([]map[string]interface{}, error) {
    return s.woodRepo.GetWoodBarrelStatistics()
}

// Отчет по выдержке напитков
func (s *BarrelManagementService) GetAgingReport(ctx context.Context) ([]map[string]interface{}, error) {
    var result []map[string]interface{}
    
    // Получаем актуальные заполнения бочек
    activeOccupations, err := s.dibRepo.FindActiveBarrelOccupations()
    if err != nil {
        return nil, fmt.Errorf("Невозможно найти актуальные заполнения бочек: %w", err)
    }
    
    for _, occupation := range activeOccupations {
        duration := time.Since(occupation.DateStart)
        
        result = append(result, map[string]interface{}{
            "barrel_id":       occupation.BarrelID,
            "drink_id":        occupation.DrinkID,
            "drink_name":      occupation.Drink.Description,
            "start_date":      occupation.DateStart,
            "aging_days":      int(duration.Hours() / 24),
            "wood_type":       occupation.Barrel.Wood.Name,
            "barrel_volume":   occupation.Barrel.Volume,
        })
    }
    
    return result, nil
}