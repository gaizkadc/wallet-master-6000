package storage

import (
	"github.com/gaizkadc/wallet-master-6000/internal/models"
	"github.com/google/uuid"
)

func GetTransactionsByCustomerId(customerId uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := DB.Model(&transactions).Where("from_customer = ?", customerId).Select()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func AddTransaction(transaction *models.Transaction) error {
	_, err := DB.Model(transaction).Insert(&transaction)
	if err != nil {
		return err
	}

	return nil
}
