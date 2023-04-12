package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/helper/SnowFlake"
	"net/http"
)

type CommonController struct{}

func (cc *CommonController) CreateSnowflakeId(c *gin.Context, worker *SnowFlake.Worker) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    http.StatusOK,
		"info": map[string]interface{}{
			"id": worker.GetId(),
		},
	})
}
