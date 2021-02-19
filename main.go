package main

import (
	mongolib "ChatQueue/lib/db/mongo"
	"ChatQueue/lib/db/mysql"
	"ChatQueue/lib/queue"
	"ChatQueue/lib/redislib"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)



func main() {
	// 初始化配置文件
	initConfig()
	// 初始化日志
	initLog()
	// 初始化redis
	redislib.Sclient()
	// 初始化mongod
	Mongos, Mongoc := mongolib.Init(viper.GetString("mongod-master.dbname"))
	fmt.Println("Mongo is connected ", Mongoc.Name)
	defer Mongos.Close()

	// 初始化数据库
	MasterDB := mysql.MasterInit()
	defer MasterDB.Close()
	Slave1DB := mysql.Slave1Init()
	defer Slave1DB.Close()

	//kafkaFunc := router.KafkaRouter()
	//kafka
	queueConfig := new(queue.Consumer)
	queueConfig.BrokerServers = viper.GetString("kafka.hosts")
	queueConfig.Consume(nil)
}

// 读取配置文件
func initConfig() {
	viper.SetConfigName("config/config")
	viper.AddConfigPath(".") // 添加搜索路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}
	//fmt.Println("config app:", viper.Get("app"))
	//fmt.Println("config redis:", viper.Get("redis"))
}

// 初始化日志
func initLog() {
	// 关闭控制台颜色
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}