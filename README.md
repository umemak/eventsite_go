# eventsite_go

```mermaid
graph TB

user --> WebApp

WebApp -- :8081 --> API

subgraph backend
  developer --> DDL
  developer --> SQL
  developer --> OpenAPI.yml
  OpenAPI.yml -.-> openapi-generator
  SQL -.-> sqlc
  DDL -.-> sqlc
  sqlc -.-> API
  admin -- :8080 --> adminer
  API -- :3306 --> MySQL
  adminer -- :3306 --> MySQL
  subgraph docker-compose
    MySQL
    API
    adminer
  end
end

subgraph frontend
  FEdeveloper --> WebApp
end

openapi-generator -. server .-> API
openapi-generator -. client .-> WebApp

DDL --> MySQL
```

## Generate OpenAPI Server

```sh
docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/openapi.yml \
  -g go-server \
  --additional-properties=router=chi \
  -o /local/out
```

## Generate sqlc

```sh
sqlc generate
```
