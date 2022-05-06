package storage_test

import (
	"errors"
	"github.com/gaizkadc/wallet-master-6000/config"
	"github.com/gaizkadc/wallet-master-6000/internal/models"
	"github.com/gaizkadc/wallet-master-6000/internal/storage"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestAddTransaction(t *testing.T) {
	errType := "internal.PGError"
	customerId1, _ := uuid.Parse("273e4de9-a5ff-42c2-bdf2-54884c1d19cc")
	customerId2, _ := uuid.Parse("12e1c07f-5305-476e-b0bb-079dd60bd1dc")
	validTransaction := &models.Transaction{
		Id:           uuid.New(),
		FromCustomer: customerId1,
		ToCustomer:   customerId2,
		Amount:       decimal.NewFromInt(1),
		Datetime:     time.Now(),
	}
	invalidTransaction := &models.Transaction{
		FromCustomer: customerId1,
		ToCustomer:   customerId2,
		Amount:       decimal.NewFromInt(1),
		Datetime:     time.Now(),
	}
	tests := []struct {
		name            string
		expectedErrType string
		transaction     *models.Transaction
	}{
		{
			name:        "with a valid transaction, then it should nt return any error",
			transaction: validTransaction,
		},
		{
			name:            "with an invalid transaction, then it should return an error",
			transaction:     invalidTransaction,
			expectedErrType: errType,
		},
	}

	config.Setup(testConfigPath)
	storage.SetupDB()

	for _, tt := range tests {
		tt := tt
		t.Run("Given a AddTransaction function, when called "+tt.name, func(t *testing.T) {
			t.Parallel()
			err := storage.AddTransaction(tt.transaction)
			if err != nil {
				require.Equal(t, tt.expectedErrType, reflect.TypeOf(err).String())
			}
		})
	}
}

func TestGetTransactionsByCustomerId(t *testing.T) {
	err := errors.New("error")
	existingId, _ := uuid.Parse("273e4de9-a5ff-42c2-bdf2-54884c1d19cc")
	tests := []struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}{
		{
			name: "with an existing customer that has transactions, then it should return those transactions",
			id:   existingId,
		},
		{
			name:        "with a non-existing customer, then it should return an error",
			id:          uuid.New(),
			expectedErr: err,
		},
	}

	config.Setup(testConfigPath)
	storage.SetupDB()

	for _, tt := range tests {
		tt := tt
		t.Run("Given a GetTransactionsByCustomerId function, when called "+tt.name, func(t *testing.T) {
			t.Parallel()
			retrievedTxs, err := storage.GetTransactionsByCustomerId(tt.id)
			require.IsType(t, tt.expectedErr, err)
			if err == nil {
				require.NotZero(t, len(retrievedTxs))
			}
		})
	}
}
