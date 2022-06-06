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
		userRouter := routerV1.Group("/users").Use(mw.Cors())
		{
			userRouter.POST("/register", v1.Register)
			userRouter.POST("/login", v1.Login)
			userRouter.POST("/update-passwd", v1.Update)

			userRouter.GET("/activate", v1.Activate)
			userRouter.POST("/rerun-email", v1.RerunEmail)
			userRouter.POST("/forget-passwd", v1.ForgetPasswd)
			userRouter.POST("/reset-passwd", v1.ResetPasswd)
			userRouter.POST("/delete-user", v1.DeleteUser)
		}

		accountRouter := routerV1.Group("/account").Use(mw.JWTAuth()).Use(mw.Cors())
		{
			// createItem
			accountRouter.POST("/create-item", v1.CreateItem)
			//  editItem
			accountRouter.POST("/edit-item", v1.EditItem)
			// collected
			accountRouter.GET("/collected", v1.Collected)
			// favorites
			accountRouter.GET("/favorites", v1.Favorites)
			// deleteItem
			accountRouter.POST("/delete-item", v1.DeleteItem)

			// createCollection
			accountRouter.POST("/create-collection", v1.CreateCollectionByAccount)
			// editCollection
			accountRouter.POST("/edit-collection", v1.EditCollection)
			accountRouter.GET("/creation", v1.Creation)
			// deleteCollection
			accountRouter.POST("/delete-collection", v1.DeleteCollection)

			// EditProfile
			accountRouter.POST("/edit-profile", v1.EditProfile)
			// getUser
			accountRouter.GET("/get-user", v1.GetUser)

		}

		collectionRouter := routerV1.Group("/collections").Use(mw.Cors())
		{
			// get all collections
			collectionRouter.GET("/", v1.GetAllCollections)

			// check one collection
			collectionRouter.GET("/:collection-id", v1.GetCollectionByID)

		}

		itemsRouter := routerV1.Group("/items").Use(mw.Cors())
		{
			// all items
			// tab = sortBy & filter
			itemsRouter.GET("/", v1.SortedItems)

			itemsRouter.GET("/:item", v1.SingleItem)

			itemsRouter.GET("/item-json", v1.GetJsonMsg)

		}

		eventRouter := routerV1.Group("/events").Use(mw.Cors())
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

		tourRouter := routerV1.Group("/tour").Use(mw.Cors())
		{
			tourRouter.GET("/", v1.GetAllTutorials)

			tourRouter.GET("/tr/articles", v1.GetAllArticles)

			tourRouter.POST("/tr/articles/:articles-id", v1.GetArticleByID)

			tourRouter.GET("/event-banner", v1.Name19)
		}

		testRouter := routerV1.Group("/test").Use(mw.Cors())
		{
			testRouter.POST("/", v1.TestContract)
			testRouter.POST("/enroll", v1.TestEnroll)
		}
	}

	return router
}
