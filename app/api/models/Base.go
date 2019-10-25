package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"net/url"
	"github.com/gin-gonic/gin"
	"web-go-skeleton/app/api/components"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

var DB *gorm.DB

type Base struct {
	gorm.Model
}

func DuplicateEntry(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	return strings.Contains(errStr,"Duplicate entry")
}

func init()  {
	user := components.Config.MustValue("mysql", "user")
	password := components.Config.MustValue("mysql", "password")
	host := components.Config.MustValue("mysql", "host","127.0.0.1")
	port := components.Config.MustValue("mysql", "port","3306")
	dbName := components.Config.MustValue("mysql", "dbname")
	prefix := components.Config.MustValue("mysql", "prefix")
	timezone := components.Config.MustValue("mysql", "timezone")

	fmt.Println(timezone)
	if len(user) == 0 || len(password) == 0 || len(dbName) == 0 {
		panic("缺少必要的数据库配置")
	}

	//设置表名前缀
	if len(prefix) > 0 {
		gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
			return prefix + defaultTableName
		}
	}

	//初始化db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, dbName)
	if len(timezone) > 0 {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil{
		panic(fmt.Sprintf("connect to mysql failed,  %v", err))
	}else{
		fmt.Println("connect to mysql suceed")
	}

	DB.DB().SetMaxIdleConns(10)  //set max idle conn nums
	DB.DB().SetMaxOpenConns(100) //set max open conn nums
	if gin.IsDebugging() {
		DB.LogMode(true)         // set sql log on
	}
}