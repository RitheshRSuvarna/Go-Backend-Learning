# README_BACKEND Template

Create this as `README_BACKEND.md` in your workspace root (same level as `Rithesh-backend-docs`) on Day28 and fill all TODOs.

## Backend Runbook

### Services

- Go API: `services/api` (default port `8080`)
- Python LLM service: `services/llm` (default port `8090`)
- Postgres: Docker compose

### Prerequisites

- Docker Desktop
- Go 1.22+
- Python 3.11+

### Environment variables

API (`services/api/.env`):

```env
PORT=8080
DATABASE_URL=postgres://askhub:askhub@localhost:5432/askhubtravel?sslmode=disable
LLM_BASE_URL=http://localhost:8090
```

### Run locally

1. Start DB:

```powershell
docker compose up -d
```

2. Start API:

```powershell
cd services/api
go run .
```

3. Start LLM:

```powershell
cd services/llm
python -m venv .venv
.\.venv\Scripts\activate
pip install -r requirements.txt
uvicorn app:app --reload --port 8090
```

### Smoke checks

```powershell
Invoke-RestMethod http://localhost:8080/health -Method Get
Invoke-RestMethod http://localhost:8090/docs
```

### Scripts

```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
. .\Rithesh-backend-docs\scripts\07_accept_plan.ps1
```

### Troubleshooting

- Use `Rithesh-backend-docs/DEBUG_PLAYBOOK.md`.

### Out of scope in this training

- Production auth hardening (CSRF, secure cookies, rate limiting)
- Real email delivery provider
- Real LLM prompt quality system
- Production infrastructure and deployment
