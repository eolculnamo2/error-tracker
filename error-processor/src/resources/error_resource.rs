use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

use crate::services::incoming_error;
use crate::structs::error_structs;

#[derive(Deserialize)]
pub struct ErrorBody {
  msg: String,
  url: String,
  lineNo: String,
  columnNo: String,
  err: String,
}

#[post("/receive-error")]
pub async fn receive_error(error_body: web::Json<ErrorBody>) -> impl Responder {
  let new_error = &error_structs::WebError {
    msg: error_body.msg.clone(),
    url: error_body.url.clone(),
    lineNo: error_body.lineNo.clone(),
    columnNo: error_body.columnNo.clone(),
    err: error_body.err.clone(),
  };
  incoming_error::save_error(new_error).await;
  HttpResponse::Ok()
}
