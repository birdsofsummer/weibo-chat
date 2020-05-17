a=require('utility')
fs=require("fs")



get_url=(id="")=>"https://weibo.com/u/"+id
format=(t=0)=>new Date(t).toLocaleString()
to_csv1=(d=[])=>{
    let k=Object.keys(d[0])
    h=k.map(x=>`"${x}"`).join(",")
    t=d.map(x=>k.map(y=>x[y]||"").map(j=>`"${j}"`).join(",")).join("\n")
   let c=[h,t].join("\n")
   return c
}

init=async ()=>{
    d=await a.readJSON('json/1001.json')
    d1=to_csv1(d)
    d1=d.reverse().slice(1).map(x=>({
            //...x.from_user,
            time:format(x.time*1e3),
            from_uid:get_url(x.from_uid),
            screen_name:x.from_user ? x.from_user.screen_name:"" ,
            content:x.content,
            file: !x.fids ? "" :x.fids.map(y=>'https://upload.api.weibo.com/2/mss/msget?fid='+y+'&source=209678993&imageType=origin').join('\n')
        }))
    s=to_csv1(d1)
    fs.writeFileSync('json/1.csv',s)
    console.log('done')
}

init()

