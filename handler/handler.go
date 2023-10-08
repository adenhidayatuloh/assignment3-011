package handler

import (
	"assigment3-GLNG-08-011/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateWaterWind(ctx *gin.Context) {

	response, err := service.UpdatePeriod()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)

}
