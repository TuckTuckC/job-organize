use actix_web::{get, App, HttpServer, Responder, HttpResponse};
use sqlx::SqlitePool;

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello from Actix + SQLite")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Connect to Database
    let db = SqlitePool::connect("sqlite://app.db").await.unwrap();

    // Start the HTTP server
    HttpServer::new(move || {
        App::new()
            // You would add the db to app state here (omitted for now)
            .service(hello)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}

