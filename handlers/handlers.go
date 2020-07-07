package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	devicesMu sync.Mutex
	devices   map[string]device
}

type device struct {
	On      bool   `json:"on"`
	Blanked bool   `json:"blanked"`
	Input   string `json:"input"`
}

func New() *Handlers {
	return &Handlers{
		devices: make(map[string]device),
	}
}

func (h *Handlers) device(addr string) device {
	h.devicesMu.Lock()
	defer h.devicesMu.Unlock()

	if _, ok := h.devices[addr]; !ok {
		h.devices[addr] = device{
			On:    true,
			Input: "input",
		}
	}

	return h.devices[addr]
}

func (h *Handlers) save(addr string, dev device) {
	h.devicesMu.Lock()
	defer h.devicesMu.Unlock()

	h.devices[addr] = dev
}

func (h *Handlers) All(c *gin.Context) {
	h.devicesMu.Lock()
	defer h.devicesMu.Unlock()

	c.JSON(http.StatusOK, h.devices)
}
