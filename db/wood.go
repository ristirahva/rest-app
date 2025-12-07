package db

import "gorm.io/gorm"

// сорт дерева, из которого изготовлена бочка; 
// может быть композитным (например, скальный/канадский дуб)

type Wood struct {
    ID      uint   `gorm:"primaryKey;autoIncrement"` // serial в PostgreSQL
    Name    string `gorm:"type:varchar(64);unique;not null"`
    NameLat string `gorm:"type:varchar(64);column:name_lat"`
    Barrels []Barrel `gorm:"foreignKey:WoodID"` // связь has many
}

// TableName задает имя таблицы (опционально)
func (Wood) TableName() string {
    return "wood"
}