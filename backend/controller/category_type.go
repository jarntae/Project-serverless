package controller

import (
	"net/http"
	"github.com/cpe-nuntawut/assignment-6-bobox/config"
	"github.com/cpe-nuntawut/assignment-6-bobox/entity"
	"github.com/gin-gonic/gin"
	// "time"
    // "gorm.io/gorm"
)

func GetCategoryType(c *gin.Context) {
	var Category []entity.Category_Type

	db := config.DB()
	results := db.Find(&Category)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, Category)
}