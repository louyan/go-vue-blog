package main

import (
	_ "vue-blog-api/routers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}


	//root:root123@tcp(127.0.0.1:3306)/vue_blog?charset=utf8
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn,30)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "DELETE", "PUT", "PATCH", "OPTIONS","POST"},
		AllowHeaders: []string{"Origin", "Access-Control-Allow-Origin","Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	
	beego.Run()
}

