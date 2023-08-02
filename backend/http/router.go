package http

import (
	"TravelGo/backend/http/controller"
	"TravelGo/backend/http/middleware"
	"TravelGo/backend/repository"
	"TravelGo/backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	ginServer *gin.Engine
}

func NewRouter() Router {
	server := gin.Default()
	server.Use(CORSMiddleware())
	return Router{ginServer: server}
}

func (r *Router) StartServer() {
	//logger middleware
	r.ginServer.Use(gin.Logger())
	gin.SetMode(gin.DebugMode)
	r.ginServer.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error", gin.H{
			"message": "Requested routing not exist",
		})
	})
	r.ginServer.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.HTML(http.StatusInternalServerError, "error", gin.H{
					"message": err,
				})
			}
		}()
		c.Next()
	})

	repoUser := repository.NewUserRepository()
	serviceUser := service.NewUserService(repoUser)
	controllerOnBoarding := controller.OnBoardingController{UserService: serviceUser}
	r.ginServer.POST("/signUp", controllerOnBoarding.PostSignUp)
	r.ginServer.POST("/signIn", controllerOnBoarding.PostSignIn)
	r.ginServer.GET("/signOut", controllerOnBoarding.GetSignOut)

	repoComment := repository.NewCommentRepository()
	serviceComment := service.NewCommentService(repoComment)
	controllerComment := controller.CommentController{CommentService: serviceComment}
	repoPost := repository.NewPostRepository()
	servicePost := service.NewPostService(repoPost)
	controllerPost := controller.PostController{PostService: servicePost, CommentService: serviceComment}
	r.ginServer.POST("/search", controllerPost.PostSearchPosts)
	r.ginServer.GET("/get_post", controllerPost.GetPosts)
	r.ginServer.POST("/post_detail", controllerPost.GetPostDetail)
	groupPost := r.ginServer.Group("post")
	groupPost.Use(middleware.RequireAuth(serviceUser))
	{
		groupPost.POST("/create", controllerPost.PostCreatePost)
		groupPost.POST("/edit", controllerPost.PostEditPost)
		groupPost.POST("/delete", controllerPost.PostDeletePost)
		groupPost.GET("/my_posts", controllerPost.GetUserPost)
	}

	groupComment := r.ginServer.Group("comment")
	groupComment.Use(middleware.RequireAuth(serviceUser))
	{
		groupComment.POST("/make", controllerComment.PostMakeComment)
		groupComment.POST("/delete", controllerComment.PostDeleteComment)
	}

	controllerMashUp := controller.MashUpController{}
	groupMeshUp := r.ginServer.Group("mashup")
	{
		groupMeshUp.GET("/forecast", controllerMashUp.GetCityTemp)
		groupMeshUp.GET("/city_pics", controllerMashUp.GetCityPics)
	}
	r.ginServer.Run(":9991")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
