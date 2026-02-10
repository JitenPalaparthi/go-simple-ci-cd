# Go Simple CI/CD (GitHub Actions) — with Unit Tests

A minimal Go HTTP service with:
- **Unit tests** (net/http/httptest)
- **CI**: gofmt check, `go vet`, `go test -race`, coverage, `golangci-lint`
- **CD**: build & push Docker image to **GitHub Container Registry (GHCR)** on version tags (`vX.Y.Z`)

## Run locally

```bash
go run ./cmd/server
curl -s localhost:8080/health
```

## Run tests

```bash
go test ./... -race -count=1
make coverage
```

## Run with Docker

```bash
docker build -t go-simple-ci-cd:local .
docker run --rm -p 8080:8080 go-simple-ci-cd:local
curl -s localhost:8080/health
```

## CI

Workflow: `.github/workflows/ci.yml`
- Runs on PRs and pushes to `main`
- Fails if any file is not `gofmt`-formatted
- Runs tests + race detector
- Prints coverage summary
- Runs golangci-lint

## CD → GHCR

Workflow: `.github/workflows/cd.yml`
- Triggers on tags like `v1.0.0`
- Pushes:
  - `ghcr.io/<owner>/<repo>:v1.0.0`
  - `ghcr.io/<owner>/<repo>:latest`

Release:

```bash
git tag v1.0.0
git push origin v1.0.0
```
