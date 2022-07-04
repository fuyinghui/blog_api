package main

import (
	"blog_api/db"
	"blog_api/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	log.Printf("准备进入Initconfig方法")
	InitConfig()
	log.Printf("执行完了Initconfig方法")
	db := db.InitDB()
	log.Printf("执行完了initDB方法")
	defer db.Close()
	r := gin.Default()
	r = router.CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
func InitConfig() {
	log.Printf("进入了Initconfig方法")
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("err")
	}
}
