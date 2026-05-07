package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func R(c *gin.Context, statusCode int, id, message string) {
	c.JSON(statusCode, Response{Id: id, Message: message})
}

func ErrorBadRequest(c *gin.Context, id, message string) {
	R(c, http.StatusBadRequest, id, message)
}

func ErrorInternalServer(c *gin.Context, id, message string) {
	R(c, http.StatusInternalServerError, id, message)
}

func ErrorServiceUnavailable(c *gin.Context, id, message string) {
	R(c, http.StatusServiceUnavailable, id, message)
}

func ResponseCreated(c *gin.Context, id, message string) {
	R(c, http.StatusCreated, id, message)
}

func ResponseOK(c *gin.Context, id, message string) {
	R(c, http.StatusOK, id, message)
}

func ResponseNotFound(c *gin.Context, id, message string) {
	R(c, http.StatusNotFound, id, message)
}
