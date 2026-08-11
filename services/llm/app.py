from fastapi import FastAPI
from pydantic import BaseModel, ConfigDict, Field
from uuid import UUID


class StrictModel(BaseModel):
    model_config = ConfigDict(strict=True, extra="forbid")


class PlanRequest(StrictModel):
    day_session_id: UUID


class Stop(StrictModel):
    position: int = Field(gt=0)
    title: str = Field(min_length=1)
    category_label: str = Field(min_length=1)
    image_url: str
    planned_arrival: str = Field(pattern=r"^(?:[01]\d|2[0-3]):[0-5]\d$")
    planned_departure: str = Field(pattern=r"^(?:[01]\d|2[0-3]):[0-5]\d$")
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


@app.post("/plan", response_model=PlanResponse)
def generate_plan(request: PlanRequest) -> PlanResponse:
    # Deterministic contract implementation. No real AI provider is called yet.
    return PlanResponse(
        stops=[
            Stop(
                position=1,
                title="Bangalore Palace",
                category_label="Sightseeing",
                image_url="",
                planned_arrival="09:00",
                planned_departure="10:30",
                travel_minutes=20,
                stay_minutes=90,
            ),
            Stop(
                position=2,
                title="Cubbon Park",
                category_label="Sightseeing",
                image_url="",
                planned_arrival="11:00",
                planned_departure="12:00",
                travel_minutes=15,
                stay_minutes=60,
            ),
            Stop(
                position=3,
                title="UB City",
                category_label="Dining",
                image_url="",
                planned_arrival="12:30",
                planned_departure="14:00",
                travel_minutes=10,
                stay_minutes=90,
            ),
        ]
    )
