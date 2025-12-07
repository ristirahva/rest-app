package repositories

import (
    "gorm.io/gorm"
)

// BaseRepository содержит общие методы для всех репозиториев
type BaseRepository[T any] struct {
    db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
    return &BaseRepository[T]{db: db}
}

// Общие CRUD методы
func (r *BaseRepository[T]) Create(entity *T) error {
    return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) FindByID(id uint, preloads ...string) (*T, error) {
    var entity T
    query := r.db
    for _, preload := range preloads {
        query = query.Preload(preload)
    }
    err := query.First(&entity, id).Error
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func (r *BaseRepository[T]) Update(entity *T) error {
    return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
    var entity T
    return r.db.Delete(&entity, id).Error
}

func (r *BaseRepository[T]) FindAll(preloads ...string) ([]T, error) {
    var entities []T
    query := r.db
    for _, preload := range preloads {
        query = query.Preload(preload)
    }
    err := query.Find(&entities).Error
    return entities, err
}

func (r *BaseRepository[T]) FindByCondition(condition interface{}, preloads ...string) ([]T, error) {
    var entities []T
    query := r.db.Where(condition)
    for _, preload := range preloads {
        query = query.Preload(preload)
    }
    err := query.Find(&entities).Error
    return entities, err
}