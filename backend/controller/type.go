package controller

import (
			"github.com/Pet002/Project-sa-65/entity"
			"github.com/gin-gonic/gin"
			"net/http"
)

// POST /type
func CreateType(c *gin.Context) {
	var MedicineType entity.Type
	if err := c.ShouldBindJSON(&MedicineType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

// GET /type/:id
func GetType(c *gin.Context) {
	var MedicineType entity.Type
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM types WHERE id = ?", id).Scan(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

// GET /type
func ListType(c *gin.Context) {
	var MedicineType []entity.Type
	if err := entity.DB().Raw("SELECT * FROM types").Scan(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}

// DELETE /type/:id
func DeleteType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicineType not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /type
func UpdateType(c *gin.Context) {
	var MedicineType entity.Type
	if err := c.ShouldBindJSON(&MedicineType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", MedicineType.ID).First(&MedicineType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})
		return
	}

	if err := entity.DB().Save(&MedicineType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineType})
}