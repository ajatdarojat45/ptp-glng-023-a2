package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	"c2/models"
	"fmt"
	"strconv"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateOrder(c *gin.Context){
	req := models.Order{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"result": "Bad Request",
		})
		return
	}

	errCreate := idb.DB.Create(&req).Error
	if errCreate != nil {
		c.JSON(500, gin.H{
			"result": "internal server error",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"result": "Data successfully created",
	})
}

func (idb *InDB) GetOrders(c *gin.Context){
	var (
		orders []models.Order
		result gin.H
	)

	idb.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count": 0,
		}
	}else {
		result = gin.H{
			"result": orders,
			"count": len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context){
	id := c.Param("orderId")
	orderId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		fmt.Println("error found: ", errConvert)
		c.JSON(400, gin.H{
			"result": "params orderId is required",
		})
		return
	}

	var order models.Order
	errFind := idb.DB.First(&order, orderId).Error
	if errFind != nil {
		c.JSON(400, gin.H{
			"result": "Data not found",
		})
		return
	}

	req := models.Order{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"result": "Bad Request",
		})
		return
	}

	errUpdate := idb.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"result": "internal server error",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"result": "Data successfully updated",
	})
}

func (idb *InDB) DeleteOrder(c *gin.Context){
	var (
		order models.Order
		result gin.H
	)

	id := c.Param("orderId")
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	} else {
		errDelete := idb.DB.Delete(&order).Error
		if errDelete != nil {
			result = gin.H{
				"result": "delete failed",
			}
		}else {
			result = gin.H{
				"result": "Data successfully deleted",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}