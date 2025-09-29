use crate::db::{self, GetUserError};
use axum::{
    body::Body,
    extract::{Path, State},
    http::{HeaderValue, StatusCode},
    response::IntoResponse,
    routing::{get, post},
    Json, Router,
};
use serde::{Deserialize, Serialize};
use std::{net::SocketAddr, sync::Arc};

enum ApiError {
    Other,
    JsonError(serde_json::Error),
    DuplicateCommittee,
    CommitteeNotFound,
    UserNotFound,
}

impl IntoResponse for ApiError {
    fn into_response(self) -> axum::response::Response {
        let (status_code, error_code, debug_msg) = match self {
            ApiError::Other => (StatusCode::INTERNAL_SERVER_ERROR, -1, "Unknown error"),
            ApiError::JsonError(_json_error) => {
                (StatusCode::INTERNAL_SERVER_ERROR, 1, "JSON error")
            }
            ApiError::DuplicateCommittee => (StatusCode::CONFLICT, 2, "Duplicate committee name"),
            ApiError::CommitteeNotFound => (StatusCode::NOT_FOUND, 3, "Committee not found"),
            ApiError::UserNotFound => (StatusCode::NOT_FOUND, 4, "User not found"),
        };

        let json_obj = serde_json::json!({
            "error_code": error_code,
            "debug_msg": debug_msg
        });

        if let Ok(json) = serde_json::to_string(&json_obj) {
            let mut response = axum::response::Response::new(Body::from(json));
            *response.status_mut() = status_code;
            response.headers_mut().insert(
                "Content-Type",
                HeaderValue::from_str("application/json").unwrap(),
            );

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

impl<R> IntoResponse for ApiResult<R>
where
    R: Serialize,
{
    fn into_response(self) -> axum::response::Response {
        match self {
            ApiResult::Success(response_object) => match serde_json::to_string(&response_object) {
                Err(error) => ApiError::JsonError(error).into_response(),
                Ok(response_string) => {
                    let mut response = axum::response::Response::new(Body::from(response_string));
                    response.headers_mut().insert(
                        "Content-Type",
                        HeaderValue::from_str("application/json").unwrap(),
                    );

                    response.into_response()
                }
            },
            ApiResult::Error(api_error) => api_error.into_response(),
        }
    }
}

fn handle_sqlx_error<R>(error: sqlx::Error) -> ApiResult<R> {
    log::error!("Got sqlx error: {:?}", error);
    ApiResult::Error(ApiError::Other)
}

async fn get_all_committees(
    State(ctx): State<Arc<Context>>,
) -> ApiResult<Vec<db::CommitteeRecord>> {
    match db::get_all_committees(&ctx.db_pool).await {
        Ok(all_committees) => ApiResult::Success(all_committees),
        Err(sqlx_error) => handle_sqlx_error(sqlx_error),
    }
}

#[derive(Deserialize)]
struct AddCommitteeRequest {
    pub full_name: String,
    pub short_name: String,
}

async fn add_committee(
    State(ctx): State<Arc<Context>>,
    Json(request): Json<AddCommitteeRequest>,
) -> ApiResult<Vec<db::CommitteeRecord>> {
    match db::add_committee(&ctx.db_pool, &request.full_name, &request.short_name).await {
        Ok(results) => ApiResult::Success(results),
        Err(db::AddCommitteeError::Duplicate) => ApiResult::Error(ApiError::DuplicateCommittee),
        Err(db::AddCommitteeError::Other(sqlx_error)) => handle_sqlx_error(sqlx_error),
    }
}

#[derive(Deserialize)]
struct RenameCommitteeRequest {
    pub id: i32,
    pub new_full_name: String,
    pub new_short_name: String,
}

async fn rename_committee(
    State(ctx): State<Arc<Context>>,
    Json(request): Json<RenameCommitteeRequest>,
) -> ApiResult<Vec<db::CommitteeRecord>> {
    match db::rename_committee(
        &ctx.db_pool,
        request.id,
        &request.new_full_name,
        &request.new_short_name,
    )
    .await
    {
        Ok(results) => ApiResult::Success(results),
        Err(db::RenameCommitteeError::NotFound) => ApiResult::Error(ApiError::CommitteeNotFound),
        Err(db::RenameCommitteeError::Duplicate) => ApiResult::Error(ApiError::DuplicateCommittee),
        Err(db::RenameCommitteeError::Other(sqlx_error)) => handle_sqlx_error(sqlx_error),
    }
}

#[derive(Deserialize)]
struct DeleteCommitteeRequest {
    pub id: i32,
}

async fn delete_committee(
    State(ctx): State<Arc<Context>>,
    Json(request): Json<DeleteCommitteeRequest>,
) -> ApiResult<Vec<db::CommitteeRecord>> {
    match db::delete_committee(&ctx.db_pool, request.id).await {
        Ok(results) => ApiResult::Success(results),
        Err(db::DeleteCommitteeError::NotFound) => ApiResult::Error(ApiError::CommitteeNotFound),
        Err(db::DeleteCommitteeError::Other(sqlx_error)) => handle_sqlx_error(sqlx_error),
    }
}

async fn get_specific_user(
    State(ctx): State<Arc<Context>>,
    Path(user_id): Path<i32>,
) -> ApiResult<db::UserRecord> {
    match db::get_user_from_id(&ctx.db_pool, user_id).await {
        Ok(user_record) => ApiResult::Success(user_record),
        Err(GetUserError::NotFound) => ApiResult::Error(ApiError::UserNotFound),
        Err(GetUserError::Other(sqlx_error)) => handle_sqlx_error(sqlx_error),
    }
}

struct Context {
    db_pool: db::Pool,
}

pub async fn serve(http_addr: SocketAddr, pool: db::Pool) {
    let app = Router::new()
        .route("/api/committees", get(get_all_committees))
        .route("/api/committees", post(add_committee))
        .route("/api/rename_committees", post(rename_committee))
        .route("/api/delete_committee", post(delete_committee))
        .route("/api/user/{user_id}", get(get_specific_user))
        .with_state(Arc::new(Context { db_pool: pool }));

    let listener = tokio::net::TcpListener::bind(http_addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
