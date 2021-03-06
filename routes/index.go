package routes

import (
	"archie/constants"
	"archie/resolver"
	"archie/schema"
	"archie/utils"
	"archie/utils/configer"
	"archie/utils/helper"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gql "github.com/graph-gophers/graphql-go"
	"log"
	"net/http"
)

type Params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

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

	parsedSchema := gql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{}, gql.UseFieldResolvers())

	router.Any("graphql", func(ctx *gin.Context) {
		res := helper.Res{}
		params := Params{}

		if err := ctx.Bind(&params); err != nil {
			res.Status(http.StatusBadRequest).Error(err).Send(ctx)
			return
		}

		newCtx := context.WithValue(ctx.Request.Context(), constants.GIN_CONTEXT, ctx)

		response := parsedSchema.Exec(newCtx, params.Query, params.OperationName, params.Variables)

		if len(response.Errors) != 0 {
			res.Error(response.Errors).Send(ctx)
			return
		}

		res.Success(response.Data).Send(ctx)
	})

	//messageRoutes(router)
	//uploadRouter(router)
	//userRouter(router)
	//organizationRouter(router)
	//docRouter(router)
	//todoRouter(router)
	//permissionRouter(router)
	//categoryRouter(router)

	utils.Logger(fmt.Sprintf("Listing on %s", config.Port))
	err := router.Run(config.GetAddress())

	if err != nil {
		log.Fatal(err)
	}
}
