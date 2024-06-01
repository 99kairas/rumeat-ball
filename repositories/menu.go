package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

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
	tx := database.DB.Delete(&models.Menu{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateMenu(data models.Menu) (models.Menu, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.Menu{}, tx.Error
	}
	return data, nil
}
