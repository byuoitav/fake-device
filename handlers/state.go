package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetPower(c *gin.Context) {
	dev := h.device(c.Param("address"))

	pow := "standby"
	if dev.On {
		pow = "on"
	}

	c.JSON(http.StatusOK, gin.H{
		"power": pow,
	})
}

func (h *Handlers) SetPower(c *gin.Context) {
	addr := c.Param("address")
	dev := h.device(addr)
	on, err := strconv.ParseBool(c.Param("on"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dev.On = on
	h.save(addr, dev)

	h.GetPower(c)
}

func (h *Handlers) GetBlanked(c *gin.Context) {
	dev := h.device(c.Param("address"))

	c.JSON(http.StatusOK, gin.H{
		"blanked": dev.Blanked,
	})
}

func (h *Handlers) SetBlanked(c *gin.Context) {
	addr := c.Param("address")
	dev := h.device(addr)
	blanked, err := strconv.ParseBool(c.Param("blanked"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	dev.Blanked = blanked
	h.save(addr, dev)

	h.GetBlanked(c)
}

func (h *Handlers) GetInput(c *gin.Context) {
	dev := h.device(c.Param("address"))

	c.JSON(http.StatusOK, gin.H{
		"input": dev.Input,
	})
}

func (h *Handlers) SetInput(c *gin.Context) {
	addr := c.Param("address")
	dev := h.device(addr)
	dev.Input = c.Param("input")

	h.save(addr, dev)

	h.GetInput(c)
}
