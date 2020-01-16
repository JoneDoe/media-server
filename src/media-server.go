package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"istorage/config"
	"istorage/controllers"
)

var envConfiguration = "server.cfg"

func main() {

	initConfig()

	host := flag.String("host", config.Config.Server.Host+config.Config.Server.Port, "host:port for iMedia server.")
	storage := flag.String("storage", config.Config.Storage.Path, "Root for storage")

	flag.Parse()

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/:uuid", controllers.ReadFile)
	router.DELETE("/:uuid", controllers.DeleteFile)
	router.POST("/upload", controllers.StoreAttachment)

	log.Printf("Storage place in: %s", *storage)
	config.Config.Storage.Path = *storage

	log.Printf("Start server on %s", *host)

	server := &http.Server{
		Addr:           config.Config.Server.Port,
		Handler:        router,
		ReadTimeout:    config.Config.Server.ReadTimeout,
		WriteTimeout:   config.Config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}

func initConfig() {
	envName := *flag.String("c", envConfiguration, "Environment config name")

	err := config.LoadConfig(envName)
	if err != nil {
		fmt.Println(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Content-Range, Content-Disposition, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		// c.Next()
	}
}
