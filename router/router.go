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
		userRouter := routerV1.Group("/users")
		{
			userRouter.POST("/register", v1.Register)
			userRouter.POST("/update", v1.Update)
			userRouter.POST("/login", v1.Login)
			userRouter.GET("/activate", v1.Activate)
			userRouter.POST("/rerunEmail", v1.RerunEmail)
		}

		accountRouter := routerV1.Group("/:id")
		{
			// getUser
			// if username == id, return all information includes password
			// else return
			accountRouter.GET("/", v1.GetUser)

			// // changeProfile
			accountRouter.POST("/change-profile", v1.ChangeProfile)

			// // collected
			accountRouter.GET("/collected", v1.Collected)

			// // favorites
			accountRouter.GET("/favorites", v1.Favorites)

			// creation
			// tab = item or collection
			accountRouter.GET("/creation", v1.Creation)

			// // createItem
			accountRouter.POST("/create-item", v1.CreateItem)

			// // editItem
			accountRouter.POST("/edit-item", v1.EditItem)

			// // editItem
			accountRouter.POST("/delete-item", v1.DeleteItem)

			// // createCollection
			accountRouter.POST("/create-collection", v1.CreateCollectionByAccount)

			// // editCollection
			accountRouter.POST("/edit-collection", v1.EditCollection)

		}

		collectionRouter := routerV1.Group("/collections")
		{
			// get all collections
			collectionRouter.GET("/", v1.GetAllCollections)

			//create collection
			collectionRouter.POST("/create", v1.CreateCollection)

			// check one collection
			collectionRouter.GET("/:collection-id", v1.GetCollectionByID)

			// items
			collectionRouter.GET("/:collection-id/items", v1.GetAllItems)

		}

		itemsRouter := routerV1.Group("/items")
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

			eventRouter.Use(mw.JWTAuth())

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
		itemJsonRouter := routerV1.Group("/tokenid")
		{

			itemJsonRouter.GET("/:tokenid", v1.GetJsonMsg)
		}

		testRouter := routerV1.Group("/test")
		{
			testRouter.POST("/", v1.TestContract)
		}	
	}
	return router
}
