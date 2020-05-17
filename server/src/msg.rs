
// Example code that deserializes and serializes the model.
// extern crate serde;
// #[macro_use]
// extern crate serde_derive;
// extern crate serde_json;
//
// use generated_module::[object Object];
//
// fn main() {
//     let json = r#"{"answer": 42}"#;
//     let model: [object Object] = serde_json::from_str(&json).unwrap();
// }

extern crate serde_json;
use std::collections::HashMap;



#[derive(Serialize, Deserialize)]
pub struct Msg{
    #[serde(rename = "last_read_mid")]
    last_read_mid: i64,

    #[serde(rename = "significant_msgs")]
    significant_msgs: Vec<Option<serde_json::Value>>,

    #[serde(rename = "group_tips")]
    group_tips: Vec<Option<serde_json::Value>>,

    #[serde(rename = "messages")]
    messages: Vec<Message>,

    #[serde(rename = "result")]
    result: bool,

    #[serde(rename = "ts")]
    ts: i64,
}


#[derive(Serialize, Deserialize)]
pub struct Message{
    #[serde(rename = "annotations")]
    annotations: Annotations,

    #[serde(rename = "appid")]
    appid: i64,

    #[serde(rename = "content")]
    content: String,

    #[serde(rename = "faith_status")]
    faith_status: Option<i64>,

    #[serde(rename = "fids")]
    fids: Option<Vec<i64>>,

    #[serde(rename = "from_uid")]
    from_uid: i64,

    #[serde(rename = "from_user")]
    from_user: FromUser,

    #[serde(rename = "gid")]
    gid: i64,

    #[serde(rename = "id")]
    id: i64,

    #[serde(rename = "mblogid")]
    mblogid: String,

    #[serde(rename = "media_type")]
    media_type: i64,

    #[serde(rename = "recall_status")]
    recall_status: i64,

    #[serde(rename = "time")]
    time: i64,

    #[serde(rename = "type")]
    top_level_type: i64,

    #[serde(rename = "url_objects")]
    url_objects: Option<Vec<UrlObject>>,
}

#[derive(Serialize, Deserialize)]
pub struct Annotations {
    #[serde(rename = "gdid")]
    gdid: String,

    #[serde(rename = "include_video")]
    include_video: i64,

    #[serde(rename = "send_from")]
    send_from: String,

    #[serde(rename = "spr")]
    spr: String,
}

#[derive(Serialize, Deserialize)]
pub struct FromUser {
    #[serde(rename = "avatar_large")]
    avatar_large: String,

    #[serde(rename = "followers_count")]
    followers_count: i64,

    #[serde(rename = "friends_count")]
    friends_count: i64,

    #[serde(rename = "id")]
    id: i64,

    #[serde(rename = "level")]
    level: i64,

    #[serde(rename = "profile_image_url")]
    profile_image_url: String,

    #[serde(rename = "profile_url")]
    profile_url: String,

    #[serde(rename = "screen_name")]
    screen_name: String,

    #[serde(rename = "verified")]
    verified: bool,
}

#[derive(Serialize, Deserialize)]
pub struct UrlObject {
    #[serde(rename = "asso_like_count")]
    asso_like_count: i64,

    #[serde(rename = "card_info_un_integrity")]
    card_info_un_integrity: bool,

    #[serde(rename = "follower_count")]
    follower_count: i64,

    #[serde(rename = "info")]
    info: Info,

    #[serde(rename = "isActionType")]
    is_action_type: bool,

    #[serde(rename = "like_count")]
    like_count: i64,

    #[serde(rename = "object_id")]
    object_id: String,

    #[serde(rename = "page_id")]
    page_id: String,

    #[serde(rename = "search_topic_count")]
    search_topic_count: i64,

    #[serde(rename = "search_topic_read_count")]
    search_topic_read_count: i64,

    #[serde(rename = "status")]
    status: Status,

    #[serde(rename = "super_topic_photo_count")]
    super_topic_photo_count: i64,

    #[serde(rename = "super_topic_status_count")]
    super_topic_status_count: i64,

    #[serde(rename = "url_ori")]
    url_ori: String,
}

#[derive(Serialize, Deserialize)]
pub struct Info {
    #[serde(rename = "description")]
    description: String,

    #[serde(rename = "ext_status")]
    ext_status: i64,

    #[serde(rename = "last_modified")]
    last_modified: i64,

    #[serde(rename = "result")]
    result: bool,

    #[serde(rename = "title")]
    title: String,

    #[serde(rename = "transcode")]
    transcode: i64,

    #[serde(rename = "type")]
    info_type: i64,

    #[serde(rename = "url_long")]
    url_long: String,

    #[serde(rename = "url_short")]
    url_short: String,
}

#[derive(Serialize, Deserialize)]
pub struct Status {
    #[serde(rename = "annotations")]
    annotations: Vec<Annotation>,

    #[serde(rename = "appid")]
    appid: i64,

    #[serde(rename = "attitudes_count")]
    attitudes_count: i64,

    #[serde(rename = "biz_feature")]
    biz_feature: i64,

    #[serde(rename = "biz_ids")]
    biz_ids: Vec<i64>,

    #[serde(rename = "can_edit")]
    can_edit: bool,

    #[serde(rename = "cardid")]
    cardid: String,

    #[serde(rename = "comment_manage_info")]
    comment_manage_info: CommentManageInfo,

    #[serde(rename = "comments_count")]
    comments_count: i64,

    #[serde(rename = "content_auth")]
    content_auth: i64,

    #[serde(rename = "created_at")]
    created_at: String,

    #[serde(rename = "darwin_tags")]
    darwin_tags: Vec<DarwinTag>,

    #[serde(rename = "extern_safe")]
    extern_safe: i64,

    #[serde(rename = "favorited")]
    favorited: bool,

    #[serde(rename = "geo")]
    geo: Geo,

    #[serde(rename = "gif_ids")]
    gif_ids: String,

    #[serde(rename = "hasActionTypeCard")]
    has_action_type_card: i64,

    #[serde(rename = "hide_flag")]
    hide_flag: i64,

    #[serde(rename = "hot_weibo_tags")]
    hot_weibo_tags: Vec<Option<serde_json::Value>>,

    #[serde(rename = "id")]
    id: i64,

    #[serde(rename = "idstr")]
    idstr: String,

    #[serde(rename = "in_reply_to_screen_name")]
    in_reply_to_screen_name: String,

    #[serde(rename = "in_reply_to_status_id")]
    in_reply_to_status_id: String,

    #[serde(rename = "in_reply_to_user_id")]
    in_reply_to_user_id: String,

    #[serde(rename = "isLongText")]
    is_long_text: bool,

    #[serde(rename = "is_paid")]
    is_paid: bool,

    #[serde(rename = "is_show_bulletin")]
    is_show_bulletin: i64,

    #[serde(rename = "mblog_vip_type")]
    mblog_vip_type: i64,

    #[serde(rename = "mblogtype")]
    mblogtype: i64,

    #[serde(rename = "mid")]
    mid: String,

    #[serde(rename = "mlevel")]
    mlevel: i64,

    #[serde(rename = "more_info_type")]
    more_info_type: i64,

    #[serde(rename = "number_display_strategy")]
    number_display_strategy: NumberDisplayStrategy,

    #[serde(rename = "pending_approval_count")]
    pending_approval_count: i64,

    #[serde(rename = "pic_ids")]
    pic_ids: Vec<String>,

    #[serde(rename = "pic_num")]
    pic_num: i64,

    #[serde(rename = "pic_types")]
    pic_types: String,

    #[serde(rename = "positive_recom_flag")]
    positive_recom_flag: i64,

    #[serde(rename = "reposts_count")]
    reposts_count: i64,

    #[serde(rename = "reward_exhibition_type")]
    reward_exhibition_type: i64,

    #[serde(rename = "rid")]
    rid: String,

    #[serde(rename = "show_additional_indication")]
    show_additional_indication: i64,

    #[serde(rename = "source")]
    source: String,

    #[serde(rename = "source_allowclick")]
    source_allowclick: i64,

    #[serde(rename = "source_type")]
    source_type: i64,

    #[serde(rename = "text")]
    text: String,

    #[serde(rename = "textLength")]
    text_length: i64,

    #[serde(rename = "text_tag_tips")]
    text_tag_tips: Vec<Option<serde_json::Value>>,

    #[serde(rename = "truncated")]
    truncated: bool,

    #[serde(rename = "user")]
    user: User,

    #[serde(rename = "userType")]
    user_type: i64,

    #[serde(rename = "visible")]
    visible: Visible,
}

#[derive(Serialize, Deserialize)]
pub struct Annotation {
    #[serde(rename = "oid")]
    oid: String,

    #[serde(rename = "type")]
    annotation_type: String,
}

#[derive(Serialize, Deserialize)]
pub struct CommentManageInfo {
    #[serde(rename = "approval_comment_type")]
    approval_comment_type: i64,

    #[serde(rename = "comment_permission_type")]
    comment_permission_type: i64,
}

#[derive(Serialize, Deserialize)]
pub struct DarwinTag {
    #[serde(rename = "bd_object_type")]
    bd_object_type: String,

    #[serde(rename = "display_name")]
    display_name: String,

    #[serde(rename = "enterprise_uid")]
    enterprise_uid: Option<serde_json::Value>,

    #[serde(rename = "object_id")]
    object_id: String,

    #[serde(rename = "object_type")]
    object_type: String,
}

#[derive(Serialize, Deserialize)]
pub struct Geo {
    #[serde(rename = "coordinates")]
    coordinates: Vec<f64>,

    #[serde(rename = "type")]
    geo_type: String,
}

#[derive(Serialize, Deserialize)]
pub struct NumberDisplayStrategy {
    #[serde(rename = "apply_scenario_flag")]
    apply_scenario_flag: i64,

    #[serde(rename = "display_text")]
    display_text: String,

    #[serde(rename = "display_text_min_number")]
    display_text_min_number: i64,
}

#[derive(Serialize, Deserialize)]
pub struct User {
    #[serde(rename = "allow_all_act_msg")]
    allow_all_act_msg: bool,

    #[serde(rename = "allow_all_comment")]
    allow_all_comment: bool,

    #[serde(rename = "avatar_hd")]
    avatar_hd: String,

    #[serde(rename = "avatar_large")]
    avatar_large: String,

    #[serde(rename = "badge")]
    badge: HashMap<String, i64>,

    #[serde(rename = "badge_top")]
    badge_top: String,

    #[serde(rename = "bi_followers_count")]
    bi_followers_count: i64,

    #[serde(rename = "block_app")]
    block_app: i64,

    #[serde(rename = "block_word")]
    block_word: i64,

    #[serde(rename = "cardid")]
    cardid: String,

    #[serde(rename = "city")]
    city: String,

    #[serde(rename = "class")]
    class: i64,

    #[serde(rename = "cover_image_phone")]
    cover_image_phone: String,

    #[serde(rename = "created_at")]
    created_at: String,

    #[serde(rename = "credit_score")]
    credit_score: i64,

    #[serde(rename = "description")]
    description: String,

    #[serde(rename = "domain")]
    domain: String,

    #[serde(rename = "extend")]
    extend: Extend,

    #[serde(rename = "favourites_count")]
    favourites_count: i64,

    #[serde(rename = "follow_me")]
    follow_me: bool,

    #[serde(rename = "followers_count")]
    followers_count: i64,

    #[serde(rename = "following")]
    following: bool,

    #[serde(rename = "friends_count")]
    friends_count: i64,

    #[serde(rename = "gender")]
    gender: String,

    #[serde(rename = "geo_enabled")]
    geo_enabled: bool,

    #[serde(rename = "has_ability_tag")]
    has_ability_tag: i64,

    #[serde(rename = "id")]
    id: i64,

    #[serde(rename = "idstr")]
    idstr: String,

    #[serde(rename = "insecurity")]
    insecurity: Insecurity,

    #[serde(rename = "is_guardian")]
    is_guardian: i64,

    #[serde(rename = "is_teenager")]
    is_teenager: i64,

    #[serde(rename = "is_teenager_list")]
    is_teenager_list: i64,

    #[serde(rename = "lang")]
    lang: String,

    #[serde(rename = "level")]
    level: i64,

    #[serde(rename = "like")]
    like: bool,

    #[serde(rename = "like_me")]
    like_me: bool,

    #[serde(rename = "location")]
    location: String,

    #[serde(rename = "mbrank")]
    mbrank: i64,

    #[serde(rename = "mbtype")]
    mbtype: i64,

    #[serde(rename = "name")]
    name: String,

    #[serde(rename = "online_status")]
    online_status: i64,

    #[serde(rename = "pagefriends_count")]
    pagefriends_count: i64,

    #[serde(rename = "profile_image_url")]
    profile_image_url: String,

    #[serde(rename = "profile_url")]
    profile_url: String,

    #[serde(rename = "province")]
    province: String,

    #[serde(rename = "ptype")]
    ptype: i64,

    #[serde(rename = "remark")]
    remark: String,

    #[serde(rename = "screen_name")]
    screen_name: String,

    #[serde(rename = "special_follow")]
    special_follow: bool,

    #[serde(rename = "star")]
    star: i64,

    #[serde(rename = "statuses_count")]
    statuses_count: i64,

    #[serde(rename = "story_read_state")]
    story_read_state: i64,

    #[serde(rename = "tab_manage")]
    tab_manage: String,

    #[serde(rename = "type")]
    user_type: i64,

    #[serde(rename = "ulevel")]
    ulevel: i64,

    #[serde(rename = "urank")]
    urank: i64,

    #[serde(rename = "url")]
    url: String,

    #[serde(rename = "user_ability")]
    user_ability: i64,

    #[serde(rename = "user_limit")]
    user_limit: i64,

    #[serde(rename = "vclub_member")]
    vclub_member: i64,

    #[serde(rename = "verified")]
    verified: bool,

    #[serde(rename = "verified_reason")]
    verified_reason: String,

    #[serde(rename = "verified_reason_url")]
    verified_reason_url: String,

    #[serde(rename = "verified_source")]
    verified_source: String,

    #[serde(rename = "verified_source_url")]
    verified_source_url: String,

    #[serde(rename = "verified_trade")]
    verified_trade: String,

    #[serde(rename = "verified_type")]
    verified_type: i64,

    #[serde(rename = "video_status_count")]
    video_status_count: i64,

    #[serde(rename = "weihao")]
    weihao: String,
}

#[derive(Serialize, Deserialize)]
pub struct Extend {
    #[serde(rename = "mbprivilege")]
    mbprivilege: String,

    #[serde(rename = "privacy")]
    privacy: Privacy,
}

#[derive(Serialize, Deserialize)]
pub struct Privacy {
    #[serde(rename = "mobile")]
    mobile: i64,
}

#[derive(Serialize, Deserialize)]
pub struct Insecurity {
    #[serde(rename = "sexual_content")]
    sexual_content: bool,
}

#[derive(Serialize, Deserialize)]
pub struct Visible {
    #[serde(rename = "list_id")]
    list_id: i64,

    #[serde(rename = "type")]
    visible_type: i64,
}
