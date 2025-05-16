use std::{net::SocketAddr, sync::{atomic::{AtomicU32, Ordering}, Arc}};

use axum::{body::Body, extract::State, http::{HeaderValue, StatusCode}, response::{IntoResponse, Response}, routing::get, Router};
use serde::Serialize;
use serde_json::json;

enum ApiError {
    JsonError(serde_json::Error),
}

impl IntoResponse for ApiError {
    fn into_response(self) -> axum::response::Response {
        let (status_code, json_obj) = match self {
            ApiError::JsonError(_json_error) => (
                StatusCode::INTERNAL_SERVER_ERROR,
                json!({
                    "error_code": 1
                })
            ),
        };

        if let Ok(json) = serde_json::to_string(&json_obj) {
            let mut response = axum::response::Response::new(Body::from(json));
            *response.status_mut() = status_code;
            response.headers_mut().insert("Content-Type", HeaderValue::from_str("application/json").unwrap());

            response
        } else {
            let mut response = axum::response::Response::new(Body::empty());
            *response.status_mut() = StatusCode::INTERNAL_SERVER_ERROR;
            
            response
        }
    }
}

enum ApiResult<R> {
    Success(R),
    Error(ApiError),
}

impl<R> IntoResponse for ApiResult<R> where R: Serialize {
    fn into_response(self) -> axum::response::Response {
        match self {
            ApiResult::Success(response_object) => {
                match serde_json::to_string(&response_object) {
                    Err(error) => ApiError::JsonError(error).into_response(),
                    Ok(response_string) => {
                        let mut response = axum::response::Response::new(Body::from(response_string));
                        response.headers_mut().insert("Content-Type", HeaderValue::from_str("application/json").unwrap());

                        response.into_response()
                    }
                }
            },
            ApiResult::Error(api_error) => api_error.into_response()
        }
    }
}

#[derive(Serialize)]
struct ApiTestResponse {
    pub n: u32,
}

async fn api_test(State(state): State<Arc<CallCounter>>) -> ApiResult<ApiTestResponse> {
    ApiResult::Success(ApiTestResponse { n: state.n.fetch_add(1, Ordering::AcqRel) })
}

struct CallCounter {
    pub n: AtomicU32,
}

pub async fn serve(http_addr: SocketAddr) {
    let app = Router::new()
        .route("/api/test", get(api_test))
        .with_state(Arc::new(CallCounter { n: 0.into() }));
    
    let listener = tokio::net::TcpListener::bind(http_addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
