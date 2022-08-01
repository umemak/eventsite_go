# eventsite_go

```mermaid
graph TB

user --> WebApp

admin -- :8080 --> adminer

WebApp -- :8081 --> API

subgraph Docker-compose
  API -- :3306 --> MySQL
  adminer -- :3306 --> MySQL
end

subgraph dev
  developer --> DDL
  developer --> SQL
  developer --> OpenAPI.yml
end

OpenAPI.yml -- server --> API
OpenAPI.yml -- client --> WebApp
SQL -- sqlc --> API
DDL -- sqlc --> API
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
