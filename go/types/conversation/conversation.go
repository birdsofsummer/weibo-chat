// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalData(bytes)
//    bytes, err = topLevel.Marshal()

package conversation

import "encoding/json"

func UnmarshalData(data []byte) (Data, error) {
	var r Data
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Data) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Data struct {
	DirectMessages []DirectMessage `json:"direct_messages"`
	NextCursor     int64           `json:"next_cursor"`    
	PreviousCursor int64           `json:"previous_cursor"`
	PrivateTime    int64           `json:"private_time"`   
	TotalNumber    int64           `json:"total_number"`   
	Following      bool            `json:"following"`      
	LastReadMid    int64           `json:"last_read_mid"`  
}

type DirectMessage struct {
	ID                  int64       `json:"id"`                   
	Idstr               string      `json:"idstr"`                
	CreatedAt           string      `json:"created_at"`           
	Text                string      `json:"text"`                 
	SysType             int64       `json:"sys_type"`             
	MsgStatus           int64       `json:"msg_status"`           
	SenderTrashUser     int64       `json:"sender_trash_user"`    
	SenderID            int64       `json:"sender_id"`            
	RecipientID         int64       `json:"recipient_id"`         
	SenderScreenName    Name        `json:"sender_screen_name"`   
	RecipientScreenName Name        `json:"recipient_screen_name"`
	Sender              Recipient   `json:"sender"`               
	Recipient           Recipient   `json:"recipient"`            
	Mid                 string      `json:"mid"`                  
	IsLargeDm           bool        `json:"isLargeDm"`            
	Source              string      `json:"source"`               
	StatusID            int64       `json:"status_id"`            
	Geo                 interface{} `json:"geo"`                  
	DmType              int64       `json:"dm_type"`              
	MediaType           int64       `json:"media_type"`           
	SubType             int64       `json:"sub_type"`             
	IP                  int64       `json:"ip"`                   
	Ipv6                interface{} `json:"ipv6"`                 
	BurnTime            int64       `json:"burn_time"`            
	EXTText             EXTText     `json:"ext_text"`             
	MatchKeyword        bool        `json:"matchKeyword"`         
	Topublic            bool        `json:"topublic"`             
	ReceiverBoxType     string      `json:"receiver_box_type"`    
	PushToMPS           bool        `json:"pushToMPS"`            
	OriImageID          int64       `json:"oriImageId"`           
	RecallStatus        int64       `json:"recall_status"`        
	URLObjects          []URLObject `json:"url_objects"`          
	Comment             *string     `json:"comment,omitempty"`    
}

type EXTText struct {
	SendFrom SendFrom `json:"send_from"`         
	Clientid *string  `json:"clientid,omitempty"`
}

type Recipient struct {
	ID                int64       `json:"id"`                         
	Idstr             string      `json:"idstr"`                      
	Class             int64       `json:"class"`                      
	ScreenName        Name        `json:"screen_name"`                
	Name              Name        `json:"name"`                       
	Province          string      `json:"province"`                   
	City              string      `json:"city"`                       
	Location          string      `json:"location"`                   
	Description       Description `json:"description"`                
	URL               string      `json:"url"`                        
	ProfileImageURL   string      `json:"profile_image_url"`          
	CoverImagePhone   *string     `json:"cover_image_phone,omitempty"`
	ProfileURL        ProfileURL  `json:"profile_url"`                
	Domain            string      `json:"domain"`                     
	Weihao            string      `json:"weihao"`                     
	Gender            Gender      `json:"gender"`                     
	FollowersCount    int64       `json:"followers_count"`            
	FriendsCount      int64       `json:"friends_count"`              
	PagefriendsCount  int64       `json:"pagefriends_count"`          
	StatusesCount     int64       `json:"statuses_count"`             
	VideoStatusCount  int64       `json:"video_status_count"`         
	FavouritesCount   int64       `json:"favourites_count"`           
	CreatedAt         CreatedAt   `json:"created_at"`                 
	Following         bool        `json:"following"`                  
	AllowAllActMsg    bool        `json:"allow_all_act_msg"`          
	GeoEnabled        bool        `json:"geo_enabled"`                
	Verified          bool        `json:"verified"`                   
	VerifiedType      int64       `json:"verified_type"`              
	Remark            string      `json:"remark"`                     
	Insecurity        Insecurity  `json:"insecurity"`                 
	Ptype             int64       `json:"ptype"`                      
	AllowAllComment   bool        `json:"allow_all_comment"`          
	AvatarLarge       string      `json:"avatar_large"`               
	AvatarHD          string      `json:"avatar_hd"`                  
	VerifiedReason    string      `json:"verified_reason"`            
	VerifiedTrade     string      `json:"verified_trade"`             
	VerifiedReasonURL string      `json:"verified_reason_url"`        
	VerifiedSource    string      `json:"verified_source"`            
	VerifiedSourceURL string      `json:"verified_source_url"`        
	FollowMe          bool        `json:"follow_me"`                  
	Like              bool        `json:"like"`                       
	LikeMe            bool        `json:"like_me"`                    
	OnlineStatus      int64       `json:"online_status"`              
	BIFollowersCount  int64       `json:"bi_followers_count"`         
	Lang              Lang        `json:"lang"`                       
	Star              int64       `json:"star"`                       
	Mbtype            int64       `json:"mbtype"`                     
	Mbrank            int64       `json:"mbrank"`                     
	BlockWord         int64       `json:"block_word"`                 
	BlockApp          int64       `json:"block_app"`                  
	CreditScore       int64       `json:"credit_score"`               
	UserAbility       int64       `json:"user_ability"`               
	Urank             int64       `json:"urank"`                      
	StoryReadState    int64       `json:"story_read_state"`           
	VclubMember       int64       `json:"vclub_member"`               
	IsTeenager        int64       `json:"is_teenager"`                
	IsGuardian        int64       `json:"is_guardian"`                
	IsTeenagerList    int64       `json:"is_teenager_list"`           
}

type Insecurity struct {
	SexualContent bool `json:"sexual_content"`
}

type URLObject struct {
	URLOri                string          `json:"url_ori"`                 
	ObjectID              string          `json:"object_id"`               
	Info                  Info            `json:"info"`                    
	Object                URLObjectObject `json:"object"`                  
	LikeCount             int64           `json:"like_count"`              
	IsActionType          bool            `json:"isActionType"`            
	FollowerCount         int64           `json:"follower_count"`          
	AssoLikeCount         int64           `json:"asso_like_count"`         
	CardInfoUnIntegrity   bool            `json:"card_info_un_integrity"`  
	SuperTopicStatusCount int64           `json:"super_topic_status_count"`
	SuperTopicPhotoCount  int64           `json:"super_topic_photo_count"` 
	SearchTopicCount      int64           `json:"search_topic_count"`      
	SearchTopicReadCount  int64           `json:"search_topic_read_count"` 
}

type Info struct {
	URLShort     string `json:"url_short"`    
	URLLong      string `json:"url_long"`     
	Type         int64  `json:"type"`         
	Result       bool   `json:"result"`       
	Title        string `json:"title"`        
	Description  string `json:"description"`  
	LastModified int64  `json:"last_modified"`
	Transcode    int64  `json:"transcode"`    
	EXTStatus    int64  `json:"ext_status"`   
}

type URLObjectObject struct {
	ObjectID       string       `json:"object_id"`       
	Containerid    string       `json:"containerid"`     
	ObjectDomainID string       `json:"object_domain_id"`
	ObjectType     string       `json:"object_type"`     
	SafeStatus     int64        `json:"safe_status"`     
	ShowStatus     string       `json:"show_status"`     
	ActStatus      string       `json:"act_status"`      
	LastModified   string       `json:"last_modified"`   
	Timestamp      int64        `json:"timestamp"`       
	UUID           int64        `json:"uuid"`            
	Uuidstr        string       `json:"uuidstr"`         
	ActivateStatus string       `json:"activate_status"` 
	Object         ObjectObject `json:"object"`          
}

type ObjectObject struct {
	Summary      string       `json:"summary"`               
	Image        Image        `json:"image"`                 
	URLShort     *string      `json:"url_short,omitempty"`   
	Keywords     *string      `json:"keywords,omitempty"`    
	AreaClass    *string      `json:"area_class,omitempty"`  
	ObjectType   string       `json:"object_type"`           
	Description  *string      `json:"description,omitempty"` 
	Generator    *string      `json:"generator,omitempty"`   
	EffectURL    *string      `json:"effect_url,omitempty"`  
	ImageOrigin  *ImageOrigin `json:"image_origin,omitempty"`
	DisplayName  string       `json:"display_name"`          
	Priority     *int64       `json:"priority,omitempty"`    
	Imglist      *bool        `json:"imglist,omitempty"`     
	URL          string       `json:"url"`                   
	Score        *string      `json:"score,omitempty"`       
	From         *string      `json:"from,omitempty"`        
	IsMobile     *bool        `json:"isMobile,omitempty"`    
	ID           string       `json:"id"`                    
	ClassLevel2  *string      `json:"class_level2,omitempty"`
	ClassLevel1  *string      `json:"class_level1,omitempty"`
	Ge           *bool        `json:"ge,omitempty"`          
	Md5          *string      `json:"md5,omitempty"`         
	Biz          Biz          `json:"biz"`                   
	TargetURL    string       `json:"target_url"`            
	Gid          *int64       `json:"gid,omitempty"`         
	Owner        *int64       `json:"owner,omitempty"`       
	OwnerName    *string      `json:"owner_name,omitempty"`  
	GroupName    *string      `json:"group_name,omitempty"`  
	MemberCount  *int64       `json:"member_count,omitempty"`
	GroupMembers []int64      `json:"group_members"`         
}

type Biz struct {
	BizID       string `json:"biz_id"`     
	Containerid string `json:"containerid"`
}

type Image struct {
	URL      string  `json:"url"`                
	RoundURL *string `json:"round_url,omitempty"`
}

type ImageOrigin struct {
}

type SendFrom string
const (
	Msg SendFrom = "msg"
	Share SendFrom = "share"
	Webchat SendFrom = "webchat"
)

type CreatedAt string
const (
	SatAPR1314100708002019 CreatedAt = "Sat Apr 13 14:10:07 +0800 2019"
	SunFeb0914555008002020 CreatedAt = "Sun Feb 09 14:55:50 +0800 2020"
)

type Description string
const (
	Empty Description = ""
	禾舀花香里说丰年 Description = "禾舀花香里说丰年"
)

type Gender string
const (
	F Gender = "f"
	M Gender = "m"
)

type Lang string
const (
	ZhCN Lang = "zh-cn"
)

type Name string
const (
	月明似雪Moon1 Name = "月明似雪moon1"
	空气动力学研究员 Name = "空气动力学研究员"
)

type ProfileURL string
const (
	U7076832004 ProfileURL = "u/7076832004"
	U7388859010 ProfileURL = "u/7388859010"
)
