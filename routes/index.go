package routes

import (
	"archie/utils"
	"archie/utils/configer"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func Serve() {
	config := configer.LoadServeConfig()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authentication"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	//router.Use(func(c *gin.Context) {
	//	for k, v := range c.Request.Header {
	//		fmt.Println("")
	//		fmt.Println(k, v)
	//		fmt.Println("")
	//	}
	//
	//	c.JSON(200, gin.H{
	//		"data": "a",
	//		"err":  "err",
	//	})
	//})

	messageRoutes(router)
	uploadRouter(router)
	userRouter(router)
	organizationRouter(router)
	DocRouter(router)
	TodoRouter(router)

	utils.Logger(fmt.Sprintf("Listing on %s", config.Port))
	err := router.Run(config.GetAddress())

	if err != nil {
		log.Fatal(err)
	}
}
