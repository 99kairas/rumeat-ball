package repositories

import (
	"errors"
	"rumeat-ball/database"
	"rumeat-ball/dto"
	"rumeat-ball/models"
	"rumeat-ball/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOrder(data models.Order) (models.Order, error) {
	for {
		// Generate new Order ID
		orderID, err := util.GenerateOrderID("RB")
		if err != nil {
			return models.Order{}, err
		}

		// Cek apakah order ID sudah ada dalam database
		var existingOrder models.Order
		tx := database.DB.Where("id = ?", orderID).First(&existingOrder)
		if tx.Error != nil {
			if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
				// Jika tidak ditemukan, set order ID dan simpan data order
				data.ID = orderID
				tx = database.DB.Save(&data)
				if tx.Error != nil {
					return models.Order{}, tx.Error
				}
				break
			}
			return models.Order{}, tx.Error
		}
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

func GetOrderByOrderID(id string) (models.Order, error) {
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

func GetOrderItemsByOrderID(orderID string, userID uuid.UUID) ([]dto.OrderItem, error) {
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

func DeleteDetailOrdersByOrderID(orderID string) error {
	tx := database.DB.Where("order_id = ?", orderID).Delete(&models.DetailOrder{})
	return tx.Error
}

func UpdateOrderCart(order models.Order) (models.Order, error) {
	tx := database.DB.Save(&order)
	return order, tx.Error
}

func AdminGetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	tx := database.DB.Preload("User").Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func AdminGetOrdersByUserName(userName string) ([]models.Order, error) {
	var orders []models.Order
	tx := database.DB.Preload("User").Joins("JOIN users ON users.id = orders.user_id").
		Where("users.name LIKE ?", "%"+userName+"%").
		Find(&orders)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}

func AdminGetOrderByOrderID(id string) (models.Order, error) {
	var order models.Order
	tx := database.DB.Preload("User").Where("id = ?", id).First(&order)
	if tx.Error != nil {
		return models.Order{}, tx.Error
	}
	return order, nil
}

func AdminGetOrderItemsByOrderID(orderID string) ([]models.DetailOrder, error) {
	var items []models.DetailOrder
	tx := database.DB.Where("order_id = ?", orderID).Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return items, nil
}

func AdminGetUserByID(userID uuid.UUID) (models.User, error) {
	var user models.User
	tx := database.DB.First(&user, "id = ?", userID)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func AdminUpdateOrderStatus(id string, status string) error {
	if err := database.DB.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
