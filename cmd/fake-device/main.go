package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/byuoitav/fake-device/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	var (
		port int
	)

	pflag.IntVarP(&port, "port", "P", 8080, "port to run the server on")
	pflag.Parse()

	handlers := handlers.New()

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/all", handlers.All)

	api := r.Group("/v1")

	api.GET("/:address/power", handlers.GetPower)
	api.GET("/:address/power/:on", handlers.SetPower)
	api.GET("/:address/blanked", handlers.GetBlanked)
	api.GET("/:address/blanked/:blanked", handlers.SetBlanked)
	api.GET("/:address/input", handlers.GetInput)
	api.GET("/:address/input/:input", handlers.SetInput)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("unable to bind listener: %s\n", err)
	}

	log.Printf("Starting server on %s\n", lis.Addr().String())

	err = r.RunListener(lis)
	switch {
	case errors.Is(err, http.ErrServerClosed):
	case err != nil:
		log.Fatalf("failed to server: %s\n", err)
	}
}
