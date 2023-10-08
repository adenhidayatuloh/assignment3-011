package main

import (
	"assigment3-GLNG-08-011/handler"
	outputlog "assigment3-GLNG-08-011/output_log"
	"assigment3-GLNG-08-011/repo/database"

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
