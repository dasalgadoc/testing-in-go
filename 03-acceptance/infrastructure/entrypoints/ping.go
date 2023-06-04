package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const pong = "pong"

type Ping struct{}

func NewPing() Ping {
	return Ping{}
}

func (p *Ping) Pong(c *gin.Context) {
	c.String(http.StatusOK, pong)
}
