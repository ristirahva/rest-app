package db

// выдерживаемые в бочках дистилляты 

type Drink struct {
    ID          uint    `gorm:"primaryKey;autoIncrement"`
    Alcohol     *int    // Может быть NULL
    Description string  `gorm:"type:varchar(255)"`
    
    // Связь many2many с бочками
    Barrels []Barrel `gorm:"many2many:drink_in_barrel;"`
    
    // Промежуточная модель для дополнительных полей (опционально)
    DrinkBarrels []DrinkInBarrel `gorm:"foreignKey:DrinkID"`
}

