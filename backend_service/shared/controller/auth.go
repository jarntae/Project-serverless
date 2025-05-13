package controller

import (
	"net/http"
	"time"


	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/cpe-nuntawut/assignment-6-bobox/config"
	"github.com/cpe-nuntawut/assignment-6-bobox/entity"
	"github.com/cpe-nuntawut/assignment-6-bobox/service"
)

type (
	Authen struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func SignIn(c *gin.Context) {
	var payload Authen
	var member entity.Member

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา email ใน members
	memberFound := config.DB().Where("email = ?", payload.Email).First(&member).Error == nil

	// ถ้าไม่พบในทั้งสองตาราง
	if !memberFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบ email นี้ในระบบ"})
		return
	}

	var err error
	if memberFound {
		// ถ้าเป็น member
		err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(payload.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "รหัสผ่านไม่ถูกต้อง (member)"})
			return
		}

		// สร้าง JWT token สำหรับ member
		jwtWrapper := services.JwtWrapper{
			SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
			Issuer:          "AuthService",
			ExpirationHours: 24,
		}

		signedToken, err := jwtWrapper.GenerateToken(member.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
			return
		}

		expirationTime := time.Now().Add(time.Hour * time.Duration(jwtWrapper.ExpirationHours)).Unix()

		c.JSON(http.StatusOK, gin.H{
			"token_type":      "Bearer",
			"token":           signedToken,
			"memberID":        member.ID,
			"token_expiration": expirationTime,
		})
	}
}
