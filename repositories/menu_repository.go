package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"
	"time"

	"github.com/google/uuid"
)

func CreateMenu(menu models.Menu) (models.Menu, error) {
	err := database.DB.Create(&menu).Error
	return menu, err
}

func GetMenu(name string, categoryID uuid.UUID) ([]models.Menu, error) {
	var data []models.Menu
	tx := database.DB

	if name != "" {
		tx = tx.Where("name LIKE ?", "%"+name+"%")
	}
	if categoryID != uuid.Nil {
		tx = tx.Where("category_id = ?", categoryID)
	}

	err := tx.Find(&data).Error
	if err != nil {
		return []models.Menu{}, err
	}
	return data, nil
}

func UpdateMenu(menu models.Menu, id uuid.UUID) (models.Menu, error) {
	err := database.DB.Model(&menu).Where("id = ?", id).Updates(&menu).Error
	return menu, err
}

func DeleteMenu(id uuid.UUID) error {
	err := database.DB.Delete(&models.Menu{}, "id = ?", id).Error
	return err
}

func GetCategoryByID(id uuid.UUID) (models.Category, error) {
	var category models.Category
	err := database.DB.First(&category, "id = ?", id).Error
	return category, err
}

func PermanentlyDeleteOldMenus(olderThan time.Duration) error {
	threshold := time.Now().Add(-olderThan)
	tx := database.DB.Unscoped().Where("deleted_at < ?", threshold).Delete(&models.Menu{})
	return tx.Error
}

func GetMenuByID(id uuid.UUID) (models.Menu, error) {
	var menu models.Menu
	if err := database.DB.Where("id = ?", id).First(&menu).Error; err != nil {
		return models.Menu{}, err
	}
	return menu, nil
}
