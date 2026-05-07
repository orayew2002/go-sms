package handler

import (
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"

	"github.com/richxcame/gosms/src/addapter"
	"github.com/richxcame/gosms/src/utils"
)

type Handler struct {
	secret_key     []string
	smm            *addapter.SmsService
	cache          *redis.Client
	default_number string
	sms_ttl        time.Duration
}

func NewHandler(keys []string, smm *addapter.SmsService, c *redis.Client) *Handler {
	redisLifeTime, err := strconv.Atoi(os.Getenv("REDIS_LIFE_TIME"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Handler{
		secret_key:     keys,
		smm:            smm,
		cache:          c,
		default_number: utils.GetEnv("SMS_NUMBER"),
		sms_ttl:        time.Second * time.Duration(redisLifeTime),
	}
}

func (h *Handler) Send(c *gin.Context) {
	var body SendMessageRequqest
	if err := c.BindJSON(&body); err != nil {
		ErrorBadRequest(c, BadRequest, err.Error())
		return
	}

	if !slices.Contains(h.secret_key, body.APIKey) {
		ErrorInternalServer(c, WronApiKey, "Please, ensure that your api_key is valid")
		return
	}

	if !utils.IsPhone(body.To) {
		ErrorBadRequest(c, InvallidPhoneNumber, "Phone number must be in format E.164")
		return
	}

	if body.From == "" {
		body.From = h.default_number
	}

	text := utils.GetTextCodec(body.TextType, body.Text)
	respId, err := h.smm.Send(&smpp.ShortMessage{
		Src:      body.From,
		Dst:      body.To,
		Text:     text,
		Register: pdufield.NoDeliveryReceipt,
	})

	if err == smpp.ErrNotConnected {
		ErrorServiceUnavailable(c, SmppErrorConnection, err.Error())
		return
	}

	if err != nil {
		ErrorServiceUnavailable(c, SmppErrorSendMessage, err.Error())
		return
	}

	err = h.cache.SetEx(c, respId, body.Text, h.sms_ttl).Err()
	if err != nil {
		ErrorServiceUnavailable(c, RedisErrorWrite, err.Error())
		return
	}

	ResponseCreated(c, respId, body.Text)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	message, err := h.cache.Get(c, id).Result()
	if err != nil {
		ResponseNotFound(c, RedisNotFound, err.Error())
		return
	}

	ResponseOK(c, id, message)
}
