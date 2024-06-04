package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

	"github.com/google/uuid"
)

func CreateDetailOrder(data models.DetailOrder) (models.DetailOrder, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.DetailOrder{}, tx.Error
	}
	return data, nil
}

func GetDetailOrder(id uuid.UUID) (models.DetailOrder, error) {
	var data models.DetailOrder
	tx := database.DB.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return models.DetailOrder{}, tx.Error
	}
	return data, nil
}

func GetDetailOrderByOrderID(id uuid.UUID) ([]models.DetailOrder, error) {
	var data []models.DetailOrder
	tx := database.DB.Where("order_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.DetailOrder{}, tx.Error
	}
	return data, nil
}

func GetDetailOrderByMenuID(id uuid.UUID) ([]models.DetailOrder, error) {
	var data []models.DetailOrder
	tx := database.DB.Where("menu_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.DetailOrder{}, tx.Error
	}
	return data, nil
}

func UpdateDetailOrder(data models.DetailOrder) (models.DetailOrder, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.DetailOrder{}, tx.Error
	}
	return data, nil
}

func DeleteDetailOrder(data models.DetailOrder) (models.DetailOrder, error) {
	tx := database.DB.Delete(&data)
	if tx.Error != nil {
		return models.DetailOrder{}, tx.Error
	}
	return data, nil
}
