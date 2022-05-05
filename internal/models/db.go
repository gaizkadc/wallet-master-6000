package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Customer struct {
	Id       uuid.UUID
	Password string
	Balance  decimal.Decimal
}

type Transaction struct {
	Id           uuid.UUID
	FromCustomer uuid.UUID
	ToCustomer   uuid.UUID
	Amount       decimal.Decimal
	Datetime     time.Time
}
