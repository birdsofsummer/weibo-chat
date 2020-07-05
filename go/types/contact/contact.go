// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalTopLevel(bytes)
//    bytes, err = topLevel.Marshal()

package  contact

import "bytes"
import "errors"
import "encoding/json"

func UnmarshalContact(data []byte) (ContactData, error) {
	var r ContactData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContactData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContactData struct {
	PublicMessgeRemindTab int64            `json:"public_messge_remind_tab"`
	PreviousCursor        int64            `json:"previous_cursor"`         
	TotalNumber           int64            `json:"totalNumber"`             
	IsBlockAllPublic      int64            `json:"is_block_all_public"`     
	NextCursor            int64            `json:"next_cursor"`             
	BackupService         int64            `json:"backup_service"`          
	ShowClearTrashButton  int64            `json:"show_clear_trash_button"` 
	Contacts              []ContactElement `json:"contacts"`                
}

type ContactElement struct {
	UnreadCount     int64            `json:"unread_count"`      
	Message         Message          `json:"message"`           
	User            User             `json:"user"`              
	Bulletin        *Bulletin        `json:"bulletin,omitempty"`
	SignificantMsgs []SignificantMsg `json:"significant_msgs"`  
}

type Bulletin struct {
	KeepTime   int64  `json:"keep_time"`  
	Scheme     string `json:"scheme"`     
	GroupID    int64  `json:"group_id"`   
	Time       int64  `json:"time"`       
	SenderUid  int64  `json:"sender_uid"` 
	BulletinID int64  `json:"bulletin_id"`
	Content    string `json:"content"`    
}

type Message struct {
	Idstr                *string              `json:"idstr,omitempty"`                  
	FaithStatus          *int64               `json:"faith_status,omitempty"`           
	CreatedAt            *string              `json:"created_at,omitempty"`             
	Mid                  *string              `json:"mid,omitempty"`                    
	SysType              *int64               `json:"sys_type,omitempty"`               
	Source               *string              `json:"source,omitempty"`                 
	OriImageID           *int64               `json:"oriImageId,omitempty"`             
	SenderID             *int64               `json:"sender_id,omitempty"`              
	DmType               *int64               `json:"dm_type,omitempty"`                
	StatusID             *int64               `json:"status_id,omitempty"`              
	SenderScreenName     *string              `json:"sender_screen_name,omitempty"`     
	IsLargeDm            *bool                `json:"isLargeDm,omitempty"`              
	SubType              *int64               `json:"sub_type,omitempty"`               
	MediaType            *int64               `json:"media_type,omitempty"`             
	MatchKeyword         *bool                `json:"matchKeyword,omitempty"`           
	ID                   int64                `json:"id"`                               
	Text                 *string              `json:"text,omitempty"`                   
	RecipientID          *int64               `json:"recipient_id,omitempty"`           
	IP                   *int64               `json:"ip,omitempty"`                     
	MsgStatus            *int64               `json:"msg_status,omitempty"`             
	EXTText              *EXTText             `json:"ext_text,omitempty"`               
	RecallStatus         *int64               `json:"recall_status,omitempty"`          
	RecipientScreenName  *string              `json:"recipient_screen_name,omitempty"`  
	BurnTime             *int64               `json:"burn_time,omitempty"`              
	GroupChatXinfo       *GroupChatXinfoUnion `json:"group_chat_xinfo"`                 
	URLObjects           []URLObject          `json:"url_objects"`                      
	Topublic             *bool                `json:"topublic,omitempty"`               
	GroupChatMessageType *int64               `json:"group_chat_message_type,omitempty"`
	AttIDS               []int64              `json:"att_ids"`                          
	Template             *string              `json:"template,omitempty"`               
	Data                 *Data                `json:"data,omitempty"`                   
	GroupChatMessageUids []int64              `json:"group_chat_message_uids"`          
	Msg                  *string              `json:"msg,omitempty"`                    
	Gid                  *int64               `json:"gid,omitempty"`                    
	Uids                 []int64              `json:"uids"`                             
	Time                 *int64               `json:"time,omitempty"`                   
	Type                 *int64               `json:"type,omitempty"`                   
	FromUid              *int64               `json:"from_uid,omitempty"`               
	Geo                  *string              `json:"geo,omitempty"`                    
	ReceiverBoxType      *string              `json:"receiver_box_type,omitempty"`      
	SenderTrashUser      *int64               `json:"sender_trash_user,omitempty"`      
	PushToMPS            *bool                `json:"pushToMPS,omitempty"`              
	Sender               *Recipient           `json:"sender,omitempty"`                 
	Recipient            *Recipient           `json:"recipient,omitempty"`              
}

type Data struct {
	Join     *Join `json:"join,omitempty"`    
	Operator *Join `json:"operator,omitempty"`
	Name     *Join `json:"name,omitempty"`    
}

type Join struct {
	Color     string  `json:"color"`               
	Scheme    *string `json:"scheme,omitempty"`    
	ColorDark *string `json:"color_dark,omitempty"`
	Value     string  `json:"value"`               
}

type EXTText struct {
	SendFrom      *SendFrom `json:"send_from,omitempty"`      
	Spr           *string   `json:"spr,omitempty"`            
	Gdid          *string   `json:"gdid,omitempty"`           
	IncludeVideo  *int64    `json:"include_video,omitempty"`  
	VideoPicFid   *int64    `json:"video_pic_fid,omitempty"`  
	AutoReply     *bool     `json:"autoReply,omitempty"`      
	SenderBoxType *string   `json:"sender_box_type,omitempty"`
	Clientid      *string   `json:"clientid,omitempty"`       
	IsPenetrate   *bool     `json:"is_penetrate,omitempty"`   
	Following     *bool     `json:"following,omitempty"`      
}

type GroupChatXinfoClass struct {
	MediaType   *int64  `json:"media_type,omitempty"`  
	FaithStatus *int64  `json:"faith_status,omitempty"`
	Template    *string `json:"template,omitempty"`    
	Data        *Data   `json:"data,omitempty"`        
}

type Recipient struct {
	IsTeenager             int64            `json:"is_teenager"`                       
	AllowAllActMsg         bool             `json:"allow_all_act_msg"`                 
	FavouritesCount        int64            `json:"favourites_count"`                  
	Urank                  int64            `json:"urank"`                             
	VerifiedTrade          string           `json:"verified_trade"`                    
	Weihao                 string           `json:"weihao"`                            
	VerifiedSourceURL      string           `json:"verified_source_url"`               
	Province               string           `json:"province"`                          
	ScreenName             string           `json:"screen_name"`                       
	ID                     int64            `json:"id"`                                
	GeoEnabled             bool             `json:"geo_enabled"`                       
	LikeMe                 bool             `json:"like_me"`                           
	Like                   bool             `json:"like"`                              
	VerifiedType           int64            `json:"verified_type"`                     
	PagefriendsCount       int64            `json:"pagefriends_count"`                 
	Domain                 string           `json:"domain"`                            
	Following              bool             `json:"following"`                         
	Name                   string           `json:"name"`                              
	Idstr                  string           `json:"idstr"`                             
	FollowMe               bool             `json:"follow_me"`                         
	FriendsCount           int64            `json:"friends_count"`                     
	CreditScore            int64            `json:"credit_score"`                      
	Gender                 string           `json:"gender"`                            
	City                   string           `json:"city"`                              
	ProfileURL             string           `json:"profile_url"`                       
	Description            string           `json:"description"`                       
	CreatedAt              string           `json:"created_at"`                        
	Remark                 string           `json:"remark"`                            
	Ptype                  int64            `json:"ptype"`                             
	VerifiedReasonURL      string           `json:"verified_reason_url"`               
	BlockWord              int64            `json:"block_word"`                        
	AvatarHD               string           `json:"avatar_hd"`                         
	Mbtype                 int64            `json:"mbtype"`                            
	BIFollowersCount       int64            `json:"bi_followers_count"`                
	UserAbility            int64            `json:"user_ability"`                      
	IsTeenagerList         int64            `json:"is_teenager_list"`                  
	VerifiedReason         string           `json:"verified_reason"`                   
	StoryReadState         int64            `json:"story_read_state"`                  
	VideoStatusCount       int64            `json:"video_status_count"`                
	Mbrank                 int64            `json:"mbrank"`                            
	Lang                   string           `json:"lang"`                              
	Class                  int64            `json:"class"`                             
	Star                   int64            `json:"star"`                              
	AllowAllComment        bool             `json:"allow_all_comment"`                 
	OnlineStatus           int64            `json:"online_status"`                     
	Verified               bool             `json:"verified"`                          
	ProfileImageURL        string           `json:"profile_image_url"`                 
	BlockApp               int64            `json:"block_app"`                         
	URL                    string           `json:"url"`                               
	AvatarLarge            string           `json:"avatar_large"`                      
	StatusesCount          int64            `json:"statuses_count"`                    
	VclubMember            int64            `json:"vclub_member"`                      
	FollowersCount         int64            `json:"followers_count"`                   
	IsGuardian             int64            `json:"is_guardian"`                       
	Location               string           `json:"location"`                          
	Insecurity             Insecurity       `json:"insecurity"`                        
	VerifiedSource         string           `json:"verified_source"`                   
	CoverImage             *string          `json:"cover_image,omitempty"`             
	VerifiedContactMobile  *string          `json:"verified_contact_mobile,omitempty"` 
	AvatargjID             *string          `json:"avatargj_id,omitempty"`             
	VerifiedTypeEXT        *int64           `json:"verified_type_ext,omitempty"`       
	VerifiedContactEmail   *string          `json:"verified_contact_email,omitempty"`  
	CoverImagePhone        *string          `json:"cover_image_phone,omitempty"`       
	VerifiedState          *int64           `json:"verified_state,omitempty"`          
	Cardid                 *string          `json:"cardid,omitempty"`                  
	VerifiedLevel          *int64           `json:"verified_level,omitempty"`          
	HasServiceTel          *bool            `json:"has_service_tel,omitempty"`         
	AttitudeStyle          *string          `json:"attitude_style,omitempty"`          
	VerifiedContactName    *string          `json:"verified_contact_name,omitempty"`   
	VerifiedReasonModified *string          `json:"verified_reason_modified,omitempty"`
	Type                   *int64           `json:"type,omitempty"`                    
	Level                  *int64           `json:"level,omitempty"`                   
	UserLimit              *int64           `json:"user_limit,omitempty"`              
	TabManage              *string          `json:"tab_manage,omitempty"`              
	Extend                 *Extend          `json:"extend,omitempty"`                  
	Badge                  map[string]int64 `json:"badge,omitempty"`                   
	BadgeTop               *string          `json:"badge_top,omitempty"`               
	HasAbilityTag          *int64           `json:"has_ability_tag,omitempty"`         
	Ulevel                 *int64           `json:"ulevel,omitempty"`                  
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

type URLObject struct {
	SearchTopicReadCount  int64            `json:"search_topic_read_count"` 
	LikeCount             int64            `json:"like_count"`              
	ObjectID              *string          `json:"object_id,omitempty"`     
	FollowerCount         int64            `json:"follower_count"`          
	SuperTopicPhotoCount  int64            `json:"super_topic_photo_count"` 
	SearchTopicCount      int64            `json:"search_topic_count"`      
	SuperTopicStatusCount int64            `json:"super_topic_status_count"`
	PageID                *string          `json:"page_id,omitempty"`       
	AssoLikeCount         int64            `json:"asso_like_count"`         
	URLOri                string           `json:"url_ori"`                 
	IsActionType          bool             `json:"isActionType"`            
	CardInfoUnIntegrity   bool             `json:"card_info_un_integrity"`  
	Info                  Info             `json:"info"`                    
	Status                *Status          `json:"status,omitempty"`        
	Object                *URLObjectObject `json:"object,omitempty"`        
}

type Info struct {
	Result       bool   `json:"result"`       
	URLShort     string `json:"url_short"`    
	EXTStatus    int64  `json:"ext_status"`   
	URLLong      string `json:"url_long"`     
	Transcode    int64  `json:"transcode"`    
	Description  string `json:"description"`  
	Type         int64  `json:"type"`         
	Title        string `json:"title"`        
	LastModified int64  `json:"last_modified"`
}

type URLObjectObject struct {
	Uuidstr        string       `json:"uuidstr"`                  
	ObjectType     string       `json:"object_type"`              
	ActivateStatus string       `json:"activate_status"`          
	SafeStatus     int64        `json:"safe_status"`              
	ObjectID       string       `json:"object_id"`                
	UUID           int64        `json:"uuid"`                     
	ActStatus      string       `json:"act_status"`               
	ObjectDomainID string       `json:"object_domain_id"`         
	Containerid    string       `json:"containerid"`              
	ShowStatus     string       `json:"show_status"`              
	LastModified   string       `json:"last_modified"`            
	Timestamp      int64        `json:"timestamp"`                
	Object         ObjectObject `json:"object"`                   
	SubObjectType  *string      `json:"sub_object_type,omitempty"`
}

type ObjectObject struct {
	Owner           *int64       `json:"owner,omitempty"`            
	Summary         *string      `json:"summary,omitempty"`          
	Image           *Image       `json:"image,omitempty"`            
	Gid             *int64       `json:"gid,omitempty"`              
	OwnerName       *string      `json:"owner_name,omitempty"`       
	ObjectType      string       `json:"object_type"`                
	GroupName       *string      `json:"group_name,omitempty"`       
	GroupMembers    []int64      `json:"group_members"`              
	TargetURL       string       `json:"target_url"`                 
	DisplayName     string       `json:"display_name"`               
	URL             string       `json:"url"`                        
	Biz             Biz          `json:"biz"`                        
	ID              string       `json:"id"`                         
	MemberCount     *int64       `json:"member_count,omitempty"`     
	Scheme          *string      `json:"scheme,omitempty"`           
	Author          *Author      `json:"author,omitempty"`           
	FullImage       *FullImage   `json:"full_image,omitempty"`       
	ToolbarHidden   *int64       `json:"toolbar_hidden,omitempty"`   
	CreateAt        *string      `json:"create_at,omitempty"`        
	Note            *string      `json:"note,omitempty"`             
	OID             *OID         `json:"oid"`                        
	Title           *string      `json:"title,omitempty"`            
	TopcolorDark    *string      `json:"topcolor_dark,omitempty"`    
	Topcolor        *string      `json:"topcolor,omitempty"`         
	ContentSummary  *string      `json:"content_summary,omitempty"`  
	LinkDisplayName *string      `json:"link_display_name,omitempty"`
	UpdateAt        *string      `json:"update_at,omitempty"`        
	Headimgurl      *string      `json:"headimgurl,omitempty"`       
	ContentTemplate *string      `json:"content_template,omitempty"` 
	ContentData     *ContentData `json:"content_data,omitempty"`     
	ObjectTypes     []string     `json:"object_types"`               
	OriginalMid     *string      `json:"original_mid,omitempty"`     
	ScreenName      *string      `json:"screen_name,omitempty"`      
	Updated         *string      `json:"updated,omitempty"`          
	Items           []Item       `json:"items"`                      
}

type Author struct {
	ObjectType  string `json:"object_type"` 
	ID          *OID   `json:"id"`          
	DisplayName string `json:"display_name"`
	URL         string `json:"url"`         
}

type Biz struct {
	BizID       string `json:"biz_id"`     
	Containerid string `json:"containerid"`
}

type ContentData struct {
	Data     *Content `json:"data,omitempty"`    
	Jifen    *Jifen   `json:"jifen,omitempty"`   
	Time     Join     `json:"time"`              
	Mark     *Join    `json:"mark,omitempty"`    
	Location *Device  `json:"location,omitempty"`
	Remark   *Device  `json:"remark,omitempty"`  
	Title    *Device  `json:"title,omitempty"`   
	Device   *Device  `json:"device,omitempty"`  
	Content  *Content `json:"content,omitempty"` 
}

type Content struct {
	Value string `json:"value"`
}

type Device struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

type Jifen struct {
	Value int64 `json:"value"`
}

type FullImage struct {
	Width  string `json:"width"` 
	URL    string `json:"url"`   
	Height string `json:"height"`
}

type Image struct {
	RoundURL *string `json:"round_url,omitempty"`
	URL      string  `json:"url"`                
	Width    *OID    `json:"width"`              
	Height   *OID    `json:"height"`             
}

type Item struct {
	Summary     string `json:"summary"`     
	Scheme      string `json:"scheme"`      
	ObjectType  string `json:"object_type"` 
	Pic         string `json:"pic"`         
	ID          string `json:"id"`          
	DisplayName string `json:"display_name"`
	URL         string `json:"url"`         
}

type Status struct {
	IsLongText               bool                  `json:"isLongText"`                
	HotWeiboTags             []interface{}         `json:"hot_weibo_tags"`            
	InReplyToStatusID        string                `json:"in_reply_to_status_id"`     
	PicIDS                   []string              `json:"pic_ids"`                   
	PicTypes                 string                `json:"pic_types"`                 
	Annotations              []Annotation          `json:"annotations"`               
	Mblogtype                int64                 `json:"mblogtype"`                 
	Source                   string                `json:"source"`                    
	AttitudesCount           int64                 `json:"attitudes_count"`           
	Rid                      string                `json:"rid"`                       
	BmiddlePic               *string               `json:"bmiddle_pic,omitempty"`     
	PositiveRecomFlag        int64                 `json:"positive_recom_flag"`       
	IsShowBulletin           int64                 `json:"is_show_bulletin"`          
	HideFlag                 int64                 `json:"hide_flag"`                 
	Mlevel                   int64                 `json:"mlevel"`                    
	InReplyToUserID          string                `json:"in_reply_to_user_id"`       
	HasActionTypeCard        int64                 `json:"hasActionTypeCard"`         
	ID                       int64                 `json:"id"`                        
	Text                     string                `json:"text"`                      
	PicStatus                *string               `json:"picStatus,omitempty"`       
	MblogVipType             int64                 `json:"mblog_vip_type"`            
	ContentAuth              int64                 `json:"content_auth"`              
	Visible                  Visible               `json:"visible"`                   
	GIFIDS                   string                `json:"gif_ids"`                   
	SourceType               int64                 `json:"source_type"`               
	BizFeature               int64                 `json:"biz_feature"`               
	CommentsCount            int64                 `json:"comments_count"`            
	UserType                 int64                 `json:"userType"`                  
	Idstr                    string                `json:"idstr"`                     
	TextTagTips              []interface{}         `json:"text_tag_tips"`             
	CreatedAt                string                `json:"created_at"`                
	Mid                      string                `json:"mid"`                       
	DarwinTags               []interface{}         `json:"darwin_tags"`               
	PendingApprovalCount     int64                 `json:"pending_approval_count"`    
	InReplyToScreenName      string                `json:"in_reply_to_screen_name"`   
	PicNum                   int64                 `json:"pic_num"`                   
	IsPaid                   bool                  `json:"is_paid"`                   
	RepostsCount             int64                 `json:"reposts_count"`             
	RewardExhibitionType     int64                 `json:"reward_exhibition_type"`    
	Favorited                bool                  `json:"favorited"`                 
	ExternSafe               int64                 `json:"extern_safe"`               
	ThumbnailPic             *string               `json:"thumbnail_pic,omitempty"`   
	OriginalPic              *string               `json:"original_pic,omitempty"`    
	CanEdit                  bool                  `json:"can_edit"`                  
	TextLength               int64                 `json:"textLength"`                
	Truncated                bool                  `json:"truncated"`                 
	NumberDisplayStrategy    NumberDisplayStrategy `json:"number_display_strategy"`   
	SourceAllowclick         int64                 `json:"source_allowclick"`         
	ShowAdditionalIndication int64                 `json:"show_additional_indication"`
	Appid                    int64                 `json:"appid"`                     
	CommentManageInfo        CommentManageInfo     `json:"comment_manage_info"`       
	User                     Recipient             `json:"user"`                      
	MoreInfoType             int64                 `json:"more_info_type"`            
	BizIDS                   []int64               `json:"biz_ids"`                   
}

type Annotation struct {
	ClientMblogid *string `json:"client_mblogid,omitempty"`
	MAPIRequest   *bool   `json:"mapi_request,omitempty"`  
	Shooting      *int64  `json:"shooting,omitempty"`      
}

type CommentManageInfo struct {
	CommentPermissionType int64 `json:"comment_permission_type"`
	ApprovalCommentType   int64 `json:"approval_comment_type"`  
}

type NumberDisplayStrategy struct {
	ApplyScenarioFlag    int64  `json:"apply_scenario_flag"`    
	DisplayTextMinNumber int64  `json:"display_text_min_number"`
	DisplayText          string `json:"display_text"`           
}

type Visible struct {
	ListID int64 `json:"list_id"`
	Type   int64 `json:"type"`   
}

type SignificantMsg struct {
	Color           string `json:"color"`           
	SignificantType int64  `json:"significant_type"`
	Level           int64  `json:"level"`           
	ColorDark       string `json:"color_dark"`      
	Mid             int64  `json:"mid"`             
	Type            string `json:"type"`            
	Title           string `json:"title"`           
}

type User struct {
	Idstr                *string `json:"idstr,omitempty"`                  
	FollowMe             *bool   `json:"follow_me,omitempty"`              
	Addsession           *bool   `json:"addsession,omitempty"`             
	ValidateType         *int64  `json:"validateType,omitempty"`           
	ProfileURL           *string `json:"profile_url,omitempty"`            
	Remark               *string `json:"remark,omitempty"`                 
	GroupType            *int64  `json:"group_type,omitempty"`             
	Type                 int64   `json:"type"`                             
	SuperGroupType       *int64  `json:"super_group_type,omitempty"`       
	JoinTime             *int64  `json:"join_time,omitempty"`              
	AffiType             []int64 `json:"affi_type"`                        
	RoundAvatarLarge     *string `json:"round_avatar_large,omitempty"`     
	RoundProfileImageURL *string `json:"round_profile_image_url,omitempty"`
	Members              []int64 `json:"members"`                          
	TrashUser            *int64  `json:"trash_user,omitempty"`             
	ID                   int64   `json:"id"`                               
	MemberCount          *int64  `json:"member_count,omitempty"`           
	Dotted               *bool   `json:"dotted,omitempty"`                 
	UserCustomMsgSetting *string `json:"user_custom_msg_setting,omitempty"`
	Creator              *int64  `json:"creator,omitempty"`                
	Level                *int64  `json:"level,omitempty"`                  
	Verified             *bool   `json:"verified,omitempty"`               
	RemindSetting        *int64  `json:"remindSetting,omitempty"`          
	AffiliationPageIDS   []int64 `json:"affiliationPageIds"`               
	ProfileImageURL      string  `json:"profile_image_url"`                
	VerifiedType         *int64  `json:"verified_type,omitempty"`          
	MaxMemberCount       *int64  `json:"max_member_count,omitempty"`       
	VerifiedTypeEXT      *int64  `json:"verified_type_ext,omitempty"`      
	AvatarLarge          string  `json:"avatar_large"`                     
	IsBlocked            *bool   `json:"is_blocked,omitempty"`             
	Following            *bool   `json:"following,omitempty"`              
	LastChangeTime       *int64  `json:"lastChangeTime,omitempty"`         
	Name                 string  `json:"name"`                             
	Admins               []int64 `json:"admins"`                           
}

type SendFrom string
const (
	Contact SendFrom = "contact"
	Msg SendFrom = "msg"
	Share SendFrom = "share"
	Webchat SendFrom = "webchat"
)

type GroupChatXinfoUnion struct {
	GroupChatXinfoClass *GroupChatXinfoClass
	String              *string
}

func (x *GroupChatXinfoUnion) UnmarshalJSON(data []byte) error {
	x.GroupChatXinfoClass = nil
	var c GroupChatXinfoClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.GroupChatXinfoClass = &c
	}
	return nil
}

func (x *GroupChatXinfoUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.GroupChatXinfoClass != nil, x.GroupChatXinfoClass, false, nil, false, nil, false)
}

type OID struct {
	Integer *int64
	String  *string
}

func (x *OID) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *OID) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
