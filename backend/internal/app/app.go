package app

import (
	"log"
	"lottery/internal/api/router"
	"lottery/internal/config"
	"lottery/internal/service"
	"lottery/pkg/database"
	"net/http"
	"strconv"
	"strings"
)

func Run() {
	config.InitEnvConfigs()
	db := database.NewDB()

	winnerPool := make(map[uint]struct{})
	for _, v := range strings.Split(config.EnvConfig.WinnerPool, " ") {
		number, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Fatal("Failed parse winner pool")
		}
		winnerPool[uint(number)] = struct{}{}
	}

	service := service.NewService(db, winnerPool, config.EnvConfig.MinPoolSize)
	server := router.NewServer(service)
	server.SetupRoutes()
	log.Println("Server started", config.EnvConfig.LocalServerPort)
	err := http.ListenAndServe(":"+config.EnvConfig.LocalServerPort, nil)
	log.Println(err)
}
