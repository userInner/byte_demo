namespace go api

struct Request {
  1: string message
  2: i64 a
  3: i64 b
}

struct Response {
  1: string message
  2: i64 sum
}

service Echo {
    Response echo(1: Request req)
}
