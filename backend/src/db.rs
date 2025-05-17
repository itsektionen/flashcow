use serde::Serialize;
use sqlx::*;

pub type Pool = sqlx::Pool<Postgres>;

pub fn create_pool(connection_string: &str) -> Pool {
    PgPool::connect_lazy(connection_string).unwrap()
}

pub enum Error {
    Other(sqlx::Error),
}

#[derive(FromRow, Serialize)]
pub struct CommitteeRecord {
    pub id: i32,
    pub full_name: String,
    pub short_name: String,
}

pub async fn get_all_committees<'e, E>(executor: E) -> Result<Vec<CommitteeRecord>, sqlx::Error>
where E: Executor<'e, Database = Postgres> {
    sqlx::query_as!(
        CommitteeRecord,
        "
            SELECT *
            FROM committee;
        "
    )
    .fetch_all(executor)
    .await
}
