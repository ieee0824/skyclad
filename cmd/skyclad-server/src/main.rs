#[macro_use] extern crate nickel;
use nickel::{Nickel, HttpRouter};

fn main() {
    let mut serv = Nickel::new();

    serv.get("/bar", middleware!("This is the /bar handler"));
    serv.get("/user/:userid", middleware! { |request|
      format!("<h1>This is user: {:?}</h1>", request.param("userid").unwrap())
    });
    serv.get("/a/*/d", middleware!("matches /a/b/d but not /a/b/c/d"));
    serv.get("/a/**/d", middleware!("This matches /a/b/d and also /a/b/c/d"));
    serv.listen("127.0.0.1:6767");
}