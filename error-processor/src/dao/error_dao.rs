use bson::doc;
use mongodb::Client;

use crate::structs::error_structs;

const WEB_ERRORS_COLLECTION: &str = "web-errors";
const ERROR_TRACKER_DB: &str = "error-tracker";
// todo move to env
const CONNECTION_STR: &str =
  "mongodb://rob2:test123@ds017195.mlab.com:17195/error-tracker?retryWrites=false";

pub async fn save_error(error: &error_structs::WebError) {
  // todo, connect on app start and save mongo client in struct or something
  let client = Client::with_uri_str(CONNECTION_STR).await.unwrap();
  let coll = client
    .database(ERROR_TRACKER_DB)
    .collection(WEB_ERRORS_COLLECTION);
  let result = coll
    .insert_one(
      doc! {
        "msg": error.msg.clone(),
        "url": error.url.clone(),
        "lineNo": error.lineNo.clone(),
        "columnNo": error.columnNo.clone(),
        "err": error.err.clone(),
      },
      None,
    )
    .await;
  match result {
    Ok(result) => println!("Saved new error: {} --  {:#?}", error.msg, result),
    Err(result) => println!(
      "MongoDB Error: Failed to save msg {} {:#?}",
      error.msg, result
    ),
  }
}
