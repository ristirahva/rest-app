package repositories

import (
    "gorm.io/gorm"

     "log"

    "github.com/ristirahva/rest-app/db"
)

type CooperRepository struct {
    db *gorm.DB
}

func NewCooperRepository(db *gorm.DB) *CooperRepository {
    return &CooperRepository{
        db: db,
    }
}

// список всех материалов

func (r *CooperRepository) FindAll() ([]db.Cooper, error) {
    var coopers []db.Cooper
    err := r.db.Find(&coopers).Error
    log.Printf("Список материалов: {}", coopers)
    return coopers, err
}


// ???
func (r *CooperRepository) FindByName(name string) (*db.Cooper, error) {
    var cooper db.Cooper
    err := r.db.Where("name = ?", name).First(&cooper).Error
    if err != nil {
        return nil, err
    }
    return &cooper, nil
}


