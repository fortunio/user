package routes

import (
	_ "oosa/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	AuthRoutes(r)
	UserRoutes(r)
	OosaUserRoutes(r)
	PaymentRoutes(r)
	ContactUsRoutes(r)
	OosaDailyRoutes(r)
	WorldRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
