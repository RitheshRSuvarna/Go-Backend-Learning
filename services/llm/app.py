from pathlib import Path
from urllib import response
from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, ConfigDict, Field
from uuid import UUID
from datetime import datetime
from prompts.plan_prompt import build_plan_prompt
from llm.client import LLMClient
import json

PROJECT_ROOT = Path(__file__).resolve().parents[2]
load_dotenv(PROJECT_ROOT / ".env")


class StrictModel(BaseModel):
    model_config = ConfigDict(strict=True, extra="forbid")


class PlanRequest(StrictModel):
    day_session_id: UUID = Field(strict=False)
    destination: str = Field(min_length=1)
    date: str = Field(min_length=1)
    start_time: str = Field(min_length=1)
    start_label: str = Field(min_length=1)

class Stop(StrictModel):
    position: int = Field(gt=0)
    title: str = Field(min_length=1)
    category_label: str = Field(min_length=1)
    image_url: str
    planned_arrival: datetime = Field(strict=False)
    planned_departure: datetime = Field(strict=False)
    travel_minutes: int = Field(ge=0)
    stay_minutes: int = Field(gt=0)


class PlanResponse(StrictModel):
    stops: list[Stop] = Field(min_length=1)


app = FastAPI(
    title="Trip Planner LLM Service",
    version="1.0.0",
)


@app.get("/")
def health() -> dict[str, str]:
    return {"status": "running"}


llm_client = LLMClient()


@app.post("/plan", response_model=PlanResponse)
def generate_plan(request: PlanRequest) -> PlanResponse:

    prompt = build_plan_prompt(
        destination=request.destination,
        date=request.date,
        start_time=request.start_time,
        start_label=request.start_label,
    )

    try:
        response = llm_client.generate(prompt)

    except Exception as exc:
        raise HTTPException(
            status_code=502,
            detail=f"LLM request failed: {exc}",
        ) from exc

    print("RAW LLM RESPONSE:")
    print(response)
    
    try:
        data = json.loads(response)

    except json.JSONDecodeError as exc:
        raise HTTPException(
            status_code=502,
            detail=f"LLM returned invalid JSON: {exc}",
        ) from exc

    try:
        plan = PlanResponse.model_validate(data)

    except Exception as exc:
        raise HTTPException(
            status_code=502,
            detail=f"LLM returned invalid plan response: {exc}",
        ) from exc

    return plan