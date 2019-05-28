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

	userRouter(router)
	organizationRouter(router)

	utils.Logger(fmt.Sprintf("Listing on %s", config.Port))
	router.Run(config.GetAddress())
}
