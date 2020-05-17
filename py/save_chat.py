import requests
import time
import json

now=lambda: int(time.time()*1e3)

c="SINAGLOBAL=8372698832750.411.1586951892828; ULV=1588989175725:1:1:1:8372698832750.411.1586951892828:; SUHB=0cSXyS3AKtTlEy; ALF=1620905033; wvr=6; UOR=,,login.sina.com.cn; webim_unReadCount=%7B%22time%22%3A1589374104726%2C%22dm_pub_total%22%3A1%2C%22chat_group_client%22%3A711%2C%22chat_group_notice%22%3A0%2C%22allcountNum%22%3A712%2C%22msgbox%22%3A0%7D; SCF=Ahf_H6nQmmH8A6OM5pQpKM8Fsj0a5iRtqRerzo_0YjjPFc-oe-79LOsaZCPWbzdAi4CrKqtZgefEY2riYXnV3Gc.; login_sid_t=ddf60580ed993c3f54e2923e872f4230; cross_origin_proto=SSL; _s_tentry=-; Apache=8372698832750.411.1586951892828; SSOLoginState=1587018632; ULOGIN_IMG=15873742810023; JSESSIONID=D48B9CB449C93F59214596FD918C554A; SUB=_2A25zv6ipDeRhGeFO7FQZ8yzMyziIHXVQzJ1hrDV8PUNbmtAKLUjbkW9NQVoZcDAleYyJcu7WuHleT5xNHVD1fY6R; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WFj3InpCz3yDbL2y0aBj1oL5JpX5KMhUgL.FoM7S0qRe0z7ehB2dJLoIE.LxKqL1-eL1hnLxKqL1KnL12-LxKBLB.eLB.eLxK.L1-zLB-8_dcvVeBtt"

h={
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
        "Accept": "application/json, text/plain, */*",
        "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
        'Sec-Fetch-Dest':'empty',
        'Sec-Fetch-Mode':'cors',
        'Sec-Fetch-Site':'same-origin',
        "Referer": "https://api.weibo.com/chat/",
        "Cookie":c,
}

FILE="json/1001.json"

s=requests.session()
s.headers.update(h)

def get(max_mid=0,id=4429200364208496):
    u="https://api.weibo.com/webim/groupchat/query_messages.json"
    d={
            'convert_emoji' : '1',
            'query_sender' : '1',
            'count' : '20',
            "id":id,
            'source' : '209678993',
            "max_mid":max_mid,
            't' : now(),
    }
    r=s.get(u,params=d,headers=h)
    return r.json()

def test():
    z=get()
    d= [(
            x["time"],
            x["from_uid"],
            x["from_user"]["screen_name"],
            x['content'],
        ) for x in z['messages']]
    print(d)
    return z

def start(i=0,n1=0,r1=[],max=0,id=4429200364208496):
    print("wait...#",i,n1,"to",max,n1)
    if max > 0 and n1 >= max: return r1
    z=get(n1,id)
    m=z['messages']
    if len(m)>0:
        n2=m[0]['id']-1
        r1.extend(m)
        return start(i+1,n2,r1,max,id)
    else:
        print(i,n1,"done")
    return r1

def get_group():
    dd=start()
    ds=json.dumps(dd,
            indent=4,
            sort_keys=True,
            ensure_ascii=False
    )
    with open(FILE,"r+") as f:
        f.write(ds)

def append_group():
    s1=""
    try:
        with open(FILE,"r+") as f:
            z=f.read()
            z1=json.loads(z)
            print(len(z1))
            id=4429200364208496
            max=z1[0]["id"]
            print("append from ...",max)
            r1=start(0,0,[],max,id)
            r1.extend(z1)
            d1=json.dumps(r1,
                indent=4,
                sort_keys=True,
                ensure_ascii=False
            )
            s1=d1
        with open(FILE,"w+") as f:
            f.write(s1)
    except:
        print("start from 0")
        return get_group()
    else:
        print('append done')


#get_group()
while True:
    append_group()
    print('sleep 1 min')
    time.sleep(60)
