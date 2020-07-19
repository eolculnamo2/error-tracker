use crate::dao::error_dao;
use crate::structs::error_structs;

pub async fn save_error(error: &error_structs::WebError) {
  error_dao::save_error(&error).await;
}
