package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

	"github.com/google/uuid"
)

func CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
    err := database.DB.Create(&transaction).Error
    return transaction, err
}

func GetTransactionByOrderID(orderID uuid.UUID) (models.Transaction, error) {
    var transaction models.Transaction
    err := database.DB.Where("order_id = ?", orderID).First(&transaction).Error
    return transaction, err
}

func GetTransactionsByUserID(userID uuid.UUID) ([]models.Transaction, error) {
    var transactions []models.Transaction
    err := database.DB.Where("user_id = ?", userID).Find(&transactions).Error
    return transactions, err
}

func UpdateTransactionStatus(transactionID uuid.UUID, status string) error {
    return database.DB.Model(&models.Transaction{}).Where("id = ?", transactionID).Update("status", status).Error
}


func GetTransactionByUserIDAndOrderID(userID, orderID uuid.UUID) (models.Transaction, error) {
	var transaction models.Transaction

	err := database.DB.Where("user_id = ? AND order_id = ?", userID, orderID).First(&transaction).Error
	return transaction, err
}

func GetTransactionByUserIDAndStatus(userID uuid.UUID, status string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := database.DB.Where("user_id = ? AND status = ?", userID, status).Find(&transactions).Error
	return transactions, err
}

func GetTransactionByOrderIDAndStatus(orderID uuid.UUID, status string) (models.Transaction, error) {
	var transaction models.Transaction

	err := database.DB.Where("order_id = ? AND status = ?", orderID, status).First(&transaction).Error
	return transaction, err
}

func UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := database.DB.Save(&transaction).Error
	return transaction, err
}

func DeleteTransaction(transaction models.Transaction) error {
	err := database.DB.Delete(&transaction).Error
	return err
}

func DeleteTransactionByOrderID(orderID uuid.UUID) error {
	err := database.DB.Where("order_id = ?", orderID).Delete(&models.Transaction{}).Error
	return err
}

func DeleteTransactionByUserID(userID uuid.UUID) error {
	err := database.DB.Where("user_id = ?", userID).Delete(&models.Transaction{}).Error
	return err
}

func DeleteTransactionByUserIDAndOrderID(userID, orderID uuid.UUID) error {
	err := database.DB.Where("user_id = ? AND order_id = ?", userID, orderID).Delete(&models.Transaction{}).Error
	return err
}

func DeleteTransactionByUserIDAndStatus(userID uuid.UUID, status string) error {
	err := database.DB.Where("user_id = ? AND status = ?", userID, status).Delete(&models.Transaction{}).Error
	return err
}

func DeleteTransactionByOrderIDAndStatus(orderID uuid.UUID, status string) error {
	err := database.DB.Where("order_id = ? AND status = ?", orderID, status).Delete(&models.Transaction{}).Error
	return err
}

func DeleteTransactionByUserIDAndOrderIDAndStatus(userID, orderID uuid.UUID, status string) error {
	err := database.DB.Where("user_id = ? AND order_id = ? AND status = ?", userID, orderID, status).Delete(&models.Transaction{}).Error
	return err
}

func GetTransactionByUserIDAndMenuID(userID, menuID uuid.UUID) (models.Transaction, error) {
	var transaction models.Transaction

	err := database.DB.Where("user_id = ? AND menu_id = ?", userID, menuID).First(&transaction).Error
	return transaction, err
}
