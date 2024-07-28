package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

	"github.com/google/uuid"
)

func CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	tx := database.DB.Create(&transaction)
	if tx.Error != nil {
		return models.Transaction{}, tx.Error
	}
	return transaction, nil
}

func GetTransactionByOrderID(orderID string) (models.Transaction, error) {
	var transaction models.Transaction
	tx := database.DB.Where("order_id = ?", orderID).First(&transaction)
	if tx.Error != nil {
		return models.Transaction{}, tx.Error
	}
	return transaction, nil
}

func UpdateTransactionStatus(orderID string, status string) (models.Transaction, error) {
	var transaction models.Transaction
	tx := database.DB.Where("order_id = ?", orderID).First(&transaction)
	if tx.Error != nil {
		return models.Transaction{}, tx.Error
	}
	transaction.Status = status
	tx = database.DB.Save(&transaction)
	if tx.Error != nil {
		return models.Transaction{}, tx.Error
	}
	return transaction, nil
}

func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	tx := database.DB.Preload("User").Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactions, nil
}

func GetTransactionsByUserID(userID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction
	tx := database.DB.Where("user_id = ?", userID).Preload("User").Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactions, nil
}

func GetTransactionsByUserName(userName string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	tx := database.DB.Joins("User").Where("name LIKE ?", "%"+userName+"%").Preload("User").Find(&transactions)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactions, nil
}

func GetTransactionByID(id uuid.UUID) (models.Transaction, error) {
	var transaction models.Transaction
	tx := database.DB.Where("id = ?", id).Preload("User").First(&transaction)
	if tx.Error != nil {
		return models.Transaction{}, tx.Error
	}
	return transaction, nil
}

func GetTransactionsByOrderID(orderID string) ([]models.Transaction, error) {
	var transaction []models.Transaction
	tx := database.DB.Where("order_id = ?", orderID).Preload("User").First(&transaction)
	if tx.Error != nil {
		return []models.Transaction{}, tx.Error
	}
	return transaction, nil
}
