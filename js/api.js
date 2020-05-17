follow={
    "cleartrashfans":{url:"/aj/f/trash/cleartrashfans",method:"post"},
    "deltrashfans":{url:"/aj/f/trash/deltrashfans",method:"post"},
    "confirmfans":{url:"/aj/f/trash/confirmfans",method:"post"},
    "recommendfollow":{url:"/aj/f/recomafterfollow",method:"get"},  //...
    "closerecommend":{url:"/aj/f/closerecommend",method:"get"},
    "newuserguide":{url:"/aj/user/interest/newuserguide",method:"get"},
    "mayinterested":{url:"/aj/user/interest/list",method:"get"},
    "uninterested":{url:"/aj/user/interest/uninterested",method:"post"},
    "userCard":{url:"/aj/user/cardv5",method:"get"},
    "userCard2":{url:"//weibo.com/aj/user/newcard",method:"get",requestMode:"jsonp",varkey:"callback"},
    "userCard2_abroad":{url:"//www.weibo.com/aj/user/newcard",method:"get",requestMode:"jsonp",varkey:"callback"},
    "follow":{url:"/aj/f/followed",method:"post"},
    "unFollow":{url:"/aj/f/unfollow",method:"post"},
    "follow_register":{url:"/nguide/aj/relation/followed",method:"post"},
    "unfollow_register":{url:"/nguide/aj/relation/unfollow",method:"post"},
    "block":{url:"/aj/f/addblack",method:"post"},
    "unBlock":{url:"/aj/f/delblack",method:"post"},
    "removeFans":{url:"/aj/f/remove",method:"post"},
    "requestFollow":{url:"/ajax/relation/requestfollow",method:"post"},
    "questions":{url:"/aj/invite/attlimit",method:"get"},
    "answer":{url:"/aj/invite/att",method:"post"},
    "setRemark":{url:"/aj/f/remarkname",method:"post"},
    "recommendusers":{url:"/aj/f/recommendusers",method:"get"},
    "recommendAttUsers":{url:"/aj/f/worthfollowusers",method:"get"},
    "recommendPopularUsers":{url:"/aj/user/interest/recommendpopularusers",method:"get"},
    "mayinterestedweiqun":{url:"/aj/weiqun/getinterestedlist",method:"get"},
    "moreData":{url:"/aj/f/listuserdetail",method:"get"},
    "getInvite":{url:"/aj/invite/unread",method:"get"},
    "quiet_addUser":{url:"/aj/f/addwhisper",method:"post"},
    "quiet_removeUser":{url:"/aj/f/delwhisper",method:"post"},
    "quiet_know":{url:"/aj/tipsbar/closetipsbar",method:"post"},
    "groupUserList":{url:"/aj/f/group/getgroupmembers",method:"get"},
    "smart_sort":{url:"/aj/mblog/mblogcard",method:"post"},
    "groupSubmit":{url:"/aj/f/group/list",method:"get"},
    "wqList":{url:"/aj/proxy?api=http://recom.i.t.sina.com.cn/1/weiqun/weiqun_may_interest.php",method:"get"},
    "uninterestedWq":{url:"/aj/proxy?api=http://recom.i.t.sina.com.cn/1/weiqun/weiqun_uninterest.php",method:"get"},
    "inviteNeglect":{url:"/aj/invite/handleinvite",method:"post"},
    "checkNeglect":{url:"/aj/invite/shieldedlist",method:"post"},
    "inviteLift":{url:"/aj/invite/lift",method:"post"},
    "inviteAccept":{url:"/aj/invite/handleinvite",method:"post"},
    "searchByTel":{url:"/aj/relation/getuserbymobile",method:"post"},
    "inviteCloseTips":{url:"/aj/invite/closetips",method:"post"},
    "checkrelation":{url:"/aj/f/checkrelation",method:"post"},
    "addCloseFriend":{url:"/aj/f/createclosefriend",method:"post"},
    "removeClsFrd":{url:"/aj/f/removeclosefriend",method:"post"},
    "cfInviteUnread":{url:"/aj/invite/unread",method:"get"},
    "recommendCf":{url:"/aj/user/closefriend/recommend",method:"get"},
    "clearInvalidUsers":{url:"/aj/f/clearinvalidfriends",method:"post"},
    "unIstCf":{url:"/aj/user/closefriend/deny",method:"post"},
    "checkcloserelation":{url:"/aj/f/checkcloserelation",method:"post"},
    "closeunfollow":{url:"/aj/profile/closeunfollow",method:"post"},
    "fanslikemore":{url:"/aj/relation/fanslikemore",method:"get"},
    "getProfileInfo":{url:"/aj/relation/getprofileinfo",method:"get"},
    "interestlist":{url:"/aj/user/interest/profileinfo",method:"get"},
    "recommendGroupMember":{url:"/aj/user/group/recommend",method:"get"},
    "followGroup":{url:"/aj/f/group/followedgroup",method:"post"},
    "recommendWholeGroup":{url:"/aj/relation/rename",method:"post"},
    "recommendUserAdd":{url:"/aj/f/group/addrecommenduser",method:"post"},
    "recommendUserRemove":{url:"/aj/f/group/remove",method:"post"},
    "specialAttentionClose":{url:"/aj/limit/increment",method:"get"},
    "friendShield":{url:"/aj/mblog/friendsmblogshield",method:"post"},
    "friendRecover":{url:"/aj/mblog/friendsshieldrecover",method:"post"},
    "rightmod_getCloseFriendRecommend":{url:"/aj/f/recomfriendsmore",method:"get"},
    "unsubscribe":{url:"/aj/relation/unsubscribe",method:"post"},
    "batch":{url:"/aj/f/group/batch",method:"post"},
    "attention":{url:"/aj/relation/attention",method:"get"},
    "recommend":{url:"/aj/f/group/recommend",method:"get"},
    "interest_uninterested":{url:"/aj/relation/uninterested",method:"post"},
    "addSpeical":{url:"/aj/f/group/addspecial",method:"post"},
    "delSpecial":{url:"/aj/f/group/delspecial",method:"post"},
}

group={
        "add":{url:"/aj/f/group/add",method:"post"},
        "modify":{url:"/aj/relation/rename",method:"post"},
        "del":{url:"/aj/relation/delete",method:"post"},
        "info":{url:"/aj/f/group/getgroupinfo",method:"get"},
        "set":{url:"/aj3/attention/aj_group_update_v4.php",method:"post"},
        "batchSet":{url:"/aj3/attention/aj_group_batchupdate_v4.php",method:"post"},
        "list":{url:"/aj/f/group/list",method:"get"},
        "order":{url:"/aj/f/group/order",method:"post"},
        "listbygroup":{url:"/aj/f/attchoose",method:"post"},
        "infolist":{url:"/aj/f/attfilterlist",method:"get"},
        "orientGroup":{url:"/aj/v6/f/group/groupmembers",method:"get"},
        "recommendfollow1":{url:"/aj3/recommend/aj_addrecommend.php",method:"post"}, //...
        "groupupdate":{url:"/aj/relation/groupupdate",method:"post"},
        "unInterest":{url:"/aj/f/group/notinterest",method:"post"},
        "editdesc":{url:"/aj/f/group/editdesc",method:"post"},
        "update":{url:"/aj/f/group/update",method:"post"},
        "followgroup":{url:"/aj/f/group/followgroup",method:"post"},
        "getGroupDesc":{url:"/aj/f/group/getdesc",method:"get"},
        "closebubble":{url:"/aj/bubble/closebubble",method:"get"},
        "recomgroupuser":{url:"/aj/user/group/recommend",method:"get"},
}





chat_api=[
    ["get",'/webim/2/users/show.json'],
    ["get",'/webim/2/account/profile/basic.json'],
    ["get","/webim/query_remark.json"],
    ["get",'/webim/pic_infos.json'],

    ["get",'/webim/2/direct_messages/public/conversation.json'],
    ["get",'/webim/2/direct_messages/public/contacts.json'],

    ["get",'/webim/2/direct_messages/contacts.json?special_source=3'],
    ["get",'/webim/2/direct_messages/conversation.json'],
    ["get",'/webim/2/direct_messages/show_batch.json?is_complete=1&is_encoded=0'],


    ["get",'/webim/query_bilateral_friends.json'],
    ["get",'/webim/2/friendships/batch_exists.json'],
    ["post",'/webim/2/friendships/remark/update.json'],

    ["post",'/webim/2/direct_messages/destroy_batch.json'],
    ["post",'/webim/2/direct_messages/set_unread_count.json'],

    ["get",'/webim/2/direct_messages/block/check_batch.json'],
    ["post",'/webim/2/direct_messages/block_batch.json'],
    ["post",'/webim/2/direct_messages/unblock_batch.json'],

    ["post",'/webim/2/direct_messages/public/destroy_batch.json'],
    ["post",'/webim/2/direct_messages/public/set_unread_count.json'],

    ["post",'/webim/2/direct_messages/new.json'],
    ["get",'/webim/2/direct_messages/messageboxsearch.json'],


    ["post",'/webim/groupchat/clear_sys_unread.json'],
    ["post",'/webim/groupchat/clear_unread.json'],

    ["post",'/webim/2/notice_center/discussion_group/set_push_settings.json'],
    ["post",'/webim/groupchat/update_user_settings.json'],
    ["get",'/webim/2/notice_center/discussion_group/push_settings.json?call_from=mobile'],


    ["get",'/2/emotions.json'],
    ["get",'/webim/2/short_url/info.json'],

    ["post",'/webim/groupchat/create.json',],
    ["post",'/webim/groupchat/update.json',],
    ["get",'/webim/groupchat/query.json',],
    ["get",'/webim/groupchat/query_nick.json',],
    ["get",'/webim/query_group.json?is_pc=1&query_member=1&sort_by_jp=1&query_member_count=5000',],
    ["get",'/webim/groupchat/query_messages.json?convert_emoji=1&query_sender=1',],
    ["get",'/webim/groupchat/query_message.json',],
    ["get",'/webim/groupchat/batch_query_messages.json',],
    ["get",'/webim/groupchat/query_join_groups.json'],
    ["get",'/webim/2/direct_messages/messageboxsearch.json?types=contact',],
    ["post",'/webim/groupchat/kick.json',],
    ["post",'/webim/groupchat/join.json',],
    ["post",'/webim/groupchat/send_message.json',],
    ["post",'/webim/groupchat/delete_message.json',],
    ["post",'/webim/2/direct_messages/delete_top_contact.json',],
    ["post",'/webim/groupchat/exit.json',],
    ["post",'/webim/2/direct_messages/delete_top_contact.json',],
    ["post",'/webim/groupchat/send_bulletin.json',],
    ["post",'/webim/groupchat/delete_user_bulletin.json',],
    ["post",'/webim/groupchat/delete_bulletin.json',],
    ["post",'/webim/fileplatform/init.json',],
    ["get",'/webim/webim_nas.json?returntype=json&v=1.1',],
    ["get",'/webim/2/direct_messages/query_receiver_setting.json',],
    ["get",'/webim/2/direct_messages/public/query_remind_type.json',],
    ["get",'/webim/2/direct_messages/recent_messages.json',],
    ["post",'/webim/report.json',],
    ["post",'/webim/2/multimedia/init.json',],
    ["get",'/webim/2/mss/query_video_trans.json',],
    ["post",'/webim/groupchat/recall.json',],
    ["post",'/webim/2/direct_messages/recall.json',],
    ["post",'/webim/groupchat/delete_message.json',],
    ["post",'/webim/2/direct_messages/destroy.json',],
    ["post",'/webim/2/mss/repost.json',],
    ["post",'/webim/groupchat/apply_join.json',],
    ["post",'/webim/follow_affi.json',],
    ["get",'/webim/groupchat/is_members.json',],
    ["get",'/webim/friendships/vipclub/relation/check.json',],
    ["post",'/webim/groupchat/add_admin.json',],
    ["post",'/webim/groupchat/delete_admin.json',],
    ["get",'/webim/groupchat/set_group_owner.json',],
    ["post",'/webim/groupchat/apply_update_to_affi.json',],
    ["get",'/webim/query_group_affi.json',],
    ["get",'/webim/groupchat/query_sys_messages.json',],
    ["post",'/webim/2/notice_center/outer/set_push_settings.json',],
    ["get",'/webim/2/notice_center/outer/push_settings.json',],
    ["post",'/webim/groupchat/delete_sys_message.json',],
    ["post",'/webim/groupchat/apply_check.json',],
    ["get",'//login.sina.com.cn/sso/login.php?entry=weibo&returntype=TEXT&crossdomain=1&cdult=3&domain=weibo.com&savestate=30&source=209678993',],
]





cors=[
    [ '//login.sina.com.cn/sso/qrcode/image',
     {
                        entry: 'weibo',
                        size: 180,
                        source: 209678993
     }
    ],
    ['//upload.api.weibo.com/2/mss/meta_query.json?source=209678993']
    ['//login.sina.com.cn/sso/qrcode/check?entry=weibo&source=209678993'],
    ['//passport.weibo.com/wbsso/logout',
     {
        r: 'https://weibo.com',
        returntype: 1
     },
    ],
    [ '//weibo.com/aj/card/show?source=209678993&_t=3' ],
    ['//api.weibo.com/webim/webim_nas.json',
    {
                    source: 209678993,
                    returntype: 'json',
                    v: '1.1'
    }
    ],

    [
    '//rm.api.weibo.com/2/remind/push_count.json',
    {
        source: 209678993,
        with_dm_group: !0,
        with_chat_group: !0,
        with_dm_unread: 1,
        with_settings: 1,
        sign_val: 'newChat'
    }
    ],


]


qs=(...a)=>{
    let s=new URLSearchParams()
    for (let i of a){
        for (let j in i){
            s.append(j,i[j]) //不覆盖
           // s.set(j,i[j])    //覆盖
        }
    }
    return s
}


qs1=(u="",...a)=>{
   let u1=new URL(u.toString())
    for (d of a){
        for (k in d){
            u1.searchParams.append(k,d[k])
        }
    }
    return u1
}



/////// chat api
// 21301 == t.data.error_code ? '登录超时或失效，请重新登录'

get=(u="/",d={})=>{
    let q={
        source: '209678993',
        t: (new Date).getTime()
    }
    let u1=qs1(u,{...d,...q})
    let o={
        headers:{
            'Content-Type': 'application/x-www-form-urlencoded',
        }
    }
     return fetch(u1,o)
}
post=(u="/",d={})=>{
    let q={
        source: '209678993',
        t: (new Date).getTime()
    }
    let o={
        headers:{
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body:qs({...d,...q}),
    }
    return fetch(u,o)
}





window.Global=
{
	"isMac": false,
	"isWebkit": false,
	"isChrome": false,
	"isMozilla": true,
	"isOpera": false,
	"isIE": false,
	"isEdge": false,
	"isSafari": false,
	"toChatUid": null,
	"isLogouting": false,
	"browserVersion": "77",
	"sid": "1587601755526pwba1ix609yulyi", ////
	"online": true
}

 {
  'data' => '{"sid":"1587601755526pwba1ix609yulyi","type":3,"client_id":"lupkvcqc1c6f0ehb38snqec1mplqc"}',
  'source' => '209678993' }
