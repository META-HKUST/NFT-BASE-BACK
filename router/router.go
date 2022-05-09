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
	docs.SwaggerInfo.BasePath = "/api/v1"
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerV1 := router.Group("/api/v1")
	{
		userRouter := routerV1.Group("/users")
		{
			userRouter.POST("/", v1.Name21)
			userRouter.POST("/login", v1.Name22)
		}

		accountRouter := routerV1.Group("/:id").Use(mw.AuthMiddleware)
		{
			// getUser
			// if username == id, return all information includes password
			// else return
			accountRouter.GET("/", v1.Name1)

			// // changeProfile
			accountRouter.POST("/change-profile", v1.Name5)

			// // collected
			accountRouter.GET("/collected", v1.Name2)

			// // favorites
			accountRouter.GET("/favorites", v1.Name3)

			// creation
			// tab = item or collection
			accountRouter.GET("/creation", v1.Name4)

			// // createItem
			accountRouter.POST("/create-item", v1.Name15_1)

			// // editItem
			accountRouter.POST("/edit-item", v1.Name15_2)

			// // delete  ？？

			// // createCollection
			accountRouter.POST("/create-collection", v1.Name9_1)

			// // editCollection
			accountRouter.POST("/edit-collection", v1.Name9_2)

		}

		collectionRouter := routerV1.Group("/collections")
		{
			// get all collections
			collectionRouter.GET("/", v1.GetAllCollectionsByID)

			//create collection
			collectionRouter.POST("/create", v1.CreateCollection)

			// check one collection
			collectionRouter.GET("/:collection-id", v1.GetCollectionByID)

			// items
			collectionRouter.GET("/:collection-id/items", v1.GetAllItems)

		}

		itemsRouter := routerV1.Group("/v1/items")
		{
			// all items
			// tab = sortBy & filter
			itemsRouter.GET("/", v1.AllItems)

			itemsRouter.GET("/:item", v1.SingleItem)

		}

		eventRouter := routerV1.Group("/events")
		{
			// all events
			eventRouter.GET("/", v1.AllEvents)

			eventRouter.GET("/:event-id", v1.SingleEvent)

			eventRouter.GET("/:event-id/items", v1.EventItems)

			eventRouter.GET("/:event-id/ranks", v1.EventItemsRank)

			eventRouter.GET("/:event-id/likes", v1.EventLikes)

			eventRouter.Use(mw.AuthMiddleware)

			eventRouter.POST("/:event-id/join", v1.JoinEvent)

			// choose
			eventRouter.POST("/:event-id/submit-item", v1.SubmitItem)

		}

		tourRouter := routerV1.Group("/tour")
		{
			tourRouter.GET("/", v1.GetAllTutorials)

			tourRouter.GET("/tr/articles", v1.GetAllArticles)

			tourRouter.POST("/tr/articles/:articles-id", v1.GetArticleByID)

			tourRouter.GET("/event-banner", v1.Name19)
		}
	}

	return router
}
