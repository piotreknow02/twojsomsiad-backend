package router

import (
	"net/http"

	"twojsomsiad/controller"
	_ "twojsomsiad/docs"
	"twojsomsiad/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// TODO: Update config and move this to middleware

	r.Use(middleware.GZIP())

	r.Use(middleware.CORS())

	// Controller
	api := controller.Controller{DB: db}
	advert := r.Group("/adverts")
	{
		advert.GET("/", api.Adverts)
	}

	// Authorization
	authorization := middleware.Auth(&api)
	// Authorization middleware for later use
	authMiddleware := authorization.MiddlewareFunc()

	// Swagger
	sg := r.Group("/")
	{
		sg.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		sg.GET("/swagger/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger-ui/index.html")
		})
	}

	r.Use(middleware.CSP())

	// Endpoints
	auth := r.Group("/auth")
	{
		auth.POST("/login", authorization.LoginHandler)
		auth.GET("/refresh", authorization.RefreshHandler)
		auth.POST("/register", api.Register)
	}

	user := r.Group("/user")
	{
		user.GET("/:id", api.GetUser)
		protected := user.Group("/")
		{
			protected.Use(authMiddleware)
			protected.GET("/", api.GetMyUser)
			protected.POST("/", api.UpdateUser)
		}
	}

	return r
}
