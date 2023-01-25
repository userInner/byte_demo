namespace go api

struct Request {
  1: string mesge
}

struct Response {
  1: string message
}

service Echo {
    Response echo(1: Request req)
}
