use std::{net::SocketAddr, sync::Arc};
use axum::{body::Body, extract::State, http::{HeaderValue, StatusCode}, response::IntoResponse, routing::get, Router};
use serde::Serialize;
use serde_json::json;
use crate::db;

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

async fn get_all_committees(State(ctx): State<Arc<Context>>) -> ApiResult<Vec<db::CommitteeRecord>> {
    let all_committees = db::get_all_committees(&ctx.db_pool).await.unwrap();
    ApiResult::Success(all_committees)
}

struct Context {
    db_pool: db::Pool,
}

pub async fn serve(http_addr: SocketAddr, pool: db::Pool) {
    let app = Router::new()
        .route("/api/committees", get(get_all_committees))
        .with_state(Arc::new(Context { db_pool: pool }));
    
    let listener = tokio::net::TcpListener::bind(http_addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
