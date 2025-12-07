package repositories

import (
    "github.com/ristirahva/rest-app/models"
)

type DrinkRepository struct {
    BaseRepository[models.Drink]
}

func NewDrinkRepository(db *gorm.DB) *DrinkRepository {
    return &DrinkRepository{
        BaseRepository: *NewBaseRepository[models.Drink](db),
    }
}

// список выдерживаемых дистиллятов
//
// возврат: список дистиллятов, выдерживающихся в настоящее время в бочках

func (r *DrinkRepository) FindCurrentlyInBarrels() ([]models.Drink, error) {
    var drinks []models.Drink
    err := r.db.Joins("JOIN drink_in_barrel dib ON dib.drink_id = drinks.id").
        Where("dib.date_end IS NULL").
        Group("drinks.id").
        Preload("Barrels").
        Find(&drinks).Error
    return drinks, err
}

// история выдержки однотипного дистиллята
//
// параметр: идентификатор дистиллята
//
// возврат: список дистиллятов по бочкам, отсортированный по убыванию дат заливки

func (r *DrinkRepository) GetDrinkHistory(drinkID uint) ([]models.DrinkInBarrel, error) {
    var history []models.DrinkInBarrel
    err := r.db.Where("drink_id = ?", drinkID).
        Preload("Barrel").
        Preload("Barrel.Wood").
        Order("date_start DESC").
        Find(&history).Error
    return history, err
}