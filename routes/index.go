package routes

import (
	"archie/utils"
	"archie/utils/configer"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Serve() {
	config := configer.LoadServeConfig()

	router := gin.Default()
	router.Use(cors.Default())

	userRouter(router)
	organizationRouter(router)
	DocRouter(router)
	TodoRouter(router)

	utils.Logger(fmt.Sprintf("Listing on %s", config.Port))
	router.Run(config.GetAddress())
}
