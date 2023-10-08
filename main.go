package main

import (
	"github.com/adenhidayatuloh/assignment3-011/handler"
	outputlog "github.com/adenhidayatuloh/assignment3-011/output_log"
	"github.com/adenhidayatuloh/assignment3-011/repo/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.HandlerDatabaseConnection()

	go outputlog.OutputLog()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.PUT("waterwind", handler.UpdateWaterWind)

	r.Run(":8080")

}
