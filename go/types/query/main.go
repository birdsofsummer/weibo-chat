// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalTopLevel(bytes)
//    bytes, err = topLevel.Marshal()

package query

import "encoding/json"

func UnmarshalTopLevel(data []byte) (TopLevel, error) {
	var r TopLevel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TopLevel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TopLevel struct {
	HasRank              bool     `json:"has_rank"`               
	ID                   int64    `json:"id"`                     
	PageObjectid         string   `json:"page_objectid"`          
	Name                 string   `json:"name"`                   
	SystemName           string   `json:"system_name"`            
	Avatar               string   `json:"avatar"`                 
	AvatarS              string   `json:"avatar_s"`               
	RoundAvatar          string   `json:"round_avatar"`           
	RoundAvatarS         string   `json:"round_avatar_s"`         
	GroupType            int64    `json:"group_type"`             
	SuperGroupType       int64    `json:"super_group_type"`       
	Summary              string   `json:"summary"`                
	MaxMember            int64    `json:"max_member"`             
	Owner                int64    `json:"owner"`                  
	MemberCount          int64    `json:"member_count"`           
	CreateTime           int64    `json:"create_time"`            
	ValidateType         int64    `json:"validate_type"`          
	GroupTs              int64    `json:"group_ts"`               
	Status               int64    `json:"status"`                 
	Active               int64    `json:"active"`                 
	Publicity            int64    `json:"publicity"`              
	AffiJoinType         int64    `json:"affi_join_type"`         
	Emeccs               []int64  `json:"emeccs"`                 
	AffiType             []int64  `json:"affi_type"`              
	Affiliation          []int64  `json:"affiliation"`            
	JoinTime             int64    `json:"join_time"`              
	BeginMid             int64    `json:"begin_mid"`              
	Push                 int64    `json:"push"`                   
	Addsession           int64    `json:"addsession"`             
	Filterfeed           int64    `json:"filterfeed"`             
	PushAirborne         int64    `json:"push_airborne"`          
	UserCustomMsgSetting string   `json:"user_custom_msg_setting"`
	Filterquery          int64    `json:"filterquery"`            
	GroupURL             string   `json:"group_url"`              
	GlobalMaxAdmin       int64    `json:"global_max_admin"`       
	MaxAdmin             int64    `json:"max_admin"`              
	Admins               []int64  `json:"admins"`                 
	SpeechForbid         string   `json:"speech_forbid"`          
	RemainAtCount        int64    `json:"remain_at_count"`        
	Bulletin             Bulletin `json:"bulletin"`               
	Harm                 int64    `json:"harm"`                   
	InviteFansStatus     int64    `json:"invite_fans_status"`     
	Result               bool     `json:"result"`                 
	Ts                   int64    `json:"ts"`                     
}

type Bulletin struct {
	KeepTime   int64  `json:"keep_time"`  
	Scheme     string `json:"scheme"`     
	BulletinID int64  `json:"bulletin_id"`
	Content    string `json:"content"`    
}
