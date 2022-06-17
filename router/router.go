package router

import (
	docs "NFT-BASE-BACK/docs"
	mw "NFT-BASE-BACK/middleware"
	v1 "NFT-BASE-BACK/router/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// accessToken + refreshToken
// authMiddleware
// initRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(mw.Cors())
	docs.SwaggerInfo.BasePath = "/api/v1"
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerV1 := router.Group("/api/v1")
	{

		itemRouter := routerV1.Group("/item")
		{

			itemRouter.POST("/create", v1.CreateItem)

			itemRouter.POST("/edit", v1.EditItem)

			itemRouter.POST("/transfer", v1.TransferItem)
			itemRouter.POST("/like", v1.LikeItem)

		}

		listRouter := routerV1.Group("/list")
		{

			listRouter.GET("/user-list",v1.UserList)

			listRouter.POST("/collection", v1.SingleCollection)

			listRouter.POST("/collection-list", v1.CollectionList)
			listRouter.POST("/item", v1.SingleItem)
			listRouter.POST("/item-list", v1.ItemList)
			listRouter.POST("/item-history", v1.ItemHistory)

		}
	}
	return router
}
