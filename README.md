# todo-api

A small **REST + JSON** service in Go: an in-memory task list with mutex-safe storage, standard-library HTTP, and a layout ready to grow (tests, Docker, CI).

Created by Aksel.

[![CI](https://github.com/YOUR_GITHUB_USERNAME/todo-api/actions/workflows/ci.yml/badge.svg)](https://github.com/YOUR_GITHUB_USERNAME/todo-api/actions/workflows/ci.yml)

## Features


- `net/http` (Go 1.22+ route patterns), `encoding/json`
- `internal/` packages: `store` (concurrent in-memory map), `api` (handlers + router)
- `GET /health` for load balancers and containers
- Unit tests (`httptest`, table-free integration-style handler tests)
- Multi-stage **Dockerfile**, **docker compose**, **Makefile**
- **GitHub Actions**: `go vet`, tests with `-race`, build

## Project layout

```
.
├── cmd/server/          # main entrypoint
├── internal/
│   ├── api/             # HTTP handlers and router
│   └── store/           # in-memory task store
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── .github/workflows/   # CI
```

## Requirements

- [Go 1.22+](https://go.dev/dl/) (uses `ServeMux` method and path patterns)

## Quick start

```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/todo-api.git
cd todo-api
go run ./cmd/server
```

Server listens on **`http://localhost:8080`** unless `PORT` is set.

## Configuration

| Variable | Default | Description        |
|----------|---------|--------------------|
| `PORT`   | `8080`  | Listen port (e.g. `3000` → `:3000`) |

## API

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Liveness: `{"status":"ok"}` |
| `GET` | `/tasks` | List tasks (sorted by numeric id) |
| `POST` | `/tasks` | Create task: JSON `{"title":"string"}` (title trimmed, required) |
| `POST` | `/tasks/{id}/toggle` | Toggle `done` |
| `DELETE` | `/tasks/{id}` | Delete task (`204` on success) |

### Examples

```bash
# Health
curl -s http://localhost:8080/health

# Create
curl -s -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Learn Go modules"}'

# List
curl -s http://localhost:8080/tasks

# Toggle done
curl -s -X POST http://localhost:8080/tasks/1/toggle

# Delete
curl -s -o /dev/null -w "%{http_code}\n" -X DELETE http://localhost:8080/tasks/1
```

## Docker

```bash
make docker-build    # image: todo-api:local
make docker-up       # compose: http://localhost:8080
```

Or:

```bash
docker compose up --build
```

## Development

```bash
make test    # race detector
make vet
make build   # writes ./todo-api
```

## Publishing to GitHub

1. Create a new repository on GitHub (empty, no README if you push this tree as-is).
2. Replace the module path in `go.mod` and imports with your repo, for example:
   - `module github.com/<you>/todo-api`
   - Update import paths under `cmd/` and `internal/` to match.
3. Replace `YOUR_GITHUB_USERNAME` in this README (badge + clone URL).
4. Push:

```bash
git init
git add .
git commit -m "Initial commit: todo JSON API"
git branch -M main
git remote add origin https://github.com/<you>/todo-api.git
git push -u origin main
```

## License

Apache License 2.0 — see [LICENSE](LICENSE).
# todo-api
