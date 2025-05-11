package controller

import (
	"net/http"
	"github.com/cpe-nuntawut/assignment-6-bobox/config"
	"github.com/cpe-nuntawut/assignment-6-bobox/entity"
	"github.com/gin-gonic/gin"
	// "time"
    // "gorm.io/gorm"
)

func GetMemberByID(c *gin.Context) {
    ID := c.Param("id")
	var user entity.Member

	db := config.DB()
	results := db.First(&user, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateMember(c *gin.Context) {
    var Member entity.Member

    // bind เข้าตัวแปร employee
    if err := c.ShouldBindJSON(&Member); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    hashedPassword, _ := config.HashPassword(Member.Password)
	
	// สร้าง Employee
	e := entity.Member{
		FirstName:  Member.FirstName, 
		LastName:   Member.LastName,  
		Email:      Member.Email,  
		Password:   hashedPassword,
	}

    if err := db.Create(&e).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "สมัครใช้งานสำเร็จ"})
}

func UpdateMember(c *gin.Context) {
    var member entity.Member
	memberID := c.Param("id")

	db := config.DB()
	result := db.First(&member, memberID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบสมาชิก"})
		return
	}

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&member)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "แก้ไขข้อมูลไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "แก้ไขข้อมูลสำเร็จ"})
}