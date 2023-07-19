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
	return Router{ginServer: gin.Default()}
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
	r.ginServer.GET("/post_detail", controllerPost.GetPostDetail)
	groupPost := r.ginServer.Group("post")
	groupPost.Use(middleware.RequireAuth(serviceUser))
	{
		groupPost.POST("/create", controllerPost.PostCreatePost)
		groupPost.POST("/edit", controllerPost.PostEditPost)
		groupPost.POST("/delete", controllerPost.PostDeletePost)
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
