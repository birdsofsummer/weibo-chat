#   https://diesel.rs/docs/
#   http://diesel.rs/guides/getting-started/
#   https://github.com/diesel-rs/diesel/tree/v1.4.4/examples/postgres/getting_started_step_3/


#cargo install diesel_cli
#note: ld: library not found for -lmysqlclient
cargo install diesel_cli --no-default-features --features postgres
diesel setup


diesel migration generate create_posts
#see migrations/YYYY-MM-DD-HHMMSS_create_posts/up.sql
#see migrations/YYYY-MM-DD-HHMMSS_create_posts/down.sql
diesel migration run
diesel migration redo


cargo run --bin write_post   # hehe cccc  ^d
cargo run --bin publish_post 1
cargo run --bin show_posts
cargo run --bin delete_post hehe
cargo run --bin show_posts
