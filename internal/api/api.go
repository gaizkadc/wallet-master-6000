package api

import (
	"fmt"
	"github.com/gaizkadc/wallet-master-6000/config"
	"github.com/gaizkadc/wallet-master-6000/internal/api/router"
	"github.com/gaizkadc/wallet-master-6000/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	storage.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "config/config.dev.yaml"
	}

	setConfiguration(configPath)
	conf := config.GetConfig()

	web := router.Setup()

	log.Info().Msg(fmt.Sprintf("%s started in port %s", conf.Server.AppName, conf.Server.Port))

	if err := web.Run(":" + conf.Server.Port); err != nil {
		log.Fatal().Err(err).Msg("error running server")
	}
}
