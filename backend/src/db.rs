use serde::Serialize;
use sqlx::*;
use crate::util::contains_duplicates;

pub type Pool = sqlx::Pool<Postgres>;

pub fn create_pool(connection_string: &str) -> Pool {
    PgPool::connect_lazy(connection_string).unwrap()
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
            SELECT id, full_name, short_name
            FROM committee
            WHERE deleted IS NULL;
        "
    )
    .fetch_all(executor)
    .await
}

#[derive(Debug)]
pub enum AddCommitteeError {
    Duplicate,
    Other(sqlx::Error),
}

impl From<sqlx::Error> for AddCommitteeError {
    fn from(value: sqlx::Error) -> Self {
        Self::Other(value)
    }
}

pub async fn add_committee(pool: &Pool, full_name: &str, short_name: &str) -> Result<Vec<CommitteeRecord>, AddCommitteeError> {
    let mut tx = pool.begin().await?;
    sqlx::query!(
        "
            INSERT INTO committee (full_name, short_name)
            VALUES ($1, $2);
        ",
        full_name,
        short_name
    )
    .execute(&mut *tx)
    .await?;

    let results = get_all_committees(&mut *tx).await?;

    if contains_duplicates(results.iter().map(|cr| &cr.full_name)) {
        return Err(AddCommitteeError::Duplicate);
    }

    if contains_duplicates(results.iter().map(|cr| &cr.short_name)) {
        return Err(AddCommitteeError::Duplicate);
    }

    tx.commit().await?;

    Ok(results)
}

#[derive(Debug)]
pub enum RenameCommitteeError {
    NotFound,
    Duplicate,
    Other(sqlx::Error)
}

impl From<sqlx::Error> for RenameCommitteeError {
    fn from(value: sqlx::Error) -> Self {
        Self::Other(value)
    }
}

pub async fn rename_committee(pool: &Pool, id: i32, new_full_name: &str, new_short_name: &str) -> Result<Vec<CommitteeRecord>, RenameCommitteeError> {
    let mut tx = pool.begin().await?;
    let update_result = sqlx::query!(
        "
            UPDATE committee
            SET full_name = $1,
                short_name = $2
            WHERE id = $3;
        ",
        new_full_name,
        new_short_name,
        id
    )
    .execute(&mut *tx)
    .await?;

    if update_result.rows_affected() == 0 {
        return Err(RenameCommitteeError::NotFound);
    }

    let results = get_all_committees(&mut *tx).await?;

    if contains_duplicates(results.iter().map(|cr| &cr.full_name)) {
        return Err(RenameCommitteeError::Duplicate);
    }

    if contains_duplicates(results.iter().map(|cr| &cr.short_name)) {
        return Err(RenameCommitteeError::Duplicate);
    }

    tx.commit().await?;

    Ok(results)
}

#[derive(Debug)]
pub enum DeleteCommitteeError {
    NotFound,
    Other(sqlx::Error),
}

impl From<sqlx::Error> for DeleteCommitteeError {
    fn from(value: sqlx::Error) -> Self {
        Self::Other(value)
    }
}

pub async fn delete_committee(pool: &Pool, id: i32) -> Result<Vec<CommitteeRecord>, DeleteCommitteeError> {
    let mut tx = pool.begin().await?;

    let result = sqlx::query!(
        "
            UPDATE committee
            SET deleted = NOW()
            WHERE id = $1 AND deleted IS NULL;
        ",
        id
    )
    .execute(&mut *tx)
    .await?;

    if result.rows_affected() == 0 {
        return Err(DeleteCommitteeError::NotFound);
    }

    let committees = get_all_committees(&mut *tx).await?;

    tx.commit().await?;

    Ok(committees)
}
