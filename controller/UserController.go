package controller

import (
	"blog_api/db"
	"blog_api/dto"
	"blog_api/model"
	"blog_api/response"
	"blog_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Info(c *gin.Context) {
	log.Println("进入了info方法")
	user, _ := c.Get("user")
	//log.Printf("user为：%v", user)
	response.Response(c, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "success!")
	//c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
func Login(c *gin.Context) {
	log.Println("进入了login方法")
	db := db.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//参数验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})*/
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位",
		})*/
		return
	}
	//判断手机号是否存在
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": 422,
		"msg":  "用户已存在"})*/
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		/*c.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  "密码错误"})*/
		return
	}
	//发放token
	token, err := utils.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		/*c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})*/
		log.Printf("token generate error: %v", err)
		return
	}
	//发会结果
	response.Response(c, http.StatusOK, 200, gin.H{"token": token}, "登录成功")
	/*c.JSON(200, gin.H{
	"code": 200,
	"data": gin.H{"token": token},
	"msg":  "登录成功"})*/
}
func Register(c *gin.Context) {
	db := db.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})*/
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位",
		})*/
		return
	}
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户名不能为空",
		})*/
		return
	}
	//判断用户是否存在
	log.Println(name, password, telephone)
	if isTelephoneExist(db, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		/*c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已存在",
		})*/
		return
	}
	//创建用户
	hasedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newuser := model.User{
		Name:      name,
		Password:  string(hasedpassword),
		Telephone: telephone,
	}
	db.Create(&newuser)
	//返回结果
	response.Response(c, http.StatusOK, 200, nil, "注册成功")
	/*c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功！",
	})*/
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
