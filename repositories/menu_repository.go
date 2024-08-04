package repositories

import (
	"database/sql"
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

func GetCommentCountAndAverageRating(menuID uuid.UUID) (int, float64, error) {
	var commentCount int64
	var averageRating sql.NullFloat64

	// Hitung jumlah komentar
	if err := database.DB.Model(&models.Rating{}).Where("menu_id = ?", menuID).Count(&commentCount).Error; err != nil {
		return 0, 0, err
	}

	// Hitung rata-rata rating
	if err := database.DB.Model(&models.Rating{}).Where("menu_id = ?", menuID).Select("AVG(rating)").Row().Scan(&averageRating); err != nil {
		return 0, 0, err
	}

	// Cek jika averageRating valid
	if !averageRating.Valid {
		return int(commentCount), 0.0, nil
	}

	return int(commentCount), averageRating.Float64, nil
}

func GetAllCommentCountsAndAverageRatings() (map[uuid.UUID]int, map[uuid.UUID]float64, error) {
	var ratings []models.Rating
	commentCounts := make(map[uuid.UUID]int)
	averageRatings := make(map[uuid.UUID]float64)

	if err := database.DB.Find(&ratings).Error; err != nil {
		return nil, nil, err
	}

	for _, rating := range ratings {
		commentCounts[rating.MenuID]++
		averageRatings[rating.MenuID] += rating.Rating
	}

	for id := range averageRatings {
		averageRatings[id] /= float64(commentCounts[id])
	}

	return commentCounts, averageRatings, nil
}

func GetCommentsByMenuID(menuID uuid.UUID) ([]models.Rating, error) {
	var comments []models.Rating
	if err := database.DB.Where("menu_id = ?", menuID).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
