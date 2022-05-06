package storage_test

import (
	"errors"
	"github.com/gaizkadc/wallet-master-6000/config"
	"github.com/gaizkadc/wallet-master-6000/internal/storage"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"testing"
)

const testConfigPath = "../../config/config.test.yaml"

func TestAddBalance(t *testing.T) {
	err := errors.New("error")
	existingId, _ := uuid.Parse("273e4de9-a5ff-42c2-bdf2-54884c1d19cc")
	validAmount := decimal.NewFromInt(5)
	invalidAmount := decimal.NewFromInt(-5)
	tests := []struct {
		name        string
		id          uuid.UUID
		amount      decimal.Decimal
		expectedErr error
	}{
		{
			name:   "with an existing customer and a valid amount, then it should not return any error",
			id:     existingId,
			amount: validAmount,
		},
		{
			name:        "with an existing customer and an invalid amount, then it should return an error",
			id:          existingId,
			amount:      invalidAmount,
			expectedErr: err,
		},
		{
			name:        "with a non-existing customer, then it should return an error",
			id:          uuid.New(),
			amount:      validAmount,
			expectedErr: err,
		},
	}

	config.Setup(testConfigPath)
	storage.SetupDB()

	for _, tt := range tests {
		tt := tt
		t.Run("Given an AddBalance function, when called "+tt.name, func(t *testing.T) {
			t.Parallel()
			err := storage.AddBalance(tt.id, tt.amount)
			require.IsType(t, tt.expectedErr, err)
		})
	}
}

func TestCustomerExists(t *testing.T) {
	existingId, _ := uuid.Parse("273e4de9-a5ff-42c2-bdf2-54884c1d19cc")
	tests := []struct {
		name           string
		id             uuid.UUID
		expectedResult bool
	}{
		{
			name:           "with an existing customer, then it should return true",
			id:             existingId,
			expectedResult: true,
		},
		{
			name:           "with an existing customer, then it should return false",
			id:             uuid.New(),
			expectedResult: false,
		},
	}

	config.Setup(testConfigPath)
	storage.SetupDB()

	for _, tt := range tests {
		tt := tt
		t.Run("Given an CustomerExists function, when called "+tt.name, func(t *testing.T) {
			t.Parallel()
			result := storage.CustomerExists(tt.id)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestGetCustomerById(t *testing.T) {
	err := errors.New("error")
	existingId, _ := uuid.Parse("273e4de9-a5ff-42c2-bdf2-54884c1d19cc")
	tests := []struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}{
		{
			name: "with an existing customer, then it should return that customer",
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
		t.Run("Given an GetCustomerById function, when called "+tt.name, func(t *testing.T) {
			t.Parallel()
			log.Debug().Interface("id", tt.id).Msg("id")
			retrieved, err := storage.GetCustomerById(tt.id)
			if err != nil {
				require.NotNil(t, tt.expectedErr)
			} else {
				require.Equal(t, tt.id, retrieved.Id)
			}
		})
	}
}

func TestSubtractBalance(t *testing.T) {
	err := errors.New("error")
	existingId, _ := uuid.Parse("273e4de9-a5ff-42c2-bdf2-54884c1d19cc")
	validAmount := decimal.NewFromInt(5)
	invalidAmount := decimal.NewFromInt(-5)
	bigAmount := decimal.NewFromInt(100000)
	tests := []struct {
		name        string
		id          uuid.UUID
		amount      decimal.Decimal
		expectedErr error
	}{
		{
			name:   "with an existing customer and a valid amount, then it should not return any error",
			id:     existingId,
			amount: validAmount,
		},
		{
			name:        "with an existing customer and an ivalid amount, then it should return an error",
			id:          existingId,
			amount:      invalidAmount,
			expectedErr: err,
		},
		{
			name:        "with an existing customer and a bigger amount than the customer's balance, then it should return an error",
			id:          existingId,
			amount:      bigAmount,
			expectedErr: err,
		},
		{
			name:        "with a non-existing customer, then it should return an error",
			id:          uuid.New(),
			amount:      validAmount,
			expectedErr: err,
		},
	}

	config.Setup(testConfigPath)
	storage.SetupDB()

	for _, tt := range tests {
		tt := tt
		t.Run("Given an SubtractBalance function, when called "+tt.name, func(t *testing.T) {
			t.Parallel()
			err := storage.SubtractBalance(tt.id, tt.amount)
			require.IsType(t, tt.expectedErr, err)
		})
	}
}
