package spider

import (

	//"crypto/hmac"
	//"crypto/sha1"
	//"crypto/sha256"
	//"strconv"
	//"encoding/base64"
	//"sort"
	"math/rand"
	"strings"
	"time"
	"os"
	"encoding/json"
	//"log"
	"net/url"
	"fmt"
	//"bufio"
	//"io"
	"io/ioutil"
	"net/http"
	. "../consts"
	. "../types"
	"../types/contact"
	"../types/query_group"
	"../types/batch_exists"
	"../types/conversation"
)

var (
   headers=GetHeaders()
   config=GetConfig()
)


func random() string{
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Intn(254))
}

func pick(s []string) string{
	rand.Seed(time.Now().UnixNano())
	l:=len(s)
	n:=rand.Intn(l)
	return s[n]
}

func ip() string{
	z:=[]string{}
	for i:=0;i<4;i++{
		t:=random()
		z=append(z,t)
	}
	return strings.Join(z , ".")
}


func now() string{
    return fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
}



func to_qs(u string,o map[string]string) string{
	 Url, _ := url.Parse(u)
	 p:=url.Values{}
	 for k, v := range o{
	 	p.Add(k,v)
	 }
	 Url.RawQuery = p.Encode()
	 return Url.String()
}

func to_qs1(o map[string]string) string{
	 p:=url.Values{}
	 for k, v := range o{
	 	p.Add(k,v)
	 }
	 return p.Encode()
}



func Read(file string) (string){
    data, err := ioutil.ReadFile(file)
	s:=""
    if err != nil {
        fmt.Println("File reading error", err)
        return s
    }
	s=string(data)
	//s=base64.StdEncoding.EncodeToString(data)
	return s
}


func ReadJSON(file string) (msg []Messages){
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("File reading error", err)
        return 
    }
	err = json.Unmarshal(data, &msg)
	if err!=nil {
			fmt.Println("错误")
	}
	return msg
}

func to_json(d map[string]string) string{
	m,_:=json.Marshal(d)
    return string(m)
}

func get_headers() map[string]string{
	//res.Header["Set-Cookie"]
	//z:=[]string{}
    //cookie:=strings.Join(z , "; ")
	//return HEADERS
	return GetHeaders()
}


func get(u string,d map[string]interface{}) (*http.Response,error){
     client := http.Client{}
	 p:=url.Values{}
	 for k, v := range d{
	 	p.Add(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	 }
	 qs:=p.Encode()
	 u1:=u+"?"+qs
	 fmt.Println("get",u1)
     req, _ := http.NewRequest("GET", u1,nil)
     //h:=get_headers()
     for k,v := range headers{
 	    req.Header.Add(k, v)
     }
	 //fmt.Println("hhhhhhhhh",h)
	 return client.Do(req)
}


func Get_contacts() (contact.ContactData, error){
	u:="https://api.weibo.com/webim/2/direct_messages/contacts.json"
	d:=map[string]interface{}{
		"special_source" : "3",
		"add_virtual_user" : "3,4",
		"is_include_group" : "0",
		"need_back" : "0,0",
		"count" : "50",
		"source" : "209678993",
		"t" : now(),
	}
	r,_:=get(u,d)
	b, _ := ioutil.ReadAll(r.Body)
	//fmt.Println("[contact]:",r.Status,string(b))
	return contact.UnmarshalContact(b)
}


func Get_gid(c contact.ContactData,err error) ([]int64){
	z:= []int64{}
	if err!=nil{
		fmt.Println("错误")
		return z
	}
	for _,v:=range(c.Contacts){
		fmt.Println(v.User.Name,v.User.ID,v.User.Type)
		if v.User.Type == 2 {
			z=append(z,v.User.ID)
		}
	}
	//fmt.Println(z)
	return z
}


func get_msg(max_mid int64,id int64)(Msg){
	u:="https://api.weibo.com/webim/groupchat/query_messages.json"
	d:=map[string]interface{}{
            "convert_emoji" : "1",
            "query_sender" : "1",
            "count" : "20",
            "id":id,
            "max_mid":max_mid,
            "source" : "209678993",
            "t" : now(),
    }
	r,_:=get(u,d)
	b, _ := ioutil.ReadAll(r.Body)
	//fmt.Println("[msg]:",r.Status,string(b))
	var msg Msg
	err := json.Unmarshal(b, &msg)
	if err!=nil {
		fmt.Println("错误")
	}
	return msg
}

//209678993
//封号
//{"error":"User does not exists!","error_code":20003,"request":"/2/account/profile/basic.json"}

func get_profile(s int64){
   u:="https://api.weibo.com/webim/2/account/profile/basic.json"
   d:=map[string]interface{}{
	   source: fmt.Sprintf("%d", s),
	   t:now(),
   }
   r,_:=get(u,d)
   b, _ := ioutil.ReadAll(r.Body)
   fmt.Println("[msg]:",r.Status,string(b))
}




// 7388859010
func get_batch_exists(uid int64[]) (batch_exists.Data, error){
	u:="https://api.weibo.com/webim/2/friendships/batch_exists.json"
	uids:= strings.Join(uid , ",")
	d:=map[string]interface{}{
		"uids" : uids, 
		"source" : "209678993", 
		"t" : now(), 
	}
    r,_:=get(u,d)
    b, _ := ioutil.ReadAll(r.Body)
   //fmt.Println("[msg]:",r.Status,string(b))
	return batch_exists.UnmarshalData(b)
}


// 7388859010
func show_person(uid int64) (show.Data, error){
	u:="https://api.weibo.com/webim/2/users/show.json"
	d:=map[string]interface{}{
		"uid" : uid, 
		"source" : "209678993", 
		"t" : now(), 
	}
    r,_:=get(u,d)
    b, _ := ioutil.ReadAll(r.Body)
   //fmt.Println("[msg]:",r.Status,string(b))
	return show.UnmarshalData(b)
}



// 7388859010 0
// 翻页与组消息一样 
// 第一页max_id=0 
// 之后修改max_id
func get_conversation(uid int64,max_id int64) (conversation.Data, error){
	u:="https://api.weibo.com/webim/2/direct_messages/conversation.json"
	d:=map[string]interface{}{
		"convert_emoji" : "1",
		"count" : "15",
		"max_id" : max_id,
		"uid" : uid,
		"is_include_group" : "0",
		"source" : "209678993",
		"t" : now()
	}
    r,_:=get(u,d)
    b, _ := ioutil.ReadAll(r.Body)
   //fmt.Println("[msg]:",r.Status,string(b))
	return conversation.UnmarshalData(b)
}



//4356260621621693
func get_group(id int64)(query_group.Data, error){
	u:="https://api.weibo.com/webim/query_group.json"
	d:=map[string]interface{}{
	  "is_pc" : "1",
	  "query_member" : "1",
	  "sort_by_jp" : "1",
	  "query_member_count" : "5000",
	  "id" : fmt.Sprintf("%d", id),
	  "source" : "209678993", //me
	  "t" : now(), 
	}
   r,_:=get(u,d)
   b, _ := ioutil.ReadAll(r.Body)
   //fmt.Println("[msg]:",r.Status,string(b))
   return query_group.UnmarshalData(b)
}









func test(){
	var max_mid int64
	max_mid=0

	var id int64
	//id=4429200364208496
	id=4356260621621693 

	r:=get_msg(max_mid,id)
	fmt.Println("[msg]:",r)
}

func save(d *[]Messages,n string){
	//b,_:=json.Marshal(d)
	b, _ := json.MarshalIndent(d, "", "\t")
	file,_:=os.Create(n)
	defer file.Close()
	_, err := file.Write(b)
	if err!=nil{
		fmt.Println("[saved] error")
	}
	fmt.Println("[saved]:",n)
}


func Start(id int64, from int64,to int64) *[]Messages{
	z:= []Messages{}
	i:=0
	fmt.Println("begin...",id,from,to)
	for {
		fmt.Println(from,to,i)
		if to>0 && int64(to)>=from{
			fmt.Println("[done]:",i,len(z))
			break
		}
		r:=get_msg(from,id)
		m:=r.Messages
		//fmt.Println("[msg]:",r)
		fmt.Println("[msg]:",len(m))
		if len(m)>0{
			z=append(z,m...)
			//from=r.LastReadMid-1 //
			from=m[0].ID-1 
		}else{
			fmt.Println("[done]:",i,len(m))
			break
		}
		i++
	}
	return &z
}

func Start1(){
	n:="json/msg1.json"
	n1:="json/msg1-1.json"
	t:=ReadJSON(n)
	last:=t[0].ID-1
	fmt.Println("[last]:",last)
	id:=config.ID
	from:=int64(0)
	z:=Start(id,from,last)
	z1:=append(*z, t...)
	save(&z1,n1)
	fmt.Println("append done!")
}


func Start2(id int64){
	n:=fmt.Sprintf("json/%d.json", id)
	from:=int64(0)
	last:=int64(0)
	z:=Start(id,from,last)
	save(z,n)
	fmt.Println("fetch all done!")
}

func test1(){
	var max_mid int64
	id:=config.ID
	max_mid=0
	z:=Start(id,max_mid,0)
	//fmt.Println("[done]:",len(*z))
	//fmt.Println("[done]:",z)
	n:="json/msg1.json"
	save(z,n)
}




func GetGroupMsg(){
	g:=Get_gid(Get_contacts())
	fmt.Println(g)
	c := make(chan int64, len(g))
	go (func(){
		for i,id:=range(g){
			fmt.Println("start",i,id)
			Start2(id)
			fmt.Println("done",i,id)
			c <- id
		}
	})()
	for i := range c {
		fmt.Println(i,"done")
	}
	fmt.Println("group done")
}

func main() {
}
