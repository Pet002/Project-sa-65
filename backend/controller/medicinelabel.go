package controller

import (
	"net/http"

	"github.com/Thanaporn4226/Project-sa-65/entity"
	"github.com/gin-gonic/gin"
)

// POST /MedicineLabel
func CreateMedicineLabel(c *gin.Context) {

	var medicinelabel entity.MedicineLabel
	var medicineuse entity.MedicineUse
	var warning entity.Warning
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&medicinelabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา medicineuse ด้วย id
	if tx := entity.DB().Where("id = ?", medicinelabel.MedicineUseID).First(&medicineuse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_use not found"})
		return
	}

	// 10: ค้นหา warning ด้วย id
	if tx := entity.DB().Where("id = ?", medicinelabel.WarningID).First(&warning); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "warning not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", medicinelabel.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 12: สร้าง MedicineLabel
	wv := entity.MedicineLabel{
		MedicineUse:   medicineuse,                 // โยงความสัมพันธ์กับ Entity medicineuse
		Warning:       warning,                     // โยงความสัมพันธ์กับ Entity warning
		Employee:      employee,                    // โยงความสัมพันธ์กับ Entity employee
		RecordingDate: medicinelabel.RecordingDate, // ตั้งค่าฟิลด์ RecordingDate
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /MedicineLabel/:id
func GetMedicineLabel(c *gin.Context) {
	var MedicineLabel entity.MedicineLabel
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&MedicineLabel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_label not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineLabel})
}

// GET /MedicineLabel
func ListMedicineLabel(c *gin.Context) {
	var MedicineLabel []entity.MedicineLabel
	if err := entity.DB().Preload("MedicineUse").Preload("Warning").Preload("Employee").Raw("SELECT * FROM medicine_labels").Find(&MedicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineLabel})
}

// DELETE /MedicineLabel/:id
func DeleteMedicineLabel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicine_label WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_label not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /MedicineLabel
func UpdateMedicineLabel(c *gin.Context) {
	var MedicineLabel entity.MedicineLabel
	if err := c.ShouldBindJSON(&MedicineLabel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", MedicineLabel.ID).First(&MedicineLabel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine_label not found"})
		return
	}

	if err := entity.DB().Save(&MedicineLabel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": MedicineLabel})
}
