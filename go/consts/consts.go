package consts

import (
    "fmt"
    "github.com/BurntSushi/toml"
)


type Config struct{
    Cookie string `toml:"cookie"` 
	ID int64 `toml:"id"` 
}

var (
  HEADERS=map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
        "Accept": "application/json, text/plain, */*",
        "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
		"DNT":	"1",
        "Pragma": "no-cache",
		"Host":	"api.weibo.com",
        "Cache-Control": "no-cache",
        "Sec-Fetch-Dest":"empty",
        "Sec-Fetch-Mode":"cors",
        "Sec-Fetch-Site":"same-origin",
        "Referer": "https://api.weibo.com/chat/",
        "Cookie":"",
		//"Accept-Encoding": "gzip, deflate, br",
	}
)
func GetConfig() *Config{
	var config *Config=new (Config)
    _, err := toml.DecodeFile("config.toml",config)
    if err!=nil{
        fmt.Println(err)
    }
	return config
}

func GetHeaders() map[string]string{
	var config=GetConfig()
	//fmt.Println("zzzzzzzz",config)

	var h=map[string]string{}
	for k,v:=range HEADERS{
		h[k]=v
	}
	h["Cookie"]=config.Cookie
	//fmt.Println("zzzzzzzz",h)
	return h
}


