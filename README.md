# AstraERP — Modular ERP Backend (v9)

[![Go](https://img.shields.io/badge/Go-1.20-blue?logo=go&logoColor=white)](https://golang.org/) [![GORM](https://img.shields.io/badge/GORM-1.x-0f172a?style=flat&logo=go&logoColor=white)](https://gorm.io/) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue?logo=postgresql&logoColor=white)](https://www.postgresql.org/) [![Swagger](https://img.shields.io/badge/Swagger-swag-blue)](https://github.com/swaggo/swag) [![Docker](https://img.shields.io/badge/Docker-24.0-blue?logo=docker&logoColor=white)](https://www.docker.com/) [![Redis](https://img.shields.io/badge/Redis-6.2-orange?logo=redis&logoColor=white)](https://redis.io/) [![Gin](https://img.shields.io/badge/Gin-Gonic-00ADD8?logo=gin&logoColor=white)](https://github.com/gin-gonic/gin)

AstraERP is a production-ready, modular backend for enterprise resource planning (ERP) systems. Built with Go, Gin and GORM, AstraERP is designed around clean architecture, testability and operational readiness — Postgres/PostGIS compatible.

Version: v9 — polished, comprehensive README with folder structure, environment guidelines, Docker & deployment notes, and best practices.

---

## Table of Contents

- [AstraERP — Modular ERP Backend (v9)](#astraerp--modular-erp-backend-v9)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Key Features](#key-features)
  - [Project Structure (short)](#project-structure-short)
  - [Quick Start](#quick-start)
  - [Environment \& Secrets Management](#environment--secrets-management)
  - [Database \& GIS](#database--gis)
  - [API Documentation (Swagger)](#api-documentation-swagger)
  - [Authentication \& Security](#authentication--security)
  - [Docker \& Deployment](#docker--deployment)
  - [Testing \& CI/CD](#testing--cicd)
  - [Developer Tools \& Best Practices](#developer-tools--best-practices)
  - [Changelog (v9)](#changelog-v9)
  - [Author \& Contact](#author--contact)
  - [License](#license)

---

## Overview

AstraERP implements a clean, modular architecture where each domain (authentication, users, geofencing, etc.) follows the same pattern:

DTO → Repository (DAO) → Service → HTTP Handler

This separation of concerns improves maintainability, testability and enables independent module evolution.

## Key Features

- Modular domain-based architecture  
- DAO / Repository pattern for testable data access  
- JWT-based authentication and permission middleware  
- Geofencing & GIS-ready data models (PostGIS-ready)  
- Auto-generated Swagger documentation (swaggo)  
- Consistent API envelope (shared utils.APIResponse)  
- Docker-ready and designed for production deployments

## Project Structure (short)

A concise, opinionated tree to show where responsibilities live:

```text
ERP-SYSTEM/
├── cmd/                 # Application entry point(s)
│   └── main.go
├── config/              # Configuration loader & DB migrations
│   └── migrations/
├── internal/
│   ├── dto/             # Request & response DTOs (per module)
│   ├── http/
│   │   ├── handlers/    # HTTP handlers / controllers
│   │   ├── middlewares/ # JWT, permissions, logging, CORS
│   │   ├── routes/      # Route registration (versioned)
│   │   └── docs/        # Swagger/docs (generated)
│   ├── models/          # GORM models (Postgres/PostGIS-ready)
│   ├── repository/      # DAO interfaces + implementations
│   ├── services/        # Business logic
│   └── utils/           # Shared helpers (responses, logger, id parser)
├── pkg/                 # Reusable exported packages (optional)
├── scripts/             # Dev / migration / seed scripts
├── deployments/         # Docker / k8s manifests & examples
├── tmp/                 # Temporary files
├── .env                 # Local env (DO NOT commit secrets)
├── .gitignore
├── Makefile
├── go.mod
└── go.sum
```

## Quick Start

1. Clone
```bash
git clone https://github.com/engrsakib/astraERP-GO-Backend.git
cd astraERP-GO-Backend
```

2. Install dependencies
```bash
go mod tidy
```

3. Environment
- Copy `.env.example` to `.env` and fill values or inject secrets from your secret store.
```bash
cp .env.example .env
```

4. Generate Swagger (developer machine)
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/main.go -o internal/http/docs
```

5. Run (development)
```bash
# run directly
go run cmd/main.go

# or use live-reload (Air)
air
```

Open: `http://localhost:8080` — Swagger: `http://localhost:8080/swagger/index.html`

## Environment & Secrets Management

- Use `.env.example` as template; never commit secrets. Add `.env` to `.gitignore`.
- For production, prefer a secrets manager (AWS Secrets Manager, HashiCorp Vault, GCP Secret Manager, or Kubernetes Secrets).
- Generate strong secrets for `JWT_SECRET` (e.g., `openssl rand -hex 32`).
- Recommended env variables (examples are provided in `.env.example`, `.env.local`, `.env.production` in the repo).

## Database & GIS

- Primary DB: PostgreSQL. Models and queries are written with Postgres compatibility in mind.
- Geospatial: PostGIS-ready models (geometry fields) are included for geofencing features.
- Migrations: store SQL/migration files in `config/migrations/` and use a migration tool (e.g., golang-migrate).
- Ensure PostGIS extension is enabled on DB when using geometry columns:
```sql
CREATE EXTENSION IF NOT EXISTS postgis;
```

## API Documentation (Swagger)

- Routes are documented with swag annotations.
- Generate docs:
```bash
swag init -g cmd/main.go -o internal/http/docs
```
- The Swagger UI is available at `/swagger/index.html` once the server runs.

## Authentication & Security

- Authentication: JWT access tokens (configurable expirations).
- Authorization: permission-based middleware for endpoint-level access.
- Best practices:
  - Use HTTPS in production.
  - Set `COOKIE_SECURE=true` for secure cookies.
  - Rotate secrets and use short-lived tokens where appropriate.

## Docker & Deployment

- `deployments/docker/` contains sample Dockerfile and docker-compose.yml. Replace placeholders and avoid embedding secrets in images.
- Build a static binary for small container images:
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd
docker build -t astraerp:latest .
```
- Use multi-stage builds and minimal base images.
- For production, use orchestrators (Kubernetes) and inject secrets from the cluster secret store.

## Testing & CI/CD

- Unit-test business logic (services). Mock repository interfaces for isolated tests.
- Integration tests: use a test database (or testcontainers) for repository tests.
- Recommended CI tasks:
  - `go test ./...`
  - `golangci-lint run`
  - Build artifacts / container images
  - Deploy to staging via pipeline (GitHub Actions, GitLab CI, etc.)

## Developer Tools & Best Practices

- Live reload: [Air](https://github.com/cosmtrek/air)  
- Linter: [golangci-lint](https://golangci-lint.run)  
- Use Docker Compose for local Postgres/Redis during development  
- Use `openssl rand -hex 32` to generate secrets  
- Keep repository lean: do not commit `.env` or other secrets

## Changelog (v9)

- README content refined and reorganized for clarity and operational guidance.
- Added explicit Project Structure section.
- Included environment templates and security guidance.
- Added Docker tips, Swagger instructions, and testing/CI notes.

## Author & Contact

Md. Nazmus Sakib — Backend Engineer  
Repository: https://github.com/engrsakib/astraERP-GO-Backend  
Website / Portfolio: https://engrsakib.com  
Contact: info@engrsakib.com

## License

Proprietary — maintained by Md. Nazmus Sakib. For licensing or commercial use, please contact the author.

---

If you like, I can also:
- Provide a production-ready `Dockerfile` and `docker-compose.yml` tailored to this repo,  
- Add a GitHub Actions workflow (build, lint, test, docker build), or  
- Scaffold a fully working sample module (geofence) including handler → service → repository → tests.

Tell me which item you'd like next and I will prepare it.