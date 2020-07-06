// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalData(bytes)
//    bytes, err = topLevel.Marshal()

package show

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
	ID                int64      `json:"id"`                 
	Idstr             string     `json:"idstr"`              
	Class             int64      `json:"class"`              
	ScreenName        string     `json:"screen_name"`        
	Name              string     `json:"name"`               
	Province          string     `json:"province"`           
	City              string     `json:"city"`               
	Location          string     `json:"location"`           
	Description       string     `json:"description"`        
	URL               string     `json:"url"`                
	ProfileImageURL   string     `json:"profile_image_url"`  
	CoverImagePhone   string     `json:"cover_image_phone"`  
	ProfileURL        string     `json:"profile_url"`        
	Domain            string     `json:"domain"`             
	Weihao            string     `json:"weihao"`             
	Gender            string     `json:"gender"`             
	FollowersCount    int64      `json:"followers_count"`    
	FriendsCount      int64      `json:"friends_count"`      
	PagefriendsCount  int64      `json:"pagefriends_count"`  
	StatusesCount     int64      `json:"statuses_count"`     
	VideoStatusCount  int64      `json:"video_status_count"` 
	VideoPlayCount    int64      `json:"video_play_count"`   
	FavouritesCount   int64      `json:"favourites_count"`   
	CreatedAt         string     `json:"created_at"`         
	Following         bool       `json:"following"`          
	AllowAllActMsg    bool       `json:"allow_all_act_msg"`  
	GeoEnabled        bool       `json:"geo_enabled"`        
	Verified          bool       `json:"verified"`           
	VerifiedType      int64      `json:"verified_type"`      
	Remark            string     `json:"remark"`             
	Email             string     `json:"email"`              
	Insecurity        Insecurity `json:"insecurity"`         
	Status            Status     `json:"status"`             
	Ptype             int64      `json:"ptype"`              
	AllowAllComment   bool       `json:"allow_all_comment"`  
	AvatarLarge       string     `json:"avatar_large"`       
	AvatarHD          string     `json:"avatar_hd"`          
	VerifiedReason    string     `json:"verified_reason"`    
	VerifiedTrade     string     `json:"verified_trade"`     
	VerifiedReasonURL string     `json:"verified_reason_url"`
	VerifiedSource    string     `json:"verified_source"`    
	VerifiedSourceURL string     `json:"verified_source_url"`
	FollowMe          bool       `json:"follow_me"`          
	Like              bool       `json:"like"`               
	LikeMe            bool       `json:"like_me"`            
	OnlineStatus      int64      `json:"online_status"`      
	BIFollowersCount  int64      `json:"bi_followers_count"` 
	Lang              string     `json:"lang"`               
	Star              int64      `json:"star"`               
	Mbtype            int64      `json:"mbtype"`             
	Mbrank            int64      `json:"mbrank"`             
	BlockWord         int64      `json:"block_word"`         
	BlockApp          int64      `json:"block_app"`          
	CreditScore       int64      `json:"credit_score"`       
	UserAbility       int64      `json:"user_ability"`       
	Urank             int64      `json:"urank"`              
	StoryReadState    int64      `json:"story_read_state"`   
	VclubMember       int64      `json:"vclub_member"`       
	IsTeenager        int64      `json:"is_teenager"`        
	IsGuardian        int64      `json:"is_guardian"`        
	IsTeenagerList    int64      `json:"is_teenager_list"`   
	PCNew             int64      `json:"pc_new"`             
	SpecialFollow     bool       `json:"special_follow"`     
	PlanetVideo       int64      `json:"planet_video"`       
	VideoMark         int64      `json:"video_mark"`         
	LiveStatus        int64      `json:"live_status"`        
}

type Insecurity struct {
	SexualContent bool `json:"sexual_content"`
}

type Status struct {
	Visible                  Visible           `json:"visible"`                   
	CreatedAt                string            `json:"created_at"`                
	ID                       int64             `json:"id"`                        
	Idstr                    string            `json:"idstr"`                     
	Mid                      string            `json:"mid"`                       
	CanEdit                  bool              `json:"can_edit"`                  
	ShowAdditionalIndication int64             `json:"show_additional_indication"`
	Text                     string            `json:"text"`                      
	SourceAllowclick         int64             `json:"source_allowclick"`         
	SourceType               int64             `json:"source_type"`               
	Source                   string            `json:"source"`                    
	Appid                    int64             `json:"appid"`                     
	Favorited                bool              `json:"favorited"`                 
	Truncated                bool              `json:"truncated"`                 
	InReplyToStatusID        string            `json:"in_reply_to_status_id"`     
	InReplyToUserID          string            `json:"in_reply_to_user_id"`       
	InReplyToScreenName      string            `json:"in_reply_to_screen_name"`   
	PicIDS                   []interface{}     `json:"pic_ids"`                   
	PicTypes                 string            `json:"pic_types"`                 
	Geo                      interface{}       `json:"geo"`                       
	IsPaid                   bool              `json:"is_paid"`                   
	MblogVipType             int64             `json:"mblog_vip_type"`            
	Annotations              []Annotation      `json:"annotations"`               
	RepostsCount             int64             `json:"reposts_count"`             
	CommentsCount            int64             `json:"comments_count"`            
	AttitudesCount           int64             `json:"attitudes_count"`           
	PendingApprovalCount     int64             `json:"pending_approval_count"`    
	IsLongText               bool              `json:"isLongText"`                
	RewardExhibitionType     int64             `json:"reward_exhibition_type"`    
	HideFlag                 int64             `json:"hide_flag"`                 
	Mlevel                   int64             `json:"mlevel"`                    
	BizFeature               int64             `json:"biz_feature"`               
	HasActionTypeCard        int64             `json:"hasActionTypeCard"`         
	DarwinTags               []interface{}     `json:"darwin_tags"`               
	HotWeiboTags             []interface{}     `json:"hot_weibo_tags"`            
	TextTagTips              []interface{}     `json:"text_tag_tips"`             
	Mblogtype                int64             `json:"mblogtype"`                 
	Rid                      string            `json:"rid"`                       
	UserType                 int64             `json:"userType"`                  
	MoreInfoType             int64             `json:"more_info_type"`            
	PositiveRecomFlag        int64             `json:"positive_recom_flag"`       
	ContentAuth              int64             `json:"content_auth"`              
	GIFIDS                   string            `json:"gif_ids"`                   
	IsShowBulletin           int64             `json:"is_show_bulletin"`          
	CommentManageInfo        CommentManageInfo `json:"comment_manage_info"`       
	RepostType               int64             `json:"repost_type"`               
	PicNum                   int64             `json:"pic_num"`                   
}

type Annotation struct {
	MAPIRequest bool `json:"mapi_request"`
}

type CommentManageInfo struct {
	CommentPermissionType int64 `json:"comment_permission_type"`
	ApprovalCommentType   int64 `json:"approval_comment_type"`  
}

type Visible struct {
	Type   int64 `json:"type"`   
	ListID int64 `json:"list_id"`
}
