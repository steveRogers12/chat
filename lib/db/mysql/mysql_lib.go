package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"time"
)

var (
	MasterDB  *xorm.Engine
	err error
	Slave1DB  *xorm.Engine
	err1 error
)

func MasterInit() *xorm.Engine{
	open := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", viper.GetString("mysql-master.username"),
		viper.GetString("mysql-master.password"),
		viper.GetString("mysql-master.addr"),
		viper.GetInt64("mysql-master.port"),
		viper.GetString("mysql-master.database"))

	MasterDB, err = xorm.NewEngine("mysql", open)
	if err != nil {
		fmt.Printf("Open mysql-master failed,err:%v\n", err)
		panic(err)
	}

	MasterDB.SetConnMaxLifetime(100 * time.Second)
	MasterDB.SetMaxOpenConns(100)
	MasterDB.SetMaxIdleConns(16)
	err = MasterDB.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql-master, err:" + err.Error())
		panic(err.Error())
	}

	fmt.Printf("mysql-master connect success\r\n")
	return MasterDB
}

func Slave1Init() *xorm.Engine{
	open := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", viper.GetString("mysql-slave1.username"),
		viper.GetString("mysql-slave1.password"),
		viper.GetString("mysql-slave1.addr"),
		viper.GetInt64("mysql-slave1.port"),
		viper.GetString("mysql-slave1.database"))

	Slave1DB, err1 = xorm.NewEngine("mysql", open)
	if err != nil {
		fmt.Printf("Open mysql-slave1 failed,err:%v\n", err)
		panic(err)
	}

	Slave1DB.SetConnMaxLifetime(100 * time.Second)
	Slave1DB.SetMaxOpenConns(100)
	Slave1DB.SetMaxIdleConns(16)
	err1 = Slave1DB.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql-slave1, err:" + err1.Error())
		panic(err.Error())
	}

	fmt.Printf("mysql-slave1 connect success\r\n")
	return Slave1DB
}