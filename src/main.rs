mod page;

#[tokio::main]
async fn main() {
    env_logger::init();
    page::serve("0.0.0.0:5000".parse().unwrap()).await;
}
