package main

import (
	"flag"
	"fmt"
	"github.com/JoneDoe/istorage/config"
	"github.com/gin-gonic/gin"
	"log"

	//"github.com/gin-gonic/contrib/static"

	"github.com/JoneDoe/istorage/controllers"
)

func main() {

	flag.Parse()

	initConfig()

	host := flag.String("host", config.Config.Server.Host+config.Config.Server.Port, "host:port for iMedia server.")
	storage := flag.String("storage", config.Config.Storage.Path, "Root for storage")

	router := gin.Default()
	router.Use(CORSMiddleware())
	//router.Use(static.Serve("/", static.LocalFile(*storage, false)))

	router.POST("/files", controllers.StoreAttachment)

	log.Printf("Storage place in: %s", *storage)
	config.Config.Storage.Path = *storage

	log.Printf("Start server on %s", *host)
	router.Run(*host)
}

func initConfig() {
	envName := *flag.String("c", "server.cfg", "Environment config name")

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
