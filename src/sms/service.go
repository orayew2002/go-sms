package sms

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/richxcame/gosms/src/addapter"
	"github.com/richxcame/gosms/src/cache"
	"github.com/richxcame/gosms/src/handler"
	"github.com/richxcame/gosms/src/utils"
)

func Run(ctx context.Context) error {
	c, err := cache.NewCacheClient(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	s, err := addapter.DefaultSmsService()
	if err != nil {
		log.Fatal(err.Error())
	}

	var clients []string
	clientJSON, _ := os.ReadFile(utils.GetEnvD("CLIENTS_FILE_PATH", "clients.json"))
	err = json.Unmarshal(clientJSON, &clients)
	if err != nil {
		log.Fatal(err.Error())
	}

	h := handler.NewHandler(nil, s, c)

	r := gin.Default()
	r.POST("messages", h.Send)
	r.GET("messages/:id", h.Get)

	return r.Run()
}
