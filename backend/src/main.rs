mod api;
mod db;

#[tokio::main]
async fn main() {
    env_logger::init();
    let pool = db::create_pool("postgres://postgres@localhost:5432/flashcow?password=flashcow");
    api::serve("0.0.0.0:5000".parse().unwrap(), pool).await;
}
