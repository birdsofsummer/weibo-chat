package types 

import "encoding/json"


func UnmarshalMsg(data []byte) (Msg, error) {
	var r Msg
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Msg) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


/*
炸群

{
  "result": false,
  "request": "/groupchat/query_messages.json",
  "error_code": 21202,
  "error": "不允许此项操作，有疑问请咨询 @粉丝群"
}
*/

type MsgError struct {
	Result    bool   `json:"result"`    
	Request   string `json:"request"`   
	ErrorCode int64  `json:"error_code"`
	Error     string `json:"error"`     
}


type Msg struct {
	LastReadMid     int64         `json:"last_read_mid"`   
	SignificantMsgs []interface{} `json:"significant_msgs"`
	GroupTips       []GroupTip    `json:"group_tips"`      
	Messages        []Messages    `json:"messages"`        
	Result          bool          `json:"result"`          
	Ts              int64         `json:"ts"`              
}

type GroupTip struct {
	GroupID         int64  `json:"group_id"`        
	ContentTemplate string `json:"content_template"`
	Time            int64  `json:"time"`            
	Type            int64  `json:"type"`            
	TipsID          int64  `json:"tips_id"`         
	ContentData     string `json:"content_data"`    
	Content         string `json:"content"`         
}

type Messages struct {
	Annotations  Annotations `json:"annotations"`           
	Appid        int64       `json:"appid"`                 
	Content      string      `json:"content"`               
	FaithStatus  *int64      `json:"faith_status,omitempty"`
	Fids         []int64     `json:"fids"`                  
	FromUid      int64       `json:"from_uid"`              
	FromUser     FromUser    `json:"from_user"`             
	Gid          int64       `json:"gid"`                   
	ID           int64       `json:"id"`                    
	Mblogid      string      `json:"mblogid"`               
	MediaType    int64       `json:"media_type"`            
	RecallStatus int64       `json:"recall_status"`         
	Time         int64       `json:"time"`                  
	Type         int64       `json:"type"`                  
	URLObjects   []URLObject `json:"url_objects"`           
}

type Annotations struct {
	Gdid         string `json:"gdid"`         
	IncludeVideo int64  `json:"include_video"`
	SendFrom     string `json:"send_from"`    
	Spr          string `json:"spr"`          
}

type FromUser struct {
	AvatarLarge     string `json:"avatar_large"`     
	FollowersCount  int64  `json:"followers_count"`  
	FriendsCount    int64  `json:"friends_count"`    
	ID              int64  `json:"id"`               
	Level           int64  `json:"level"`            
	ProfileImageURL string `json:"profile_image_url"`
	ProfileURL      string `json:"profile_url"`      
	ScreenName      string `json:"screen_name"`      
	Verified        bool   `json:"verified"`         
}

type URLObject struct {
	AssoLikeCount         int64  `json:"asso_like_count"`         
	CardInfoUnIntegrity   bool   `json:"card_info_un_integrity"`  
	FollowerCount         int64  `json:"follower_count"`          
	Info                  Info   `json:"info"`                    
	IsActionType          bool   `json:"isActionType"`            
	LikeCount             int64  `json:"like_count"`              
	ObjectID              string `json:"object_id"`               
	PageID                string `json:"page_id"`                 
	SearchTopicCount      int64  `json:"search_topic_count"`      
	SearchTopicReadCount  int64  `json:"search_topic_read_count"` 
	Status                Status `json:"status"`                  
	SuperTopicPhotoCount  int64  `json:"super_topic_photo_count"` 
	SuperTopicStatusCount int64  `json:"super_topic_status_count"`
	URLOri                string `json:"url_ori"`                 
}

type Info struct {
	Description  string `json:"description"`  
	EXTStatus    int64  `json:"ext_status"`   
	LastModified int64  `json:"last_modified"`
	Result       bool   `json:"result"`       
	Title        string `json:"title"`        
	Transcode    int64  `json:"transcode"`    
	Type         int64  `json:"type"`         
	URLLong      string `json:"url_long"`     
	URLShort     string `json:"url_short"`    
}

type Status struct {
	Annotations              []Annotation          `json:"annotations"`               
	Appid                    int64                 `json:"appid"`                     
	AttitudesCount           int64                 `json:"attitudes_count"`           
	BizFeature               int64                 `json:"biz_feature"`               
	BizIDS                   []int64               `json:"biz_ids"`                   
	CanEdit                  bool                  `json:"can_edit"`                  
	Cardid                   string                `json:"cardid"`                    
	CommentManageInfo        CommentManageInfo     `json:"comment_manage_info"`       
	CommentsCount            int64                 `json:"comments_count"`            
	ContentAuth              int64                 `json:"content_auth"`              
	CreatedAt                string                `json:"created_at"`                
	DarwinTags               []DarwinTag           `json:"darwin_tags"`               
	ExternSafe               int64                 `json:"extern_safe"`               
	Favorited                bool                  `json:"favorited"`                 
	Geo                      Geo                   `json:"geo"`                       
	GIFIDS                   string                `json:"gif_ids"`                   
	HasActionTypeCard        int64                 `json:"hasActionTypeCard"`         
	HideFlag                 int64                 `json:"hide_flag"`                 
	HotWeiboTags             []interface{}         `json:"hot_weibo_tags"`            
	ID                       int64                 `json:"id"`                        
	Idstr                    string                `json:"idstr"`                     
	InReplyToScreenName      string                `json:"in_reply_to_screen_name"`   
	InReplyToStatusID        string                `json:"in_reply_to_status_id"`     
	InReplyToUserID          string                `json:"in_reply_to_user_id"`       
	IsLongText               bool                  `json:"isLongText"`                
	IsPaid                   bool                  `json:"is_paid"`                   
	IsShowBulletin           int64                 `json:"is_show_bulletin"`          
	MblogVipType             int64                 `json:"mblog_vip_type"`            
	Mblogtype                int64                 `json:"mblogtype"`                 
	Mid                      string                `json:"mid"`                       
	Mlevel                   int64                 `json:"mlevel"`                    
	MoreInfoType             int64                 `json:"more_info_type"`            
	NumberDisplayStrategy    NumberDisplayStrategy `json:"number_display_strategy"`   
	PendingApprovalCount     int64                 `json:"pending_approval_count"`    
	PicIDS                   []string              `json:"pic_ids"`                   
	PicNum                   int64                 `json:"pic_num"`                   
	PicTypes                 string                `json:"pic_types"`                 
	PositiveRecomFlag        int64                 `json:"positive_recom_flag"`       
	RepostsCount             int64                 `json:"reposts_count"`             
	RewardExhibitionType     int64                 `json:"reward_exhibition_type"`    
	Rid                      string                `json:"rid"`                       
	ShowAdditionalIndication int64                 `json:"show_additional_indication"`
	Source                   string                `json:"source"`                    
	SourceAllowclick         int64                 `json:"source_allowclick"`         
	SourceType               int64                 `json:"source_type"`               
	Text                     string                `json:"text"`                      
	TextLength               int64                 `json:"textLength"`                
	TextTagTips              []interface{}         `json:"text_tag_tips"`             
	Truncated                bool                  `json:"truncated"`                 
	User                     User                  `json:"user"`                      
	UserType                 int64                 `json:"userType"`                  
	Visible                  Visible               `json:"visible"`                   
}

type Annotation struct {
	OID  string `json:"oid"` 
	Type string `json:"type"`
}

type CommentManageInfo struct {
	ApprovalCommentType   int64 `json:"approval_comment_type"`  
	CommentPermissionType int64 `json:"comment_permission_type"`
}

type DarwinTag struct {
	BdObjectType  string      `json:"bd_object_type"`
	DisplayName   string      `json:"display_name"`  
	EnterpriseUid interface{} `json:"enterprise_uid"`
	ObjectID      string      `json:"object_id"`     
	ObjectType    string      `json:"object_type"`   
}

type Geo struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`       
}

type NumberDisplayStrategy struct {
	ApplyScenarioFlag    int64  `json:"apply_scenario_flag"`    
	DisplayText          string `json:"display_text"`           
	DisplayTextMinNumber int64  `json:"display_text_min_number"`
}

type User struct {
	AllowAllActMsg    bool             `json:"allow_all_act_msg"`  
	AllowAllComment   bool             `json:"allow_all_comment"`  
	AvatarHD          string           `json:"avatar_hd"`          
	AvatarLarge       string           `json:"avatar_large"`       
	Badge             map[string]int64 `json:"badge"`              
	BadgeTop          string           `json:"badge_top"`          
	BIFollowersCount  int64            `json:"bi_followers_count"` 
	BlockApp          int64            `json:"block_app"`          
	BlockWord         int64            `json:"block_word"`         
	Cardid            string           `json:"cardid"`             
	City              string           `json:"city"`               
	Class             int64            `json:"class"`              
	CoverImagePhone   string           `json:"cover_image_phone"`  
	CreatedAt         string           `json:"created_at"`         
	CreditScore       int64            `json:"credit_score"`       
	Description       string           `json:"description"`        
	Domain            string           `json:"domain"`             
	Extend            Extend           `json:"extend"`             
	FavouritesCount   int64            `json:"favourites_count"`   
	FollowMe          bool             `json:"follow_me"`          
	FollowersCount    int64            `json:"followers_count"`    
	Following         bool             `json:"following"`          
	FriendsCount      int64            `json:"friends_count"`      
	Gender            string           `json:"gender"`             
	GeoEnabled        bool             `json:"geo_enabled"`        
	HasAbilityTag     int64            `json:"has_ability_tag"`    
	ID                int64            `json:"id"`                 
	Idstr             string           `json:"idstr"`              
	Insecurity        Insecurity       `json:"insecurity"`         
	IsGuardian        int64            `json:"is_guardian"`        
	IsTeenager        int64            `json:"is_teenager"`        
	IsTeenagerList    int64            `json:"is_teenager_list"`   
	Lang              string           `json:"lang"`               
	Level             int64            `json:"level"`              
	Like              bool             `json:"like"`               
	LikeMe            bool             `json:"like_me"`            
	Location          string           `json:"location"`           
	Mbrank            int64            `json:"mbrank"`             
	Mbtype            int64            `json:"mbtype"`             
	Name              string           `json:"name"`               
	OnlineStatus      int64            `json:"online_status"`      
	PagefriendsCount  int64            `json:"pagefriends_count"`  
	ProfileImageURL   string           `json:"profile_image_url"`  
	ProfileURL        string           `json:"profile_url"`        
	Province          string           `json:"province"`           
	Ptype             int64            `json:"ptype"`              
	Remark            string           `json:"remark"`             
	ScreenName        string           `json:"screen_name"`        
	SpecialFollow     bool             `json:"special_follow"`     
	Star              int64            `json:"star"`               
	StatusesCount     int64            `json:"statuses_count"`     
	StoryReadState    int64            `json:"story_read_state"`   
	TabManage         string           `json:"tab_manage"`         
	Type              int64            `json:"type"`               
	Ulevel            int64            `json:"ulevel"`             
	Urank             int64            `json:"urank"`              
	URL               string           `json:"url"`                
	UserAbility       int64            `json:"user_ability"`       
	UserLimit         int64            `json:"user_limit"`         
	VclubMember       int64            `json:"vclub_member"`       
	Verified          bool             `json:"verified"`           
	VerifiedReason    string           `json:"verified_reason"`    
	VerifiedReasonURL string           `json:"verified_reason_url"`
	VerifiedSource    string           `json:"verified_source"`    
	VerifiedSourceURL string           `json:"verified_source_url"`
	VerifiedTrade     string           `json:"verified_trade"`     
	VerifiedType      int64            `json:"verified_type"`      
	VideoStatusCount  int64            `json:"video_status_count"` 
	Weihao            string           `json:"weihao"`             
}

type Extend struct {
	Mbprivilege string  `json:"mbprivilege"`
	Privacy     Privacy `json:"privacy"`    
}

type Privacy struct {
	Mobile int64 `json:"mobile"`
}

type Insecurity struct {
	SexualContent bool `json:"sexual_content"`
}

type Visible struct {
	ListID int64 `json:"list_id"`
	Type   int64 `json:"type"`   
}
