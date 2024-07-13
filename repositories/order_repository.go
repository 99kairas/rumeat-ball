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

func GetUserProfile(userID uuid.UUID) (models.User, error) {
	var user models.User
	tx := database.DB.Where("id = ?", userID).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func UpdateUserProfile(userID uuid.UUID, updatedData models.User) (models.User, error) {
	var user models.User
	// Ambil data pengguna yang ada
	tx := database.DB.Where("id = ?", userID).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	// Perbarui field hanya jika ada nilai baru yang diberikan
	if updatedData.Name != "" {
		user.Name = updatedData.Name
	}
	if updatedData.Password != "" {
		user.Password = updatedData.Password
	}
	if updatedData.Address != "" {
		user.Address = updatedData.Address
	}
	if updatedData.Phone != "" {
		user.Phone = updatedData.Phone
	}
	if updatedData.ProfileImage != "" {
		user.ProfileImage = updatedData.ProfileImage
	}

	// Simpan perubahan
	tx = database.DB.Save(&user)
	return user, tx.Error
}
