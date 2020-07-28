use actix_cors::Cors;
use actix_web::{http, App, HttpServer};

pub mod error_resource;

// todo need to update cors to allow for outside origins to make requests
pub async fn start_resources() -> std::io::Result<()> {
  HttpServer::new(|| {
    App::new()
      .wrap(
        Cors::new() // <- Construct CORS middleware builder
          // .allowed_origin(AllOrSome::All)
          .send_wildcard()
          .allowed_methods(vec!["GET", "POST", "OPTIONS"])
          .allowed_headers(vec![http::header::AUTHORIZATION, http::header::ACCEPT])
          .allowed_header(http::header::CONTENT_TYPE)
          .max_age(3600)
          .finish(),
      )
      .service(error_resource::receive_error)
  })
  .bind("127.0.0.1:8088")?
  .run()
  .await
}
