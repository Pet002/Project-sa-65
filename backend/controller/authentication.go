package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/Pet002/Project-sa-65/services"
	"github.com/gin-gonic/gin"
)

type LoginPayload struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string
	UserID   uint   `json:"user_id"`
	RoleName string `json:"role_name"`
}

// POST /signin
func Signin(c *gin.Context) {
	var payload LoginPayload
	var login entity.Login
	var role entity.Role
	var employee entity.Employee

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา login ด้วย Username ที่ผู้ใช้กรอกมา
	if err := entity.DB().Raw("SELECT * FROM logins WHERE user = ?", payload.User).Scan(&login).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ตรวจสอบ Password
	err := services.VerifyPassword(login.Password, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user credentials"})
		return
	}

	//ค้นหา Employee Role ID ด้วย login_id
	if err := entity.DB().Raw("SELECT * FROM employees WHERE login_id = ?", login.ID).Scan(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา Role ด้วย role_id
	if err := entity.DB().Raw("SELECT * FROM roles WHERE id = ?", employee.RoleID).Scan(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtWrapper := services.JwtWrapper{
		SecretKey:      "Secret",
		Issuer:         "AuthService",
		ExpirationHour: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(login.ID, role.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Token:    signedToken,
		UserID:   login.ID,
		RoleName: role.Name,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}
