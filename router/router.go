package router

import (
	docs "NFT-BASE-BACK/docs"
	mw "NFT-BASE-BACK/middleware"
	v1 "NFT-BASE-BACK/router/api/v1"
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
	routerV1 := router.Group("/api/v1")
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
	itemRouter := routerV1.Group("/item")
	{

		itemRouter.POST("/create", v1.CreateItem)

		itemRouter.POST("/edit", v1.EditItem)

		itemRouter.POST("/transfer", v1.TransferItem)
		itemRouter.POST("/like", v1.LikeItem)

	}

	listRouter := routerV1.Group("/list")
	{

		listRouter.GET("/user-list", v1.UserList)

		listRouter.POST("/collection", v1.SingleCollection)

		listRouter.POST("/collection-list", v1.CollectionList)
		listRouter.POST("/item", v1.SingleItem)
		listRouter.POST("/item-list", v1.ItemList)
		listRouter.POST("/item-history", v1.ItemHistory)

	}
	return router
}
