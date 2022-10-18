package controller

import (
	"net/http"

	"github.com/Pet002/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

// POST /MedicineUse
func CreateMedicineUse(c *gin.Context) {
	var MedicineUse entity.MedicineUse
	if err := c.ShouldBindJSON(&MedicineUse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&MedicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": MedicineUse})
}

// GET /MedicineUse/:id
func GetMedicineUse(c *gin.Context) {
	var MedicineLabel entity.MedicineLabel
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&MedicineLabel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineLabel})
}

// GET /MedicineUse
func ListMedicineUse(c *gin.Context) {
	var MedicineUse []entity.MedicineUse
	if err := entity.DB().Raw("SELECT * FROM medicine_uses").Scan(&MedicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineUse})
}

// DELETE /MedicineUse/:id
func DeleteMedicineUse(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine_use WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /MedicineUse
func UpdateMedicineUse(c *gin.Context) {
	var MedicineUse entity.MedicineUse
	if err := c.ShouldBindJSON(&MedicineUse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", MedicineUse.ID).First(&MedicineUse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	if err := entity.DB().Save(&MedicineUse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineUse})
}
