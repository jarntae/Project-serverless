package controller

import (
	"net/http"
	"github.com/cpe-nuntawut/assignment-6-bobox/config"
	"github.com/cpe-nuntawut/assignment-6-bobox/entity"
	"github.com/gin-gonic/gin"
	// "time"
    // "gorm.io/gorm"
 )

 func GetMoneyType(c *gin.Context) {
	var MoneyType []entity.Money_Type

	db := config.DB()
	results := db.Find(&MoneyType)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, MoneyType)
}