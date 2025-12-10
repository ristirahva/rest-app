package services
/*
import (
    "context"
    "fmt"
    "github.com/ristirahva/rest-app/db"
    "github.com/ristirahva/rest-app/repositories"
)

type DrinkService struct {
    BaseService[db.Drink]
    repo *repositories.DrinkRepository
}

// Создание сервиса

func NewDrinkService(repo *repositories.DrinkRepository) *DrinkService {
    return &DrinkService{
        BaseService: *NewBaseService[db.Drink](&repo.BaseRepository),
        repo:        repo,
    }
}

// Создание нового напитка
//
// Параметры:
//
// контекст
// начальная крепость
// описание

func (s *DrinkService) CreateDrink(ctx context.Context, alcohol *int, description string) (*db.Drink, error) {
    if alcohol != nil && (*alcohol < 0 || *alcohol > 96) {
        return nil, fmt.Errorf("Крепость напитка должна быть от 0 до 96 градусов")
    }
    
    drink := &db.Drink{
        Alcohol:     alcohol,
        Description: description,
    }
    
    if err := s.repo.Create(drink); err != nil {
        return nil, fmt.Errorf("Невозможно создать новый напиток: %w", err)
    }
    
    return drink, nil
}

// Список наличных напитков в бочках
//
// Параметр: контекст

func (s *DrinkService) GetDrinksInBarrels(ctx context.Context) ([]db.Drink, error) {
    return s.repo.FindCurrentlyInBarrels()
}

// История напитка по бочкам
//
// Параметры:
//
// контекст
// идентификатор напитка

func (s *DrinkService) GetDrinkHistory(ctx context.Context, drinkID uint) ([]db.DrinkInBarrel, error) {
    return s.repo.GetDrinkHistory(drinkID)
}

// Список бочек, в которых был напиток
//
// Параметры:
//
// контекст
// идентификатор напитка

func (s *DrinkService) GetDrinkBarrels(ctx context.Context, drinkID uint) ([]db.Barrel, error) {
    drink, err := s.GetByID(ctx, drinkID, "Barrels", "Barrels.Wood")
    if err != nil {
        return nil, fmt.Errorf("Невозможно найти напиток: %w", err)
    }
    return drink.Barrels, nil
}
*/