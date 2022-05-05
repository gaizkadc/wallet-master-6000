package controllers

import (
	"errors"
	"github.com/gaizkadc/wallet-master-6000/internal/errxrs"
	"github.com/gaizkadc/wallet-master-6000/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

func GetCustomerById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		return
	}

	if !isAuthorized(*c.Request, id) {
		errxrs.NewError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	customer, err := storage.GetCustomerById(id)
	if err != nil {
		errxrs.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func isAuthorized(req http.Request, customerId uuid.UUID) bool {
	authedUserId, ok := isAuthenticated(req)
	if !ok {
		return false
	}

	return *authedUserId == customerId
}

func isAuthenticated(req http.Request) (*uuid.UUID, bool) {
	user, pass, ok := req.BasicAuth()

	userId, err := uuid.Parse(user)
	if err != nil {
		log.Error().Err(err).Msg("error parsing user id")

		return nil, false
	}

	if !storage.CustomerExists(userId) {
		return nil, false
	}

	customer, err := storage.GetCustomerById(userId)
	if err != nil {
		log.Error().Err(err).Msg("error retrieving customer")

		return nil, false
	}

	if ok && pass == customer.Password {
		return &userId, true
	}

	log.Info().Str("user id", userId.String()).Msg("user not authenticated")

	return nil, false
}
