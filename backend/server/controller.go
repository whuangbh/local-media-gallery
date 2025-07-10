package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MediaFolderInfoRequest struct {
	Url string `json:"url"  binding:"required"`
}

func GetMediaFolderInfo(c *gin.Context) {
	var req MediaFolderInfoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	directory, directories, videos, images, err := QueryFolderInfo(req.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"directory":   directory,
		"directories": directories,
		"videos":      videos,
		"images":      images,
	})
}
