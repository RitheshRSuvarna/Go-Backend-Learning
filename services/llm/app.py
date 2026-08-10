from fastapi import FastAPI
from pydantic import BaseModel
from typing import List

class Stop(BaseModel):
    position: int
    title: str
    category_label: str
    image_url: str
    planned_arrival: str
    planned_departure: str
    travel_minutes: int
    stay_minutes: int

class PlanRequest(BaseModel):
    day_session_id: str

class PlanResponse(BaseModel):
    stops: List[Stop]

class ReplanRequest(BaseModel):
    day_session_id: str

class ReplanResponse(BaseModel):
    stops: List[Stop]

app = FastAPI(
    title="Trip Planner LLM Service",
    version="1.0.0"
)

@app.get("/")
def health():
    return {
        "status": "running"
    }

@app.post("/plan", response_model=PlanResponse)
def generate_plan(request: PlanRequest):
    return {
        "stops": [
            {
                "position": 1,
                "title": "Bangalore Palace",
                "category_label": "Sightseeing",
                "image_url": "",
                "planned_arrival": "09:00",
                "planned_departure": "10:30",
                "travel_minutes": 20,
                "stay_minutes": 90
            },
            {
                "position": 2,
                "title": "Cubbon Park",
                "category_label": "Sightseeing",
                "image_url": "",
                "planned_arrival": "11:00",
                "planned_departure": "12:00",
                "travel_minutes": 15,
                "stay_minutes": 60
            },
            {
                "position": 3,
                "title": "UB City",
                "category_label": "Dining",
                "image_url": "",
                "planned_arrival": "12:30",
                "planned_departure": "14:00",
                "travel_minutes": 10,
                "stay_minutes": 90
            }
        ]
    }

@app.post("/replan", response_model=ReplanResponse)
def replan(request: ReplanRequest):
    return {
        "stops": [
            {
                "position": 1,
                "title": "Cubbon Park",
                "category_label": "Sightseeing",
                "image_url": "",
                "planned_arrival": "09:00",
                "planned_departure": "10:00",
                "travel_minutes": 15,
                "stay_minutes": 60
            },
            {
                "position": 2,
                "title": "Bangalore Palace",
                "category_label": "Sightseeing",
                "image_url": "",
                "planned_arrival": "10:30",
                "planned_departure": "12:00",
                "travel_minutes": 20,
                "stay_minutes": 90
            },
            {
                "position": 3,
                "title": "UB City",
                "category_label": "Dining",
                "image_url": "",
                "planned_arrival": "12:30",
                "planned_departure": "14:00",
                "travel_minutes": 10,
                "stay_minutes": 90
            }
        ]
    }