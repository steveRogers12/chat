package mongolib

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"os"
)

var (
	Mgsession *mgo.Session
	Mgdatabase *mgo.Database
	MgoError error
)

func Init(dbname string) (*mgo.Session, *mgo.Database) {
	Mgsession, MgoError = mgo.Dial(fmt.Sprintf("%s:%s", viper.GetString("mongod-master.addr"), viper.GetString("mongod-master.port")))
	if MgoError != nil {
		fmt.Println("mongo 链接失败！")
		os.Exit(1)
	}
	// 选择DB
	Mgdatabase = Mgsession.DB(dbname)
	// 登陆
	uname := viper.GetString("mongod-master.username")
	if uname != "" {
		MgoError = Mgdatabase.Login(viper.GetString("mongod-master.username"), viper.GetString("mongod-master.password"))
		if MgoError != nil {
			fmt.Println("mongo 登陆验证失败！")
			os.Exit(1)
		}
	}
	// defer Session.Close()
	return Mgsession, Mgdatabase
}