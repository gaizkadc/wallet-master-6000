package controllers

import (
	"github.com/gaizkadc/wallet-master-6000/config"
	"github.com/gaizkadc/wallet-master-6000/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	info := models.AppInfo{
		Name:    config.Config.Server.AppName,
		Version: config.Config.Server.Version,
	}
	c.JSON(http.StatusOK, info)
}
