use std::{net::SocketAddr, sync::{atomic::{AtomicU32, Ordering}, Arc}};
use askama::Template;
use axum::{
    extract::State,
    response::{Html, IntoResponse},
    routing::get, Router
};

enum Error {
    TemplatingError(askama::Error),
}

impl From<askama::Error> for Error {
    fn from(value: askama::Error) -> Self {
        Error::TemplatingError(value)
    }
}

impl Error {
    pub fn to_html(&self) -> String {
        "Error!".to_owned()
    }
}

struct PageResult(Result<String, Error>);

impl From<String> for PageResult {
    fn from(value: String) -> Self {
        PageResult(Ok(value))
    }
}

impl<E> From<Result<String, E>> for PageResult where E: Into<Error> {
    fn from(value: Result<String, E>) -> Self {
        PageResult(value.map_err(|err| err.into()))
    }
}

impl IntoResponse for PageResult {
    fn into_response(self) -> axum::response::Response {
        match self.0 {
            Ok(response) => Html(response).into_response(),
            Err(err) => Html(err.to_html()).into_response(),
        }
    }
}

#[derive(Template)]
#[template(path = "test.html")]
struct TestPage {
    n: u32,
}

impl TestPage {
    pub async fn get(State(state): State<Arc<VisitorCounter>>) -> PageResult {
        let test_page_template = TestPage { n: state.n.fetch_add(1, Ordering::AcqRel) };
        test_page_template.render().into()
    }
}

struct VisitorCounter {
    pub n: AtomicU32,
}

pub async fn serve(http_addr: SocketAddr) {
    let app = Router::new()
        .route("/test.html", get(TestPage::get))
        .with_state(Arc::new(VisitorCounter { n: 0.into() }));
    
    let listener = tokio::net::TcpListener::bind(http_addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
