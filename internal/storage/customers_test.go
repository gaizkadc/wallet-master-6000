package storage_test

import (
	"errors"
	"github.com/gaizkadc/wallet-master-6000/config"
	"github.com/gaizkadc/wallet-master-6000/internal/storage"
	"github.com/google/uuid"
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
			name:        "with an existing customer and an ivalid amount, then it should return an error",
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

//func TestCustomerExists(t *testing.T) {
//	type args struct {
//		id uuid.UUID
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := storage.CustomerExists(tt.args.id); got != tt.want {
//				t.Errorf("CustomerExists() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGetCustomerById(t *testing.T) {
//	type args struct {
//		id uuid.UUID
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *models.Customer
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := storage.GetCustomerById(tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetCustomerById() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetCustomerById() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSubstractBalance(t *testing.T) {
//	type args struct {
//		id     uuid.UUID
//		amount decimal.Decimal
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := storage.SubstractBalance(tt.args.id, tt.args.amount); (err != nil) != tt.wantErr {
//				t.Errorf("SubstractBalance() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
