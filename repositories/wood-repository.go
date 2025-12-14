package repositories

import (
    "gorm.io/gorm"

     "log"

    "github.com/ristirahva/rest-app/db"
)

type WoodRepository struct {
    db *gorm.DB
}

func NewWoodRepository(db *gorm.DB) *WoodRepository {
    return &WoodRepository{
        db: db,
    }
}

// список всех материалов

func (r *WoodRepository) FindAll() ([]db.Wood, error) {
    var woods []db.Wood
    err := r.db.Find(&woods).Error
    log.Printf("Список материалов: {}", woods)
    return woods, err
}


// ???
func (r *WoodRepository) FindByName(name string) (*db.Wood, error) {
    var wood db.Wood
    err := r.db.Where("name = ?", name).First(&wood).Error
    if err != nil {
        return nil, err
    }
    return &wood, nil
}

// ???
func (r *WoodRepository) FindByNameLatin(nameLat string) ([]db.Wood, error) {
    var woods []db.Wood
    err := r.db.Where("name_lat ILIKE ?", "%"+nameLat+"%").Find(&woods).Error
    return woods, err
}

// ???
func (r *WoodRepository) FindWithBarrels() ([]db.Wood, error) {
    var woods []db.Wood
    err := r.db.Preload("Barrels").Find(&woods).Error
    return woods, err
}

