use crate::dao::error_dao;
use crate::structs::error_structs;
use mongodb::Cursor;

pub async fn save_error(error: &error_structs::WebError) {
  error_dao::save_error(&error).await;
}

pub async fn get_errors(team_id: &String) -> Cursor {
  error_dao::get_errors_by_team_id(&team_id).await
}
