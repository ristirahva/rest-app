package db

//import "gorm.io/gorm"

// сорт дерева, из которого изготовлена бочка; 
// может быть композитным (например, скальный/канадский дуб)

type Cooper struct {
    ID      uint   `gorm:"primaryKey;autoIncrement"` // serial в PostgreSQL
    Name    string `gorm:"type:varchar(64);unique;not null"`
    URL     string `gorm:"type:varchar(64);unique;not null"`
}

// TableName задает имя таблицы (опционально)
func (Cooper) TableName() string {
    return "cooper"
}