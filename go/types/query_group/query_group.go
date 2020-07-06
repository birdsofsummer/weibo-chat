// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    topLevel, err := UnmarshalData(bytes)
//    bytes, err = topLevel.Marshal()

package query_group

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
	Addsession           int64        `json:"addsession"`             
	Harm                 int64        `json:"harm"`                   
	AvatarS              string       `json:"avatar_s"`               
	GroupType            int64        `json:"group_type"`             
	SuperGroupType       int64        `json:"super_group_type"`       
	AffiType             []int64      `json:"affi_type"`              
	JoinTime             int64        `json:"join_time"`              
	RemainAtCount        int64        `json:"remain_at_count"`        
	GroupURL             string       `json:"group_url"`              
	Result               bool         `json:"result"`                 
	GroupTs              int64        `json:"group_ts"`               
	HasRank              bool         `json:"has_rank"`               
	Affiliation          []int64      `json:"affiliation"`            
	Members              []int64      `json:"members"`                
	SpeechForbid         string       `json:"speech_forbid"`          
	ID                   int64        `json:"id"`                     
	Publicity            int64        `json:"publicity"`              
	MemberCount          int64        `json:"member_count"`           
	PageObjectid         string       `json:"page_objectid"`          
	Summary              string       `json:"summary"`                
	Owner                int64        `json:"owner"`                  
	BeginMid             int64        `json:"begin_mid"`              
	Filterfeed           int64        `json:"filterfeed"`             
	UserCustomMsgSetting string       `json:"user_custom_msg_setting"`
	Filterquery          int64        `json:"filterquery"`            
	CreateTime           int64        `json:"create_time"`            
	SystemName           string       `json:"system_name"`            
	MaxMember            int64        `json:"max_member"`             
	Active               int64        `json:"active"`                 
	MemberInfos          []MemberInfo `json:"member_infos"`           
	Avatar               string       `json:"avatar"`                 
	PushAirborne         int64        `json:"push_airborne"`          
	Push                 int64        `json:"push"`                   
	MaxAdmin             int64        `json:"max_admin"`              
	Name                 string       `json:"name"`                   
	RoundAvatarS         string       `json:"round_avatar_s"`         
	ValidateType         int64        `json:"validate_type"`          
	AffiJoinType         int64        `json:"affi_join_type"`         
	GlobalMaxAdmin       int64        `json:"global_max_admin"`       
	Admins               []int64      `json:"admins"`                 
	Bulletin             Bulletin     `json:"bulletin"`               
	RoundAvatar          string       `json:"round_avatar"`           
	Status               int64        `json:"status"`                 
	Ts                   int64        `json:"ts"`                     
}

type Bulletin struct {
	KeepTime   int64  `json:"keep_time"`  
	Scheme     string `json:"scheme"`     
	BulletinID int64  `json:"bulletin_id"`
	Content    string `json:"content"`    
}

type MemberInfo struct {
	Uid             int64  `json:"uid"`              
	ScreenName      string `json:"screen_name"`      
	Level           int64  `json:"level"`            
	Jp              string `json:"jp"`               
	Verified        bool   `json:"verified"`         
	ProfileImageURL string `json:"profile_image_url"`
	Qp              string `json:"qp"`               
}
