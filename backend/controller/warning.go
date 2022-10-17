package controller

import (
	"net/http"

	"github.com/Thanaporn4226/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

// POST /Warning
func CreateWarning(c *gin.Context) {
	var Warning entity.Warning
	if err := c.ShouldBindJSON(&Warning); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&Warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": Warning})
}

// GET /Warning/:id
func GetWarning(c *gin.Context) {
	var Warning entity.Warning
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&Warning); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Warning})
}

// GET /Warning
func ListWarning(c *gin.Context) {
	var Warning []entity.Warning
	if err := entity.DB().Raw("SELECT * FROM warnings").Scan(&Warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Warning})
}

// DELETE /Warning/:id
func DeleteWarning(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM warning WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Warning
func UpdateWarning(c *gin.Context) {
	var Warning entity.Warning
	if err := c.ShouldBindJSON(&Warning); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Warning.ID).First(&Warning); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	if err := entity.DB().Save(&Warning).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Warning})
}
