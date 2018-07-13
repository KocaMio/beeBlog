package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "beeBlog/routers"
	_ "github.com/lib/pq"
)

func init() {
	beegoConf := beego.AppConfig

	// Register SQL Driver
	orm.RegisterDriver("postgres", orm.DRPostgres)
	
	// Register PostgresSQL
	orm.RegisterDataBase(
		"default", //dbname
		"postgres", //driver
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
			beegoConf.String("postgres::user"),
			beegoConf.String("postgres::password"),
			beegoConf.String("postgres::dbname"),
			beegoConf.String("postgres::host"),
			beegoConf.String("postgres::port"),
			beegoConf.String("postgres::sslmode"),
		),
	)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.InsertFilter("/article/read", beego.BeforeRouter, FilterSomething)

	beego.Run()
}

var FilterSomething = func(ctx *context.Context) {
	
}
