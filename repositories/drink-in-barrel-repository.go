package repositories

import (
    "time"
    "gorm.io/gorm"

    "github.com/ristirahva/rest-app/db"
)

type DrinkInBarrelRepository struct {
    db *gorm.DB
}

func NewDrinkInBarrelRepository(db *gorm.DB) *DrinkInBarrelRepository {
    return &DrinkInBarrelRepository{
        db: db,
    }
}


// Список напитков в бочках

func (r *DrinkInBarrelRepository) FindActiveBarrelOccupations() ([]db.DrinkInBarrel, error) {
    var occupations []db.DrinkInBarrel
    now := time.Now()
    err := r.db.Where("date_start IS NOT NULL ? AND date_end IS NULL", now, now).
        Preload("Barrel").
        Preload("Drink").
        Find(&occupations).Error
    return occupations, err
}

// история использования бочки 
//
// параметр: идентификатор бочки
//
// возврат: список выдерживавшихся в бочке напитков, отсортированный по убыванию дат

func (r *DrinkInBarrelRepository) GetBarrelHistory(barrelID uint) ([]db.DrinkInBarrel, error) {
    var history []db.DrinkInBarrel
    err := r.db.Where("barrel_id = ?", barrelID).
        Preload("Drink").
        Order("date_start DESC").
        Find(&history).Error
    return history, err
}

//?????

func (r *DrinkInBarrelRepository) GetDrinkBarrelDuration(drinkID, barrelID uint) (time.Duration, error) {
    var occupation db.DrinkInBarrel
    err := r.db.Where("drink_id = ? AND barrel_id = ?", drinkID, barrelID).
        First(&occupation).Error
    if err != nil {
        return 0, err
    }
    
    var endTime time.Time
    if occupation.DateEnd != nil {
        endTime = *occupation.DateEnd
    } else {
        endTime = time.Now()
    }
    
    return endTime.Sub(occupation.DateStart), nil
}

