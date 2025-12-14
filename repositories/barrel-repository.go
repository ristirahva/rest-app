package repositories

import (
    "log"
    "time"
    "gorm.io/gorm"

    "github.com/ristirahva/rest-app/db"
)

type BarrelRepository struct {
    db *gorm.DB
}

func NewBarrelRepository(db *gorm.DB) *BarrelRepository {
    return &BarrelRepository{
        db: db,
    }
}

func (r *BarrelRepository) Create(barrel *db.Barrel) error {
    return r.db.Create(barrel).Error
}


// список всех бочек

func (r *BarrelRepository) FindAll() ([]db.Barrel, error) {
    var barrels []db.Barrel
    err := r.db.Find(&barrels).Error
    log.Printf("Список бочек: {}", barrels)
    return barrels, err
}

// поиск бочек по породе дерева
//
// параметр:
//
// woodID - порода дерева (может быть компощитной, например, дуб кавказкий и канадский)

func (r *BarrelRepository) FindByWood(woodID uint) ([]db.Barrel, error) {
    var barrels []db.Barrel
    err := r.db.Where("wood_id = ?", woodID).Preload("Wood").Find(&barrels).Error
    return barrels, err
}

// список всех бочек с залитыми в них напитками

func (r *BarrelRepository) FindNonEmptyBarrels() ([]db.Barrel, error) {
    var barrels []db.Barrel
    
    err := r.db.Preload("Wood").
        Preload("Drinks").
        Joins("JOIN drink_in_barrel dib ON dib.barrel_id = barrels.id").
        Where("dib.alcohol_end IS NULL").
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
    drinkInBarrel := db.DrinkInBarrel{
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
    return r.db.Model(&db.DrinkInBarrel{}).
        Where("barrel_id = ? AND drink_id = ? AND date_end IS NULL", barrelID, drinkID).
        Updates(map[string]interface{}{
            "date_end":    dateEnd,
            "alcohol_end": alcoholEnd,
        }).Error
}
