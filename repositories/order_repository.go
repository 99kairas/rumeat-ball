package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

	"github.com/google/uuid"
)

func CreateOrder(data models.Order) (models.Order, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.Order{}, tx.Error
	}
	return data, nil
}

func GetOrder(id uuid.UUID) (models.Order, error) {
	var data models.Order
	tx := database.DB.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return models.Order{}, tx.Error
	}
	return data, nil
}

func GetOrders() ([]models.Order, error) {
	var data []models.Order
	tx := database.DB.Find(&data)
	if tx.Error != nil {
		return []models.Order{}, tx.Error
	}
	return data, nil
}

func UpdateOrder(data models.Order, id uuid.UUID) (models.Order, error) {
	tx := database.DB.Where("id = ?", id).Updates(&data)
	if tx.Error != nil {
		return models.Order{}, tx.Error
	}
	return data, nil
}

func DeleteOrder(id uuid.UUID) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Order{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetOrderByUserID(id uuid.UUID) ([]models.Order, error) {
	var data []models.Order
	tx := database.DB.Where("user_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.Order{}, tx.Error
	}
	return data, nil
}

func GetOrderByMenuID(id uuid.UUID) ([]models.Order, error) {
	var data []models.Order
	tx := database.DB.Where("menu_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.Order{}, tx.Error
	}
	return data, nil
}

func GetOrderByOrderID(id uuid.UUID) ([]models.Order, error) {
	var data []models.Order
	tx := database.DB.Where("order_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.Order{}, tx.Error
	}
	return data, nil
}
