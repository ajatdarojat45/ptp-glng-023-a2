package models

import (
  "gorm.io/gorm"
	"time"
)

type Order struct {
  gorm.Model
	ID        int64     `json:"id" gorm:"primaryKey"`
	Customer_Name string `json:"customerName" gorm:"column:customer_name;type:varchar(25)"`
	Order_At time.Time `json:"updated_at" gorm:"column:updated_at"`
	Items []Item `json:"items" gorm:"foreignKey:Order_Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}