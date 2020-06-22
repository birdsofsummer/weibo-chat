package db

//https://gobook.io/read/gitea.com/xorm/manual-en-US/chapter-05/1.conditions.html

import (
	//. "./spider"
   // _ "github.com/go-sql-driver/mysql"
  _ "github.com/lib/pq"
    "xorm.io/xorm"
)

func add(engine *xorm.EngineGroup){
	user := new(User)
	user.Name = "ddd" + Random()
	affected, err := engine.Insert(user)

	if err!=nil{
		println("eeee",err.Error())
	}
	println(affected, err)
}

func add1(engine *xorm.EngineGroup){
	users := make([]*User, 1)
	users[0] = new(User)
	users[0].Name = "name_" + Random()
	affected, err := engine.Insert(users)
	if err!=nil{
		println("eeee",err.Error())
	}
	println(affected, err)
}

func list(engine *xorm.EngineGroup) []User{
	var allusers []User
	err := engine.Find(&allusers)
	if err!=nil{
		println("oo",err.Error())
	}
	return allusers
}

func Test(engine *xorm.EngineGroup){
	add(engine)
	add1(engine)
	u:=list(engine)
	for _,i :=range u {
		println(i.Id,i.Name)
	}
//	engine.DropTables(&User)
//	engine.CreateTables()
//  engine.DBMetas() 

}


