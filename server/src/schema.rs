table! {
    posts (id) {
        id -> Int4,
        title -> Varchar,
        body -> Text,
        published -> Bool,
    }
}

table! {
    user (id) {
        id -> Int8,
        usr_name -> Varchar,
    }
}

allow_tables_to_appear_in_same_query!(
    posts,
    user,
);
