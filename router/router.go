package router

import (
	"blog_api/controller"
	"blog_api/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.PUT("/api/auth/articleUpdate", middleware.AuthMiddleware(), controller.ArticleUpdate)
	r.DELETE("/api/auth/articleDel", middleware.AuthMiddleware(), controller.ArticleDel)
	r.GET("/api/auth/articleInfo", middleware.AuthMiddleware(), controller.ArticleInfo)
	r.GET("/api/auth/articleList", middleware.AuthMiddleware(), controller.ArticleListInfo)
	r.POST("/api/auth/submitArticle", middleware.AuthMiddleware(), controller.SubmitArticle)
	r.POST("/api/auth/addArticle", middleware.AuthMiddleware(), controller.AddArticle)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r

}
