package db

import (
	//. "./spider"
   // _ "github.com/go-sql-driver/mysql"
  _ "github.com/lib/pq"
    "xorm.io/xorm"
	"math/rand"
	"time"
	"fmt"
)

func Random() string{
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Intn(254))
}

//var engine *xorm.Engine
var engine *xorm.EngineGroup 

func Conn()(error, *xorm.EngineGroup){
	var engine *xorm.EngineGroup 
    var err error
	conns := []string{
		"postgres://postgres:postgres@localhost:5432/chat1?sslmode=disable;", 
		"postgres://postgres:postgres@localhost:5432/chat2?sslmode=disable;",
		"postgres://postgres:postgres@localhost:5432/chat3?sslmode=disable",
	}

  //  engine, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
    engine, err = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())

	if err != nil {
		println(err.Error())
		return err,engine
	}
	//engine.SetMapper(names.SameMapper{})
	//engine.SetTableMapper(names.SameMapper{})
	//engine.SetColumnMapper(names.SnakeMapper{})
	err = engine.Sync2(new(User))
	if err!=nil{
		println("eeee",err.Error())
		return err,engine
	}
	return err,engine
}

