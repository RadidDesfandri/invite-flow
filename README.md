# Invite Flow Setup

Monorepo project with:

- `frontend`: Next.js (latest) + Tailwind CSS + shadcn/ui + Axios
- `backend`: Golang API + Redis + PostgreSQL
- Docker Compose for local development

## Project Structure

```text
.
├─ frontend/
├─ backend/
└─ docker-compose.yml
```

## Prerequisites

- Docker Desktop
- (Optional, if running without Docker) Node.js 22+ and Go 1.25+

## Run With Docker (Recommended)

From project root:

```bash
docker compose up --build
```

Services:

- Frontend: [http://localhost:3000](http://localhost:3000)
- Backend: [http://localhost:8080](http://localhost:8080)
- Backend Health: [http://localhost:8080/health](http://localhost:8080/health)
- PostgreSQL: `localhost:5432`
- Redis: `localhost:6379`

## Hot Reload

- Frontend runs with `next dev` in Docker.
- Source code is mounted as a volume, so changes in `frontend/` auto-reload.
- Polling env vars are enabled for stable file watching on Windows.

## API Endpoints

- `GET /` -> backend running message
- `GET /health` -> checks PostgreSQL + Redis connectivity
- `GET /api/v1/ping` -> simple ping response

## Local Run Without Docker (Optional)

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Backend

Start PostgreSQL and Redis first, then:

```bash
cd backend
go run main.go
```

Default backend env (if not set):

- `PORT=8080`
- `POSTGRES_URL=postgres://postgres:postgres@postgres:5432/app_db?sslmode=disable`
- `REDIS_ADDR=redis:6379`
- `REDIS_PASSWORD=`

## Stop Services

```bash
docker compose down
```

To also remove volumes:

```bash
docker compose down -v
```
