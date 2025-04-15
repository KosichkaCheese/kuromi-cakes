package models

type Cake struct {
	ID    int     `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"` //id
	Image string  `json:"image"`
	Name  string  `json:"name" gorm:"not null"`            //name
	Price float64 `json:"price" gorm:"type:decimal(10,2)"` //price
}

type CakePost struct {
	Image string  `json:"image"`
	Name  string  `json:"name" binding:"required"`         //name
	Price float64 `json:"price" gorm:"type:decimal(10,2)"` //price
}
