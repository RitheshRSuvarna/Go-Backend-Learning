# Demo Script Template (Day30)

## 1) Start dependencies

```powershell
docker compose up -d
```

## 2) Run API

```powershell
cd services/api
go run .
```

## 3) Run LLM service

```powershell
cd services/llm
python -m venv .venv
.\.venv\Scripts\activate
pip install -r requirements.txt
uvicorn app:app --reload --port 8090
```

## 4) Open Swagger

- `http://localhost:8080/swagger/index.html`

## 5) Execute story

```powershell
. .\Rithesh-backend-docs\scripts\01_auth_signup_or_login.ps1
. .\Rithesh-backend-docs\scripts\02_create_trip.ps1
. .\Rithesh-backend-docs\scripts\03_create_day_session.ps1
. .\Rithesh-backend-docs\scripts\04_create_plan_v1.ps1
. .\Rithesh-backend-docs\scripts\05_replan.ps1
. .\Rithesh-backend-docs\scripts\06_diff_plans.ps1
. .\Rithesh-backend-docs\scripts\07_accept_plan.ps1
```

## 6) Demo talking points

- What problem this backend solves
- Which parts are deterministic rules vs LLM
- What is production-ready vs training-only
