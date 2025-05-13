package controller

import (
	"net/http"
	"github.com/cpe-nuntawut/assignment-6-bobox/config"
	"github.com/cpe-nuntawut/assignment-6-bobox/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
    // "gorm.io/gorm"
)

func GetLedgerByMemberID(c *gin.Context) {
    MemberIDStr := c.Param("id")
	MemberID, err := strconv.Atoi(MemberIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}
	
	var Ledger entity.Ledger

	db := config.DB()

	results := db.Where("member_id = ?", MemberID).First(&Ledger)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if Ledger.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, Ledger)
}

func CreateLedger(c *gin.Context) {
    var Ledger entity.Ledger

    // bind เข้าตัวแปร member
    if err := c.ShouldBindJSON(&Ledger); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := config.DB()

    // ค้นหา employee ด้วย id
    var Member entity.Member
    db.First(&Member, Ledger.MemberID)
    if Member.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลสมาชิก"})
        return
    }
	
	// ค้นหา gender ด้วย id
	var MoneyType entity.Money_Type
	db.First(&MoneyType, Ledger.Money_Type_ID)
	if MoneyType.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลประเภทการเงิน"})
		return
	}

	var CategoryType entity.Category_Type
	db.First(&CategoryType, Ledger.Category_Type_ID)
	if CategoryType.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลหมวดหมู่"})
		return
	}

    // สร้าง Ledger
    l := entity.Ledger {
        Description:  		Ledger.Description,
		Amount: 			Ledger.Amount,					
	    Date:   			Ledger.Date,												
	    MemberID: 			Ledger.MemberID,				
        Member:   			Member,
		Money_Type_ID:   	Ledger.Money_Type_ID,
		Money_Type:     	MoneyType,
		Category_Type_ID: 	Ledger.Category_Type_ID,
        Category_Type:   	CategoryType,
    }

    if err := db.Create(&l).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "สร้างรายการสำเร็จ"})
}

func UpdateLeder(c *gin.Context) {
    var Ledger entity.Ledger
	LedgerID := c.Param("id")

	db := config.DB()
	result := db.First(&Ledger, LedgerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบรายการข้อมูล"})
		return
	}

	if err := c.ShouldBindJSON(&Ledger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&Ledger)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "แก้ไขรายการข้อมูลไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "แก้ไขรายการข้อมูลสำเร็จ"})
}

func DeleteLedger(c *gin.Context) {
    id := c.Param("id")

	db := config.DB()

    var Ledger entity.Ledger
    if err := db.First(&Ledger, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบรายการข้อมูล"})
        return
    }

    tx := db.Begin()
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
        return
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed: "})
        }
    }()

    // Soft delete ledger
    if err := tx.Model(&Ledger).Update("deleted_at", time.Now()).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "ลบรายการข้อมูลไม่สำเร็จ"})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "ลบรายการข้อมูลสำเร็จ"})
}