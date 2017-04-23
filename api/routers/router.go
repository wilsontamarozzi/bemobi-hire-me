package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/middleware"
)

func InitRoutes() *gin.Engine {
	routes := gin.New()

	routes.Use(gin.Logger())
	routes.Use(gin.Recovery())
	routes.Use(middleware.CORS())

	v1 := routes.Group("api/v1")
	{
		AddRoutesURL(v1)
	}

	return routes
}
