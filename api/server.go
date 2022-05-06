package api

import (
	"NFT-BASE-BACK/db"
	docs "NFT-BASE-BACK/docs"
	"NFT-BASE-BACK/util"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Server serves HTTP requests for our service.
type Server struct {
	config util.Config
	store  db.Store // interface or struct or pointer
	router *gin.Engine
	// tokenMaker ??
}

// accessToken + refreshToken
// authMiddleware

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
		// tokenMaker ??
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		userRouter := v1.Group("/users")
		{
			userRouter.POST("/", server.name21)
			userRouter.POST("/login", server.name22)
		}

		accountRouter := v1.Group("/:id").Use(authMiddleware)
		{
			// getUser
			// if username == id, return all information includes password
			// else return
			accountRouter.GET("/", server.name1)

			// // changeProfile
			accountRouter.POST("/change-profile", server.name5)

			// // collected
			accountRouter.GET("/collected", server.name2)

			// // favorites
			accountRouter.GET("/favorites", server.name3)

			// creation
			// tab = item or collection
			accountRouter.GET("/creation", server.name4)

			// // createItem
			accountRouter.POST("/create-item", server.name15_1)

			// // editItem
			accountRouter.POST("/edit-item", server.name15_2)

			// // delete  ？？

			// // createCollection
			accountRouter.POST("/create-collection", server.name9_1)

			// // editCollection
			accountRouter.POST("/edit-collection", server.name9_2)

		}

		collectionRouter := v1.Group("/collections")
		{
			// all collections
			collectionRouter.GET("/", server.name10_1)

			// one collection
			collectionRouter.GET("/:collection-id", server.name10_2)

			// items
			collectionRouter.GET("/:collection-id/items", server.name10_3)

		}

		itemsRouter := router.Group("/v1/items")
		{
			// all items
			// tab = sortBy & filter
			itemsRouter.GET("/", server.name11_1)

			itemsRouter.GET("/:item", server.name11_2)

		}

		eventRouter := v1.Group("/events")
		{
			// all events
			eventRouter.GET("/", server.name12_1)

			eventRouter.GET("/:event-id", server.name12_2)

			eventRouter.GET("/:event-id/items", server.name12_3)

			eventRouter.GET("/:event-id/ranks", server.name12_4)

			eventRouter.GET("/:event-id/likes", server.name12_5)

			eventRouter.Use(authMiddleware)

			eventRouter.POST("/:event-id/join", server.name12_6)

			// choose
			eventRouter.POST("/:event-id/submit-item", server.name12_7)

		}

		tourRouter := v1.Group("/tour")
		{
			tourRouter.GET("/", server.name16)

			tourRouter.GET("/what-is-nft", server.name17)

			tourRouter.GET("/web-tutorial", server.name18)

			tourRouter.GET("/event-banner", server.name19)
		}
	}
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {

	return server.router.Run(address)
}

func authMiddleware(ctx *gin.Context) {
	ctx.Next()
}

// @Description  get user
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna"
// @Router       /{id} [GET]
func (server *Server) name1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  collected
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/collected"
// @Router       /{id}/collected [GET]
func (server *Server) name2(ctx *gin.Context) {
	//

	//

	//
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  favorites
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/favorites"
// @Router       /{id}/favorites [GET]
func (server *Server) name3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  creation
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/creation"
// @Router       /{id}/creation [GET]
func (server *Server) name4(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  change profile
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 username   query   string   false   "user name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/change-profile"
// @Router       /{id}/change-profile [POST]
func (server *Server) name5(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  create-collection
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 collection-name   query   string   false   "collection name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/create-collection"
// @Router       /{id}/create-collection [POST]
func (server *Server) name9_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  edit-collection
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 collection-name   query   string   false   "collection name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/edit-collection"
// @Router       /{id}/edit-collection [POST]
func (server *Server) name9_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  create-item
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 item-name   query   string   false   "item name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/create-item"
// @Router       /{id}/create-item [POST]
func (server *Server) name15_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  edit-item
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 item-name   query   string   false   "item name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/edit-item"
// @Router       /{id}/edit-item [POST]
func (server *Server) name15_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all collections
// @Tags         collection
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections"
// @Router       /collections [GET]
func (server *Server) name10_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  single collection
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections/hahahah"
// @Router       /collections/{collection-id} [GET]
func (server *Server) name10_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all items in collection
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections/hahahah/items"
// @Router       /collections/{collection-id}/items [GET]
func (server *Server) name10_3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all items
// @Tags         item
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/items"
// @Router       /items [GET]
func (server *Server) name11_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  single item
// @Tags         item
// @param 		 item-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/items/yiiiiiii"
// @Router       /items/{item-id} [GET]
func (server *Server) name11_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all events
// @Tags         event
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/events"
// @Router       /events [GET]
func (server *Server) name12_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  single event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/events/balala"
// @Router       /events/{event-id} [GET]
func (server *Server) name12_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  items in event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/events/balala/items"
// @Router       /events/{event-id}/items [GET]
func (server *Server) name12_3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  item ranks in event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/events/balala/ranks"
// @Router       /events/{event-id}/ranks [GET]
func (server *Server) name12_4(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  item likes in event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/events/balala/likes"
// @Router       /events/{event-id}/likes [GET]
func (server *Server) name12_5(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  user join event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/events/balala/join"
// @Router       /events/{event-id}/join [POST]
func (server *Server) name12_6(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  user submit item
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @param        item-id    query  string    true    "item id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/events/balala/submit-item"
// @Router       /events/{event-id}/submit-item [POST]
func (server *Server) name12_7(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  tour
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/tour"
// @Router       /tour [GET]
func (server *Server) name16(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  what is nft
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/what-is-nft"
// @Router       /what-is-nft [GET]
func (server *Server) name17(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  web tutorial
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/web-tutorial"
// @Router       /web-tutorial [GET]
func (server *Server) name18(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  event-banner
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/event-banner"
// @Router       /event-banner [GET]
func (server *Server) name19(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  user enroll
// @Tags         user
// @param 		 email   query   string    true    "email"
// @param 		 password   query   string    true    "password"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users [POST]
func (server *Server) name21(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  user login
// @Tags         user
// @param 		 email   query   string    true    "email"
// @param 		 password   query   string    true    "password"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/login [POST]
func (server *Server) name22(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
