package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"log/slog"

	slogformatter "github.com/samber/slog-formatter"
	sloggin "github.com/samber/slog-gin"

	"github.com/bzhtux/tsa/internal/tsa"
	"github.com/bzhtux/tsa/internal/utils"
	"github.com/bzhtux/tsa/models"

	_ "go.uber.org/automaxprocs"

	"github.com/gin-gonic/gin"
)

func main() {
	projectDir := utils.GetProjectDir()
	utils.DefaultDBFile = projectDir + "/tsa.db"
	// Create a slog logger, which:
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	logger := slog.New(
		slogformatter.NewFormatterHandler(
			slogformatter.TimezoneConverter(time.UTC),
			slogformatter.TimeFormatter(time.DateTime, nil),
		)(
			slog.NewTextHandler(os.Stdout, nil),
		),
	)

	log.Printf("Tanzu Sample Application is starting ...")
	config := utils.GetConfig()
	log.Printf("Loading SQLite file: %s", config)
	tsa.DBLoader()
	db := utils.ConnectDB(config)
	err := db.AutoMigrate(&models.HttpStatusCode{})
	if err != nil {
		log.Printf("Error Migration Model: %s", err.Error())
	}
	h := tsa.NewBaseHandler(db)

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(sloggin.New(logger))
	router.LoadHTMLGlob(projectDir + "/data/public/templates/*")
	router.Static("/assets", projectDir+"/data/public/assets")
	router.Static("/picture", projectDir+"/data/public/img")
  router.Static("/html", projectDir+"/data/public/html")
	router.MaxMultipartMemory = 16 << 32 // 16 MiB

	router.GET("/healtz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/web")
	})
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", h.GetAllHttpStatusCodes)
		v1.GET("/code/:codename", h.GetOneHttpStatusCode)
	}

	web := router.Group("/web")
	{
		web.GET("/", h.GetIndex)
		web.GET("/code/:codename", h.DisplayCode)
	}

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/html/404.html")
	})

	router.Run("0.0.0.0:8080")
}
