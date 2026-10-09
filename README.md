# Kosmos

Kosmos is the orchestrator of the Lykeion backend. The gateway talks only to Kosmos, and Kosmos combines
Aegis (tokens), Demes (users) and Lexora (dictionary) to answer each request.

## Running the whole backend locally

`deploy/docker-compose.yml` starts both Postgres databases, RabbitMQ and all six Go services
(aegis, demes, lexora-dictionary, lexora-controller, kosmos, gateway). Check the service repositories
out next to this one, then run from this repository:

```bash
docker compose -f deploy/docker-compose.yml up --build
```

The REST API is then served on http://localhost:8080/v1 (`GET /v1/ping` answers `Pong`).
Repositories in another location can be pointed to with `AEGIS_DIR`, `DEMES_DIR`, `GATEWAY_DIR`,
`LEXORA_CONTROLLER_DIR` and `LEXORA_DICTIONARY_DIR`. Container settings live in `deploy/config/`.

New accounts get role 1. Editing the dictionary needs role 3, which for now is set in the database:

```bash
docker compose -f deploy/docker-compose.yml exec postgres-demes \
  psql -U postgres -d demes -c "update lykeion_users set role=3 where email_address='you@example.com'"
```

To run a single service outside Docker instead, start the infrastructure only
(`docker compose -f deploy/docker-compose.yml up postgres-lexora postgres-demes rabbitmq`) and use
`make run-dev` in the service repository; `dev_config.yaml` / `config.yaml` point at localhost.

## Working on go-shared

The services depend on a published version of `github.com/Lykeion-org/go-shared`. To build against a local
checkout while you change it, create a `go.work` in the folder that contains the repositories:

```bash
go work init ./kosmos ./go-shared   # add the other services you are working on
```
