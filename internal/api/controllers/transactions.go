package controllers

import (
	"errors"
	"github.com/gaizkadc/wallet-master-6000/internal/errxrs"
	"github.com/gaizkadc/wallet-master-6000/internal/models"
	"github.com/gaizkadc/wallet-master-6000/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func GetTransactionsByCustomerId(c *gin.Context) {
	customerIdStr := c.Param("id")

	customerId, err := uuid.Parse(customerIdStr)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
	}

	if !isAuthorized(*c.Request, customerId) {
		errxrs.NewError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	transactions, err := storage.GetTransactionsByCustomerId(customerId)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func AddTransaction(c *gin.Context) {
	customerId, ok := isAuthenticated(*c.Request)
	if !ok {
		errxrs.NewError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	input := &models.Transaction{}
	err := c.BindJSON(input)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		log.Error().Err(err).Interface("transaction", input).Msg("error marshaling json")

		return
	}

	input.Id = uuid.New()
	input.FromCustomer = *customerId
	input.Datetime = time.Now()

	if !storage.CustomerExists(input.ToCustomer) {
		errxrs.NewError(c, http.StatusBadRequest, errors.New("destination customer doesn't exist"))

		return
	}

	err = storage.SubstractBalance(input.FromCustomer, input.Amount)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = storage.AddBalance(input.ToCustomer, input.Amount)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		return
	}

	err = storage.AddTransaction(input)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		log.Error().Err(err).Msg("database error")

		return
	}

	c.JSON(http.StatusOK, input)
}
