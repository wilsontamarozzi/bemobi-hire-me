package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/controllers"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/repositories"
)

func AddRoutesURL(r *gin.RouterGroup) {
	controller := controllers.URLController{Repository: repositories.NewURLRepository()}
	routes := r.Group("/url")
	{
		routes.POST("/shorten", controller.Create)
		routes.GET("/details/:alias", controller.GetByAlias)
		routes.GET("/ranking", controller.GetAllRanking)
	}
}
