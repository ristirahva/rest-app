package services
/*
import (
    "context"
    "github.com/ristirahva/rest-app/db"
    "github.com/ristirahva/rest-app/repositories"
)

// Базовый сервис

type BaseService[T any] struct {
    repo *repositories.BaseRepository[T]
}

func NewBaseService[T any](repo *repositories.BaseRepository[T]) *BaseService[T] {
    return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) Create(ctx context.Context, entity *T) error {
    return s.repo.Create(entity)
}

func (s *BaseService[T]) GetByID(ctx context.Context, id uint, preloads ...string) (*T, error) {
    return s.repo.FindByID(id, preloads...)
}

func (s *BaseService[T]) Update(ctx context.Context, entity *T) error {
    return s.repo.Update(entity)
}

func (s *BaseService[T]) Delete(ctx context.Context, id uint) error {
    return s.repo.Delete(id)
}

func (s *BaseService[T]) GetAll(ctx context.Context, preloads ...string) ([]T, error) {
    return s.repo.FindAll(preloads...)
}
*/