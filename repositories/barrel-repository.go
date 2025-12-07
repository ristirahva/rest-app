package repositories

import (
    "time"
    "github.com/ristirahva/rest-app/models"
)

type BarrelRepository struct {
    BaseRepository[models.Barrel]
}

// сщздание репозитория

func NewBarrelRepository(db *gorm.DB) *BarrelRepository {
    return &BarrelRepository{
        BaseRepository: *NewBaseRepository[models.Barrel](db),
    }
}

// поиск бочек по породе дерева
//
// параметр:
//
// woodID - порода дерева (может быть компощитной, например, дуб кавказкий и канадский)

func (r *BarrelRepository) FindByWood(woodID uint) ([]models.Barrel, error) {
    var barrels []models.Barrel
    err := r.db.Where("wood_id = ?", woodID).Preload("Wood").Find(&barrels).Error
    return barrels, err
}

// список всех бочек с залитыми в них напитками

func (r *BarrelRepository) FindNonEmptyBarrels() ([]models.Barrel, error) {
    var barrels []models.Barrel
    now := time.Now()
    
    err := r.db.Preload("Wood").
        Preload("Drinks").
        Joins("JOIN drink_in_barrel dib ON dib.barrel_id = barrels.id").
        Group("barrels.id").
        Find(&barrels).Error
    
    return barrels, err
}

// заливка нового дистиллята в бочку
//
// параметры:
//
// barrellID    - идентификатор бочки
// drinkID      - идентификатор заливаемого дистиллята
// dateStart    - дата заливки
// alcoholStart - крепость дистиллята на дату заливки

// TODO сделать проверку dateStart <= now

func (r *BarrelRepository) AddDrinkToBarrel(barrelID, drinkID uint, dateStart time.Time, alcoholStart *int, description string) error {
    drinkInBarrel := models.DrinkInBarrel{
        BarrelID:     barrelID,
        DrinkID:      drinkID,
        DateStart:    dateStart,
        AlcoholStart: alcoholStart,
        Description:  description,
    }
    return r.db.Create(&drinkInBarrel).Error
}

// окончание выдержки дистиллята
//
// параметры:
//
// barrellID  - идентификатор бочки
// drinkID    - идентификатор выливаемого из бочки дистиллята
// dateEnd    - дата окончания выдержки
// alcoholEnd - итоговая крепость

func (r *BarrelRepository) RemoveDrinkFromBarrel(barrelID, drinkID uint, dateEnd time.Time, alcoholEnd *int) error {
    return r.db.Model(&models.DrinkInBarrel{}).
        Where("barrel_id = ? AND drink_id = ? AND date_end IS NULL", barrelID, drinkID).
        Updates(map[string]interface{}{
            "date_end":    dateEnd,
            "alcohol_end": alcoholEnd,
        }).Error
}
