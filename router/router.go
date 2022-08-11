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
	//router.Use(mw.TlsHandler())
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
		userRouter.POST("/update-passwd", mw.JWTAuth(), v2.Update_Passwd)
		userRouter.POST("/forget-passwd", v2.Forget_Passwd)
		userRouter.POST("/reset-passwd", v2.Reset_Passwd)
		//用户个人信息
		userRouter.POST("/edit-profile", mw.JWTAuth(), v2.Edit_Profile)
		userRouter.GET("/info", v2.GetUserInfo)
	}
	collectionRouter := routerV2.Group("/collection").Use(mw.JWTAuth())
	{
		collectionRouter.POST("/create", v2.Create)
		collectionRouter.POST("/edit", v2.Edit)
	}

	itemsRouter := routerV2.Group("/item").Use(mw.JWTAuth())
	{
		itemsRouter.POST("/create", v2.CreateItem)
		itemsRouter.POST("/edit", v2.EditItem)
		itemsRouter.POST("/transfer", v2.TransferItem)
		itemsRouter.POST("/like", v2.LikeItem)
		itemsRouter.POST("/update", v2.UpdateItem)
	}
	tkRouter := routerV2.Group("/tk").Use(mw.JWTAuth())
	{
		tkRouter.POST("/transfer", mw.JWTAuth(), v2.PostTokenTransfer)
		tkRouter.GET("/info", mw.JWTAuth(), v2.GetTokenInfo)
	}

	listsRouter := routerV2.Group("/list")
	{
		listsRouter.GET("/user-list", v2.UserList)
		listsRouter.GET("/collection", v2.SingleColletction)
		listsRouter.GET("/collection-list", v2.CollectionList)
		listsRouter.GET("/item", mw.GetUserInfo(), v2.SingleItem)
		listsRouter.GET("/item-list", mw.GetUserInfo(), v2.ItemList)
		listsRouter.GET("/item-history", v2.ItemHistory)

	}

	actRouter := routerV2.Group("/act")
	{
		actRouter.POST("/create", mw.JWTAuth(), v2.PostActCreate)
		actRouter.POST("/delete", mw.JWTAuth(), v2.PostActDelete)
		actRouter.POST("/edit", mw.JWTAuth(), v2.PostActEdit)
		actRouter.GET("/info", v2.GetActInfo)
		actRouter.POST("/upload-item", mw.JWTAuth(), v2.PostActUploadItem)
		actRouter.GET("/item-list", mw.GetUserInfo(), v2.GetActItemList)
		actRouter.POST("/vote", mw.JWTAuth(), v2.PostActVote)
		actRouter.GET("/act-count", v2.GetActCount)
		actRouter.GET("/all-action", v2.GetAllAct)
		actRouter.GET("/can-upload", v2.CanUpload)
	}

	uploadRouter := routerV2.Group("/upload")
	{
		uploadRouter.POST("/cos", v2.UploadToCos)
		uploadRouter.POST("/ipfs-and-cos", v2.UploadToIpfs)
	}
	UserListRouter := routerV2.Group("/userList")
	{
		UserListRouter.POST("/add_white", mw.JWTAuth(), v2.AddWhiteList)
		UserListRouter.POST("/add_black", mw.JWTAuth(), v2.AddBlackList)
		UserListRouter.POST("/delete_white", mw.JWTAuth(), v2.DeleteWhiteList)
		UserListRouter.POST("/delete_black", mw.JWTAuth(), v2.DeleteBlackList)
		UserListRouter.GET("/search_white", mw.JWTAuth(), v2.SearchWhiteList)
		UserListRouter.GET("/search_black", mw.JWTAuth(), v2.SearchBlackList)
	}
	routerV2.GET("/metadata", v2.GetMetaInfo)
	return router
}
