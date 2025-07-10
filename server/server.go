package server

import (
	"fmt"
	"local-media-gallery/config"
	"net"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"127.0.0.0"}); err != nil {
		return err
	}
	router.Use(cors.Default())

	router.Static("/assets", "./web/dist/assets")
	router.Static("/favicon.ico/", "./web/dist/favicon.ico/")
	router.Static("/static", config.Config.MEDIA_HOST_PATH)

	apiGroup := router.Group("/api")

	// Media Folder Info
	{
		apiGroup.POST("/media-folder-info", GetMediaFolderInfo)
	}

	// Favourites
	{
		// TODO
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})

	printIpAddress()

	// TODO
	if err := router.Run(":80"); err != nil {
		return err
	}

	return nil
}

func printIpAddress() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		// Check if the address is a loopback address (127.0.0.1, ::1)
		// and skip it if it is.
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// Check if it's an IPv4 or IPv6 unicast address
			if ipnet.IP.To4() != nil {
				fmt.Println("Local IP:", ipnet.IP.String())
			}
		}
	}
}
