use bson::doc;
use mongodb::options::FindOptions;
use mongodb::{Client, Collection, Cursor};

use crate::structs::error_structs;

const WEB_ERRORS_COLLECTION: &str = "web-errors";
const ERROR_TRACKER_DB: &str = "error-tracker";
// todo move to env
const CONNECTION_STR: &str =
  "mongodb://rob2:test123@ds017195.mlab.com:17195/error-tracker?retryWrites=false";

pub async fn save_error(error: &error_structs::WebError) {
  let coll = get_web_errors_collection().await;
  let result = coll
    .insert_one(
      doc! {
        "msg": error.msg.clone(),
        "url": error.url.clone(),
        "lineNo": error.line_no.clone(),
        "columnNo": error.column_no.clone(),
        "err": error.err.clone(),
        "teamId": error.team_id.clone(),
        "website": error.website.clone(),
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

pub async fn get_errors_by_team_id(team_id: &String) -> Cursor {
  let coll = get_web_errors_collection().await;
  let filter = doc! { "teamId": team_id };

  let filter_options = FindOptions::builder().sort(doc! { "name": "a" }).build();
  let result = coll.find(filter, filter_options).await;
  match result {
    Ok(result) => result,
    Err(result) => panic!("Could not find results for team_id {}", team_id),
  }
}

async fn get_web_errors_collection() -> Collection {
  // todo, connect on app start and save mongo client in struct or something
  let client = Client::with_uri_str(CONNECTION_STR).await.unwrap();
  let coll = client
    .database(ERROR_TRACKER_DB)
    .collection(WEB_ERRORS_COLLECTION);
  coll
}
