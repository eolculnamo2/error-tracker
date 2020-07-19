use actix_web::{App, HttpServer};

pub mod error_resource;

pub async fn start_resources() -> std::io::Result<()> {
  HttpServer::new(|| App::new().service(error_resource::receive_error))
    .bind("127.0.0.1:8088")?
    .run()
    .await
}
