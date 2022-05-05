package storage

import (
	"errors"
	"github.com/gaizkadc/wallet-master-6000/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func GetCustomerById(id uuid.UUID) (*models.Customer, error) {
	var customer models.Customer

	err := DB.Model(&customer).Where("id = ?", id).Select()
	if err != nil {
		log.Error().Err(err).Msg("error querying database")
		return nil, err
	}

	return &customer, nil
}

func CustomerExists(id uuid.UUID) bool {
	var customer models.Customer

	err := DB.Model(&customer).Where("id = ?", id).Select()
	if err != nil {
		log.Error().Err(err).Msg("error querying database")
		return false
	}

	if customer.Password != "" {
		return true
	}

	return false
}

func AddBalance(id uuid.UUID, amount decimal.Decimal) error {
	if !CustomerExists(id) {
		return errors.New("customer doesn't exist")
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("invalid amount")
	}

	customer, err := GetCustomerById(id)
	if err != nil {
		return err
	}

	newBalance := customer.Balance.Add(amount)
	customer.Balance = newBalance

	_, err = DB.Model(customer).OnConflict("(id) DO UPDATE").Set("balance = ?", newBalance).Insert()
	if err != nil {
		return err
	}

	return nil
}

func SubtractBalance(id uuid.UUID, amount decimal.Decimal) error {
	if !CustomerExists(id) {
		return errors.New("customer doesn't exist")
	}

	customer, err := GetCustomerById(id)
	if err != nil {
		return err
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("invalid amount")
	}

	if customer.Balance.LessThan(amount) {
		return errors.New("not enough balance")
	}

	newBalance := customer.Balance.Sub(amount)
	customer.Balance = newBalance

	_, err = DB.Model(customer).OnConflict("(id) DO UPDATE").Set("balance = ?", newBalance).Insert()
	if err != nil {
		return err
	}

	return nil
}
