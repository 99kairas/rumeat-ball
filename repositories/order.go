package repositories

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