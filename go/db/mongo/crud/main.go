// https://www.cnblogs.com/Dr-wei/p/11742293.html

/*

Drop

InsertOne
InsertMany

FindOne
Find

UpdateOne
UpdateMany

ReplaceOne
ReplaceMany

DeleteOne
DeleteMany

*/

package crud

import (
	"context"
	"fmt"
	//"io/ioutil"
	//logger "log"
	"log"
	//"sync/atomic"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readconcern"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	//"go.mongodb.org/mongo-driver/mongo/writeconcern"

)

type mgo struct {
    uri string //数据库网络地址
    database string //要连接的数据库
    collection string //要连接的集合
}

type Student struct {
	Name string
	Age int
}

type Z struct{
	Name string
	Student []Student
}






func conn(m *mgo)(*mongo.Client,error){
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()
	client, err :=mongo.Connect(context.Background(), options.Client().ApplyURI(m.uri).SetMaxPoolSize(20))
    if err != nil {
        log.Print(err)
    }
	//defer client.Disconnect(ctx)
	return client, err
}

// --------------------------------------------------------------------------------

func add(collection *mongo.Collection,s1 Student) (interface {}){
	r, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}
	id:=r.InsertedID
	fmt.Println(id)
	return id
}

func adds(collection *mongo.Collection, s []interface{} ) ( []interface{}){
	insertManyResult, err := collection.InsertMany(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	ids:=insertManyResult.InsertedIDs
	fmt.Println(ids)
	return ids
}


func add1(collection *mongo.Collection,d primitive.M){
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	res,  err:= collection.InsertOne(ctx, d)
	if err!=nil{
		fmt.Println("9999")
		log.Fatal(err)
		return
	}
	id := res.InsertedID
	fmt.Println("222222",res,id)
}


func add2(collection *mongo.Collection,s1 Z) (interface {}){
	r, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}
	id:=r.InsertedID
	fmt.Println(id)
	return id
}



func del(collection *mongo.Collection, filter primitive.D) (*mongo.DeleteResult, error){
	return collection.DeleteMany( context.Background(),filter,)
}


func update(collection *mongo.Collection, filter primitive.D, d primitive.D) (*mongo.UpdateResult){
	 updateResult, err := collection.UpdateOne(context.TODO(), filter, d)
	 if err != nil {
		log.Fatal(err)
	 }
	 fmt.Printf("%d %d",updateResult.MatchedCount,updateResult.ModifiedCount) 
	 return updateResult
}


func findOne(collection *mongo.Collection,filter primitive.D) (Student){
	var result Student
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	return result
}


func list(collection *mongo.Collection) ([]*Student){
	filter := bson.D{{}}
	cur, err := collection.Find(
		context.Background(),
		filter,
	)
	if err != nil {
	   log.Fatal(err)
	}
	var results []*Student
	for cur.Next(context.TODO()) {
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	for k,v:=range(results){
		fmt.Printf("%d %v",k,v)
	}

	//fmt.Printf( "%#v\n", results)
	return results
}

// --------------------------------------------------------------------------------


func test() {
	m:=&mgo{
		//"mongodb://localhost:27017",
		"mongodb://localhost:27017/?retryWrites=true&w=majority&connect=direct",
		"testing",
		"numbers",
	}
	client, err := conn(m)
	if err!=nil{
		fmt.Println("eeeee")
		log.Fatal(err)
		return
	}

	db:=client.Database(m.database)
	coll:= db.Collection(m.collection)
	coll.Drop(context.Background())

	s1 := Student{"小红", 12}
	s2 := Student{"小兰", 10}
	s3 := Student{"小黄", 11}
	filter := bson.D{{"name", "小兰"}}
	d := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	d1:=bson.M{ "name": "pi", "value": 3.14159 }
	d2:=bson.D{ {"name", "pi"}, }
	add1(coll,d1)
    del(coll,d2)

	add(coll,s1)
	students := []interface{}{s2, s3}
	adds(coll,students)
	list(coll)
	findOne(coll,filter)
	update(coll, filter,d)
	findOne(coll,filter)
	list(coll)

	coll.Drop(context.Background())


	d3:=Z{"ddd",[]Student{s1,s2, s3}}
	add2(coll,d3)

	coll.Drop(context.Background())
}


func main(){
	//test()
}
