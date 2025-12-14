package db

// деревянная бочка для созревания дистиллятов

type Barrel struct {
    ID          uint   `gorm:"primaryKey;autoIncrement"`
    WoodID      *uint  `gorm:"not null"`// Внешний ключ 
    Wood        *Wood  `gorm:"foreignKey:WoodID"` // связь belongs to
    Volume      int    `gorm:"not null"`
    Description string `gorm:"type:varchar(255)"`

 // Связь many2many с напитками через таблицу drink_in_barrel
    Drinks []Drink `gorm:"many2many:drink_in_barrel;"`
}

// название таблицы

func (Barrel) TableName() string {
    return "barrel"
}

