package handler

import (
	"net/http"

	"github.com/adenhidayatuloh/assignment3-011/service"

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
