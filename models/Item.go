package models

import (
  "gorm.io/gorm"
)

type Item struct {
  gorm.Model
	ID        int64     `json:"lineItemId" gorm:"primaryKey"`
  Item_Code string `json:"itemCode" gorm:"column:item_code;type:varchar(25)"`
	Description string `json:"description" gorm:"column:description;type:varchar(255)"`
	Quantity int64 `json:"quantity" gorm:"column:quantity"`
	Order_Id uint `json:"orderId" gorm:"foreignKey:ID"`
}