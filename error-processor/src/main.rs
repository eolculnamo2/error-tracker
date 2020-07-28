pub mod dao;
pub mod resources;
pub mod services;
pub mod structs;

#[actix_rt::main]
async fn main() {
    println!("Starting server...");
    let result = resources::start_resources().await;
    match result {
        Ok(_) => println!("Error processing server started..."),
        Err(error) => panic!("Failed to start error processing server: {}", error),
    }
}
