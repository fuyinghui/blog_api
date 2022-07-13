package controller

import (
	"blog_api/db"
	"blog_api/dto"
	"blog_api/model"
	"blog_api/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

func ArticleUpdate(c *gin.Context) {
	log.Println("进入了ArticleUpdate")
	db := db.GetDB()
	var requestArticle = model.Article{}
	log.Println(requestArticle.ID)
	log.Println(requestArticle.Title)
	c.Bind(&requestArticle)
	var article model.Article
	db.Debug().Model(&article).Where("id=?", requestArticle.ID).Updates(map[string]interface{}{"title": requestArticle.Title, "content": requestArticle.Content, "status": 0, "updated_at": time.Now().Format("2006-01-02 15:04:05")})
	response.Response(c, http.StatusOK, 200, gin.H{"article": dto.ToArticleInfoDto(article)}, "更新成功!")

}
func ArticleDel(c *gin.Context) {
	db := db.GetDB()
	id := c.Query("id")
	var article model.Article
	db.Where("id=?", id).Delete(&article)
	response.Response(c, http.StatusOK, 200, nil, "删除成功!")
}
func ArticleInfo(c *gin.Context) {
	db := db.GetDB()
	id := c.Query("id")
	var article model.Article
	db.Where("id=?", id).Find(&article)
	response.Response(c, http.StatusOK, 200, gin.H{"article": dto.ToArticleInfoDto(article)}, "success!")
	//log.Println(article)
}
func ArticleListInfo(c *gin.Context) {
	log.Println("进入了ArticleListInfo")
	db := db.GetDB()
	var articles []model.Article
	db.Order("created_at desc").Find(&articles)
	//log.Println(dto.ToArticleArrayDto(articles))
	response.Response(c, http.StatusOK, 200, gin.H{"articles": dto.ToArticleArrayDto(articles)}, "success!")
}
func isArticleExist(db *gorm.DB, title string) bool {
	var article model.Article
	db.Where("title = ?", title).First(&article)
	if article.ID != 0 {
		return true
	}
	return false
}
func SubmitArticle(c *gin.Context) {
	db := db.GetDB()
	var article = model.Article{}
	c.Bind(&article)
	if len(article.Title) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "请先保存")
		return
	}

	if article.Status == 1 {
		response.Response(c, http.StatusOK, 200, nil, "已是发布状态，不需要再次发布！")
	} else {
		db.Model(&article).Where("title = ?", article.Title).Update("status", 1)
		log.Println(gin.H{"article": dto.ToArticleInfoDto(article)})
		response.Response(c, http.StatusOK, 200, gin.H{"article": dto.ToArticleInfoDto(article)}, "发布成功！")
	}

}

func AddArticle(c *gin.Context) {
	log.Println("进入了AddArticle方法")
	db := db.GetDB()
	//获取参数
	var article = model.Article{}
	//log.Println(requestUser)
	c.Bind(&article)
	//log.Println(requestArticle.Title)
	if len(article.Title) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "标题不能为空")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "标题不能为空",
		})*/
		return
	}
	if len(article.Content) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "内容不能为空")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "内容不能为空",
		})*/
		return
	}

	//判断文章是否已经存在
	if isArticleExist(db, article.Title) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "文章标题已存在！")
	} else {
		db.Create(&article)
		//log.Println(gin.H{"article": dto.ToArticleInfoDto(article)})
		response.Response(c, http.StatusOK, 200, nil, "新增文章成功！")
	}

}
