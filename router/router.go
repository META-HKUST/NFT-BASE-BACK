package router

import (
	docs "NFT-BASE-BACK/docs"
	mw "NFT-BASE-BACK/middleware"
	v2 "NFT-BASE-BACK/router/api/v2"

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
	docs.SwaggerInfo.BasePath = "/api/v2"
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerV2 := router.Group("/api/v2")
	userRouter := routerV2.Group("/user")
	{
		//注册相关
		userRouter.POST("/register", v2.Register)
		userRouter.POST("/rerun-email", v2.Rerun_Email)
		userRouter.GET("/activate", v2.Activate)
		//登录
		userRouter.POST("/login", v2.Login)
		//密码管理
		userRouter.POST("/update-passwd", v2.Update_Passwd)
		userRouter.POST("/forget-passwd", v2.Forget_Passwd)
		userRouter.POST("/reset-passwd", v2.Reset_Passwd)
		//用户个人信息
		userRouter.POST("/edit-profile", v2.Edit_Profile)
		userRouter.GET("/info", v2.GetUserInfo)
	}
	collectionRouter := routerV2.Group("/collection")
	{
		collectionRouter.POST("/create", v2.Create)
		collectionRouter.POST("/edit", v2.Edit)
	}

	itemsRouter := routerV2.Group("/item")
	{
		itemsRouter.POST("/create", v2.CreateItem)
		itemsRouter.POST("/edit", v2.EditItem)
		itemsRouter.POST("/transfer", v2.TransferItem)
		itemsRouter.POST("/like", v2.LikeItem)
	}

	listsRouter := routerV2.Group("/list")
	{
		listsRouter.GET("/user-list", v2.UserList)
		listsRouter.GET("/collection", v2.SingleColletction)
		listsRouter.GET("/collection-list", v2.CollectionList)
		listsRouter.GET("/item", v2.SingleItem)
		listsRouter.GET("/item-list", v2.ItemList)
		listsRouter.GET("/item-history", v2.ItemHistory)

	}
	return router
}
