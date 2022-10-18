package controller

import (
			"github.com/Pet002/Project-sa-65/entity"
			"github.com/gin-gonic/gin"
			"net/http"
)

// POST /storage
func CreateStorage(c *gin.Context) {
	var storage entity.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// GET /storage/:id
func GetStorage(c *gin.Context) {
	var storage entity.Storage
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM storages WHERE id = ?", id).Scan(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// GET /storage
func ListStorage(c *gin.Context) {
	var storage []entity.Storage
	if err := entity.DB().Raw("SELECT * FROM storages").Scan(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storage})
}

// DELETE /storage/:id
func DeleteStorage(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM storages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storages not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /storage
func UpdateStorage(c *gin.Context) {
	var storage entity.Storage
	if err := c.ShouldBindJSON(&storage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", storage.ID).First(&storage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storages not found"})
		return
	}

	if err := entity.DB().Save(&storage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": storage})
}