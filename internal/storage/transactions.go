package storage

import (
	"github.com/gaizkadc/wallet-master-6000/internal/models"
	"github.com/google/uuid"
)

func GetTransactionsByCustomerId(customerId uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction

	_, err := DB.Query(&transactions, "select * from transactions where from_customer = ?", customerId)
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
