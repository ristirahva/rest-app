package services

import (
    "context"
    "fmt"
    "time"

    "github.com/ristirahva/rest-app/config"
    "github.com/ristirahva/rest-app/db"
    "github.com/ristirahva/rest-app/repositories"
)

type BarrelService struct {
    BaseService[db.Barrel]
    repo *repositories.BarrelRepository
}

func NewBarrelService(repo *repositories.BarrelRepository) *BarrelService {
    return &BarrelService{
        BaseService: *NewBaseService[db.Barrel](&repo.BaseRepository),
        repo:        repo,
    }
}

// CreateBarrel создает новую бочку с валидацией по минимальному и максимальному объёмам
func (s *BarrelService) CreateBarrel(ctx context.Context, woodID *uint, volume int, description string) (*db.Barrel, error) {
    cfg, err := config.LoadConfig("config/config.json")
    if err != nil {
        log.Fatalf("Невозможно прочитать конфигурационный файл: %v", err)
    }

    if volume < cfg.MinBarrelCapacity {
        return nil, fmt.Errorf("Объём бочки не может быть меньше %d", cfg.MinBarrelCapacity)
    }
    
    if volume > cfg.MaxBarrelCapacity {
        return nil, fmt.Errorf("Объём бочки не может быть больше %d", cfg.MaxBarrelCapacity)
    }
    
    barrel := &db.Barrel{
        WoodID:      woodID,
        Volume:      volume,
        Description: description,
    }
    
    if err := s.repo.Create(barrel); err != nil {
        return nil, fmt.Errorf("Невозможно добавить бочку: %w", err)
    }
    
    return barrel, nil
}

// GetBarrelsByWood возвращает бочки по породе дерева
//
// параметры:
// 
// контекст
// идентификатор материала

func (s *BarrelService) GetBarrelsByWood(ctx context.Context, woodID uint) ([]db.Barrel, error) {
    return s.repo.FindByWood(woodID)
}

// GetAvailableBarrels возвращает свободные бочки
//
// параметр: контекст
//
// возврат: список пустых бочек

func (s *BarrelService) GetAvailableBarrels(ctx context.Context) ([]db.Barrel, error) {
    return s.repo.FindEmptyBarrels()
}

// GetBarrelsInUse возвращает занятые бочки
//
// параметр: контекст
//
// возврат: список занятых бочек
func (s *BarrelService) GetBarrelsInUse(ctx context.Context) ([]db.Barrel, error) {
    return s.repo.FindNonEmptyBarrels()
}

// FillBarrel заполняет бочку напитком
//
// Параметры:
//
// контекст
// идентификатор бочки
// идентификатор дистиллята
// начальная крепость
// описание
//
// Возврат: ошибка

func (s *BarrelService) FillBarrel(ctx context.Context, barrelID, drinkID uint, alcoholStart *int, description string) error {
    // Проверяем, свободна ли бочка
    // TODO проверить заполненность бочки без перебора с выдачей, каким именно
    activeBarrels, err := s.repo.FindNonEmptyBarrels()
    if err != nil {
        return fmt.Errorf("Невозможно проверить заполненность бочек: %w", err)
    }
    
    for _, b := range activeBarrels {
        if b.ID == barrelID {
            return fmt.Errorf("Бочка уже заполнена дистиллятом")
        }
    }
    
    // Заполняем бочку
    return s.repo.AddDrinkToBarrel(barrelID, drinkID, time.Now(), alcoholStart, description)
}

// Слив дистиллята из бочки
//
// Параметры:
//
// контекст
// идентификатор бочки
// итоговая крепость дистиллята
//
// Возврат: ошибка

func (s *BarrelService) EmptyBarrel(ctx context.Context, barrelID uint, alcoholEnd *int) error {
    // Находим текущий напиток в бочке
    // TODO сделать без перебора
    activeBarrels, err := s.repo.FindNonEmptyBarrels()
    if err != nil {
        return fmt.Errorf("Невозможно проверить заполненность бочек: %w", err)
    }
    
    var drinkID uint
    found := false
    for _, b := range activeBarrels {
        if b.ID == barrelID {
            if len(b.Drinks) > 0 {
                drinkID = b.Drinks[0].ID
                found = true
                break
            }
        }
    }
    
    if !found {
        return fmt.Errorf("Бочка пуста")
    }
    
    return s.repo.RemoveDrinkFromBarrel(barrelID, drinkID, time.Now(), alcoholEnd)
}

// История использования бочки
//
// Параметры: 
//
// контекст
// идентификатор бочки

func (s *BarrelService) GetBarrelHistory(ctx context.Context, barrelID uint) ([]db.DrinkInBarrel, error) {
    dibRepo := repositories.NewDrinkInBarrelRepository(s.repo.BaseRepository.Db)
    return dibRepo.GetBarrelHistory(barrelID)
}

