package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/dto"
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

func GetOrderByStatus(status string) ([]models.Order, error) {
	var data []models.Order
	tx := database.DB.Where("status = ?", status).Find(&data)
	if tx.Error != nil {
		return []models.Order{}, tx.Error
	}
	return data, nil
}

func GetOrderByOrderID(id uuid.UUID) (models.Order, error) {
	var data models.Order
	tx := database.DB.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return models.Order{}, tx.Error
	}
	return data, nil
}

func GetOrdersByUserID(userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order
	tx := database.DB.Where("user_id = ?", userID).Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func GetOrderItemsByOrderID(orderID, userID uuid.UUID) ([]dto.OrderItem, error) {
	var orderItems []models.DetailOrder
	tx := database.DB.Where("order_id = ?", orderID).Find(&orderItems)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var items []dto.OrderItem
	for _, item := range orderItems {
		menu, err := GetMenuByID(item.MenuID)
		if err != nil {
			return nil, err
		}

		items = append(items, dto.OrderItem{
			MenuID:       item.MenuID,
			UserID:       userID,
			Quantity:     item.Quantity,
			PricePerItem: menu.Price,
			TotalPrice:   item.TotalPrice,
		})
	}
	return items, nil
}

func CancelOrder(orderID uuid.UUID) error {
	// CHANGE STATUS TO CANCELED
	tx := database.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", "cancelled").Delete(&models.Order{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
