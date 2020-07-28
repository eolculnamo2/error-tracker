use actix_web::web::Json;
use actix_web::{get, post, web, HttpResponse, Responder};
use serde::Deserialize;
use std::error::Error;

use crate::services::incoming_error;
use crate::structs::error_structs;

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct ErrorBody {
  msg: String,
  url: String,
  line_no: String,
  column_no: String,
  err: String,
  website: String,
  team_id: String,
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct RetreiveByTeamId {
  team_id: String,
}

#[post("/receive-error")]
pub async fn receive_error(error_body: web::Json<ErrorBody>) -> impl Responder {
  let new_error = &error_structs::WebError {
    msg: error_body.msg.clone(),
    url: error_body.url.clone(),
    line_no: error_body.line_no.clone(),
    column_no: error_body.column_no.clone(),
    err: error_body.err.clone(),
    website: error_body.website.clone(),
    team_id: error_body.team_id.clone(),
  };
  incoming_error::save_error(new_error).await;
  HttpResponse::Ok()
}

// #[get("/retreive-errors/")]
// pub async fn retreive_errors(req: web::Json<RetreiveByTeamId>) -> Result<Json<ErrorBody>, Error> {
//   let results = incoming_error::get_errors(&req.team_id);
//   // HttpResponse::Ok().body(results)
//   Ok(Json(results))
// }
