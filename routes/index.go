package routes

import (
	"archie/utils"
	"archie/utils/configer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Serve() {
	config := configer.LoadServeConfig()

	router := gin.Default()

	//router.Use(cors.New(cors.Config{
	//	AllowOriginFunc:  func(origin string) bool { return true },
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	//	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	//	AllowCredentials: true,
	//}))

	userRouter(router)
	organizationRouter(router)
	DocRouter(router)

	utils.Logger(fmt.Sprintf("Listing on %s", config.Port))
	router.Run(config.GetAddress())
}
