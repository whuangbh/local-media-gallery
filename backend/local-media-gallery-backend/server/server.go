package server

import (
	"github.com/gin-gonic/gin"
)

func Init() error {
	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"127.0.0.0"}); err != nil {
		return err
	}

	apiGroup := router.Group("/api")

	// Media Folder Info
	{
		apiGroup.POST("/media-folder-info", GetMediaFolderInfo)
	}

	// Favourites
	{
		// TODO
	}

	// TODO
	if err := router.Run(":8000"); err != nil {
		return err
	}

	return nil
}
