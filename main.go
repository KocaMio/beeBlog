package main

import (
	_ "beeBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	
	// Register SQL Driver
	orm.RegisterDriver("postgres", orm.DRPostgres)
	
	// Register PostgresSQL
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=miochang password=hungshih dbname=postgres host=127.0.0.1 port=5432 sslmode=disable",
	)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.InsertFilter("/article/read", beego.BeforeRouter, FilterSomething)

	beego.Run()
}

var FilterSomething = func(ctx *context.Context) {
	
}
