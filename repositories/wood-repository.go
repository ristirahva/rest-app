package repositories

import (
    "github.com/ristirahva/rest-app/models"
)

type WoodRepository struct {
    BaseRepository[models.Wood]
}

// создание репозитория

func NewWoodRepository(db *gorm.DB) *WoodRepository {
    return &WoodRepository{
        BaseRepository: *NewBaseRepository[models.Wood](db),
    }
}

// ???
func (r *WoodRepository) FindByName(name string) (*models.Wood, error) {
    var wood models.Wood
    err := r.db.Where("name = ?", name).First(&wood).Error
    if err != nil {
        return nil, err
    }
    return &wood, nil
}

// ???
func (r *WoodRepository) FindByNameLatin(nameLat string) ([]models.Wood, error) {
    var woods []models.Wood
    err := r.db.Where("name_lat ILIKE ?", "%"+nameLat+"%").Find(&woods).Error
    return woods, err
}

// ???
func (r *WoodRepository) FindWithBarrels() ([]models.Wood, error) {
    var woods []models.Wood
    err := r.db.Preload("Barrels").Find(&woods).Error
    return woods, err
}

