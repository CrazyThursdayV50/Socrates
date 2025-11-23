package gemini

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Chatter) Listen(ctx *gin.Context) {
	c.logger.Debugf("ws incoming new conn")
	err := c.wsserver.Run(ctx, ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, struct{}{})
}
