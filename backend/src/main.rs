mod api;

#[tokio::main]
async fn main() {
    env_logger::init();
    api::serve("0.0.0.0:5000".parse().unwrap()).await;
}
