package middleware

import (
	"blog_api/db"
	"blog_api/model"
	"blog_api/response"
	"blog_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		//log.Println("开始执行Authmiddleware")
		tokenString := c.GetHeader("Authorization")
		//log.Println("tokenString:" + tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足1,请先登录!")
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足2,请先登录!")
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			c.Abort()
			return
		}
		//验证通过，获取claims中的userid
		userId := claims.UserId
		db := db.GetDB()
		var user model.User
		db.First(&user, userId)
		//用户不存在
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足3,请先登录!")
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足3"})
			c.Abort()
			return
		}
		//用户存在
		c.Set("user", user)
		c.Next()
	}
}
