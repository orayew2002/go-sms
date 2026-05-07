package handler

const (
	BadRequest           = "wrong_request_body"
	WronApiKey           = "wrong_api_key"
	InvallidPhoneNumber  = "invalid_phone_number"
	SmppErrorConnection  = "smpp_connection_error"
	SmppErrorSendMessage = "message_send_error"
	RedisErrorWrite      = "redis_set_error"
	RedisNotFound        = "message_not_found"
)
