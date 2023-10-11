package router

import (
	"fmt"
	"log"

	"poc-consumer-go-endpoint-integrador/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/sinhashubham95/go-actuator"
)

func Roteador() {
	port := utils.GetVariavelAmbiente("app_port", "8080")

	r := gin.New()

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter,
			"/actuator/ping"),
		gin.Recovery(),
	)

	rotaActuator(r)
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}

func rotaActuator(router *gin.Engine) {
	actuatorHandler := actuator.GetActuatorHandler(nil)
	ginActuatorHandler := func(c *gin.Context) {
		actuatorHandler(c.Writer, c.Request)
	}

	router.GET("/actuator/*endpoint", ginActuatorHandler)
}
