package db

type DrinkInBarrel struct {
    BarrelID     uint      `gorm:"primaryKey;column:barrel_id"`
    DrinkID      uint      `gorm:"primaryKey;column:drink_id"`
    DateStart    time.Time `gorm:"column:date_start"`
    DateEnd      *time.Time `gorm:"column:date_end"` // Может быть NULL
    AlcoholStart *int      `gorm:"column:alcohol_start"`
    AlcoholEnd   *int      `gorm:"column:alcohol_end"`
    Description  string    `gorm:"type:varchar(255)"`
    
    // Связи для удобства доступа
    Barrel Barrel `gorm:"foreignKey:BarrelID"`
    Drink  Drink  `gorm:"foreignKey:DrinkID"`
}

func (DrinkInBarrel) TableName() string {
    return "drink_in_barrel"
}