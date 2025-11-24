package services

import (
	"github.com/gin-gonic/gin"
)

type Chat interface {
	Listen(*gin.Context)
	ReloadSystem(*gin.Context)
}
