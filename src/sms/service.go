package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/richxcame/gosms/src/addapter"
	"github.com/richxcame/gosms/src/cache"
	"github.com/richxcame/gosms/src/handler"
	"github.com/richxcame/gosms/src/logger"
	"github.com/richxcame/gosms/src/middleware"
	"github.com/richxcame/gosms/src/utils"
)

func Run(ctx context.Context) error {
	c, err := cache.NewCacheClient(ctx)
	if err != nil {
		return err
	}
	fmt.Println("redis success connected")

	var clients []string
	clientJSON, _ := os.ReadFile(utils.GetEnvD("CLIENTS_FILE_PATH", "clients.json"))
	err = json.Unmarshal(clientJSON, &clients)
	if err != nil {
		return err
	}
	fmt.Println("clients success parsed, clients:", len(clients))

	s, err := addapter.DefaultSmsService()
	if err != nil {
		return err
	}
	fmt.Println("smpp success connected")

	h := handler.NewHandler(clients, s, c)

	log := logger.New()

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.TraceID())
	r.Use(middleware.Logger(log,
		"172.16.208.240",
		"172.16.208.235",
	))

	r.POST("messages", h.Send)
	r.GET("messages/:id", h.Get)

	host := fmt.Sprintf("%s:%s",
		utils.GetEnvD("HOST", "127.0.0.1"),
		utils.GetEnvD("PORT", "8000"))

	fmt.Println("running server", host)
	return r.Run(host)
}
