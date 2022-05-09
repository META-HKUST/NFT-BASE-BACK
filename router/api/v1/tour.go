package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description  tour
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/tour"
// @Router       /tour [GET]
func Name16(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  what is nft
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/what-is-nft"
// @Router       /what-is-nft [GET]
func Name17(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  web tutorial
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/web-tutorial"
// @Router       /web-tutorial [GET]
func Name18(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  event-banner
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/event-banner"
// @Router       /event-banner [GET]
func Name19(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
