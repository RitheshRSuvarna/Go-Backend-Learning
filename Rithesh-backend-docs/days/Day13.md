# Day 13 - Python LLM service skeleton (FastAPI, contracts-only)

## Goal (end of day)
- Python service runs locally on port `8090`
- Endpoints exist:
  - `POST /plan`
  - `POST /replan`
- Service does not write to DB (JSON in/out only)

## Learn (45-90 minutes)
- FastAPI tutorial: `https://fastapi.tiangolo.com/tutorial/`
- Pydantic models: `https://docs.pydantic.dev/latest/`
- Uvicorn basics: `https://www.uvicorn.org/`

## Prerequisites
- Day12 complete
- Python 3.11+ available (`python --version`)

## Tasks
### 1) Create folder and files
Create:
- `services/llm/app.py`
- `services/llm/requirements.txt`

`requirements.txt` (minimum):
- `fastapi`
- `uvicorn`

### 2) Define strict request/response models
In `app.py`, define Pydantic models for:
- plan request
- plan response
- replan request
- replan response

Use strict field names and required fields.

### 3) Implement deterministic endpoints
- `POST /plan`: return fixed stop list
- `POST /replan`: return deterministic changed sequence (for example swap first two stops)

No randomness in output.

### 4) Run service
```powershell
cd services/llm
python -m venv .venv
.\.venv\Scripts\activate
pip install -r requirements.txt
uvicorn app:app --reload --port 8090
```

### 5) Manual test
In another terminal:
```powershell
Invoke-RestMethod http://localhost:8090/docs
```

## Acceptance checks
- `http://localhost:8090/docs` loads
- `POST /plan` returns JSON with non-empty `stops`
- `POST /replan` returns deterministic output for same input

## Deliverables
- `services/llm` runs locally and returns valid contract responses.
