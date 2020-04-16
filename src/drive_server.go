package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"istorage/config"
	"istorage/controllers"
	"istorage/logger"
)

var envConfiguration = "server.cfg"

func init() {
	godotenv.Load()

	initConfig()
	SentryInit()
}

func main() {
	host := flag.String("host", config.Config.Server.Host+config.Config.Server.Port, "host:port for iMedia server.")
	storage := flag.String("storage", config.Config.Storage.Path, "Root for storage")

	flag.Parse()

	logger.Infof("Storage place in: %s", *storage)
	logger.Infof("Start server on %s", *host)

	config.Config.Storage.Path = *storage

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/:uuid/:profile", controllers.ReadFile, controllers.ReadFileWithResize)
	router.GET("/:uuid", controllers.ReadFile)
	router.DELETE("/:uuid", controllers.DeleteFile)
	router.POST("/upload", controllers.StoreAttachment)

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
		logger.Fatal(err)
	}
}

func SentryInit() {
	dsn, exists := os.LookupEnv("SENTRY_DSN")
	if !exists {
		panic("Not found SENTRY_DSN environment variable")
	}

	raven.SetDSN(dsn)
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
