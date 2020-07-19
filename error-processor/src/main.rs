pub mod dao;
pub mod resources;
pub mod services;
pub mod structs;

#[actix_rt::main]
async fn main() {
    println!("Starting server...");
    resources::start_resources().await;
}
