package services

import (
//    "fmt"
    "log"

    "github.com/ristirahva/rest-app/db"
    "github.com/ristirahva/rest-app/repositories"
)

type WoodService struct {
    repo *repositories.WoodRepository
}

func NewWoodService(repo *repositories.WoodRepository) *WoodService {
    return &WoodService{
        repo:        repo,
    }
}

// GetAllWoods список всех материалов

func (s *WoodService) GetAllWoods() ([]db.Wood, error) {
    woods, error := s.repo.FindAll()
    log.Printf("Сервис Список материалов: {}", woods)
    return woods, error
}

/*
type WoodService struct {
    BaseService[db.Wood]
    repo *repositories.WoodRepository
}

// создание сервиса

func NewWoodService(repo *repositories.WoodRepository) *WoodService {
    return &WoodService{
        BaseService: *NewBaseService[db.Wood](&repo.BaseRepository),
        repo:        repo,
    }
}

// поиск породы дерева по названию
func (s *WoodService) GetByName(ctx context.Context, name string) (*db.Wood, error) {
    wood, err := s.repo.FindByName(name)
    if err != nil {
        return nil, fmt.Errorf("Такой материал для бочки не найден: %w", err)
    }
    return wood, nil
}

// Список всех материалов со сделанными из них бочками
//
// Параметр: контекст

func (s *WoodService) GetWoodsWithBarrels(ctx context.Context) ([]db.Wood, error) {
    return s.repo.FindWithBarrels()
}

// Добавление нового материала
//
// Параметры:
//
// контекст
// название породы
// латинское название породы

func (s *WoodService) CreateWood(ctx context.Context, name, nameLat string) (*db.Wood, error) {
    // Проверяем, существует ли уже такая порода
    existing, err := s.repo.FindByName(name)
    if err == nil && existing != nil {
        return nil, fmt.Errorf("Материал '%s' уже существует", name)
    }
    
    wood := &db.Wood{
        Name:    name,
        NameLat: nameLat,
    }
    
    if err := s.repo.Create(wood); err != nil {
        return nil, fmt.Errorf("Невозможно создать материал: %w", err)
    }
    
    return wood, nil
}

// Поиск пород дерева по латинскому названию
//
// Параметры:
//
// контекст
// латинское название породы дерева

func (s *WoodService) SearchWoodsByNameLatin(ctx context.Context, latinName string) ([]db.Wood, error) {
    if latinName == "" {
        return s.GetAll(ctx)
    }
    return s.repo.FindByNameLatin(latinName)
}
*/