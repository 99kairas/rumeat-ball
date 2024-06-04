package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"
	"time"

	"github.com/google/uuid"
)

func CreateMenu(data models.Menu) (models.Menu, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.Menu{}, tx.Error
	}
	return data, nil
}

func GetMenu() ([]models.Menu, error) {
	var data []models.Menu
	tx := database.DB.Find(&data)
	if tx.Error != nil {
		return []models.Menu{}, tx.Error
	}
	return data, nil
}

func GetMenuByID(id uuid.UUID) (models.Menu, error) {
	var data models.Menu
	tx := database.DB.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return models.Menu{}, tx.Error
	}
	return data, nil
}

func DeleteMenu(id uuid.UUID) error {
	tx := database.DB.Delete(&models.Menu{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func PermanentlyDeleteOldMenus(olderThan time.Duration) error {
	threshold := time.Now().Add(-olderThan)
	tx := database.DB.Unscoped().Where("deleted_at < ?", threshold).Delete(&models.Menu{})
	return tx.Error
}

func UpdateMenu(data models.Menu, id uuid.UUID) (models.Menu, error) {
	tx := database.DB.Where("id = ?", id).Updates(&data)
	if tx.Error != nil {
		return models.Menu{}, tx.Error
	}
	return data, nil
}
