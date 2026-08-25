def build_plan_prompt(
    destination: str,
    date: str,
    start_time: str,
    start_label: str,
) -> str:

    return f"""
Create a one-day travel itinerary.

Destination: {destination}
Date: {date}
Starting time: {start_time}
Starting location: {start_label}

Generate suitable tourist stops for this day.

For every stop provide:
- position
- title
- category_label
- image_url
- planned_arrival
- planned_departure
- travel_minutes
- stay_minutes

IMPORTANT:
- planned_arrival MUST be a complete ISO 8601 datetime.
- planned_departure MUST be a complete ISO 8601 datetime.
- Travel time between stops must be realistic.
- Travel minutes must be greater than 0
- Stay minutes must be greater than 0
- Give some random Image URL for each stop, but it canno't be empty.
- Use the provided date: {date}.
- Use the local timezone for the destination.
- Do NOT return only HH:MM.
- Example: {date}T09:00:00+05:30

The response must have exactly this structure:

{{
    "stops": [
        {{
            "position": 1,
            "title": "Example Place",
            "category_label": "Sightseeing",
            "image_url": "",
            "planned_arrival": "{date}T09:00:00+05:30",
            "planned_departure": "{date}T10:00:00+05:30",
            "travel_minutes": 20,
            "stay_minutes": 60
        }}
    ]
}}

Return only the required structured data.
"""