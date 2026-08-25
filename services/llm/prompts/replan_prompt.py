def build_replan_prompt(
    destination: str,
    date: str,
    start_time: str,
    start_label: str,
    existing_plan_stops: list,
) -> str:

    return f"""
Replan a one-day travel itinerary.

Destination: {destination}
Date: {date}
Starting time: {start_time}
Starting location: {start_label}

Here is the existing itinerary:

{existing_plan_stops}

Create an improved/replanned itinerary based on the existing plan.

IMPORTANT:
- Use the existing itinerary as the starting point.
- Reorganize, replace, add, or remove stops when necessary to create a better itinerary.
- Keep suitable existing stops when possible.
- Make sure the itinerary is realistic and follows the available time.
- The first stop should be reachable from the starting location.
- Travel time between stops must be realistic.
- Travel minutes must be greater than 0.
- Stay minutes must be greater than 0.
- Every stop must have a non-empty image_url.
- planned_arrival MUST be a complete ISO 8601 datetime.
- planned_departure MUST be a complete ISO 8601 datetime.
- Use the provided date: {date}.
- Use the local timezone for the destination.
- Do NOT return only HH:MM.
- Make sure planned_arrival and planned_departure are chronologically ordered.
- Do not create overlapping stops.
- The replanned itinerary should represent a meaningfully improved version of the existing itinerary.

For every stop provide:
- position
- title
- category_label
- image_url
- planned_arrival
- planned_departure
- travel_minutes
- stay_minutes

Example datetime:
{date}T09:00:00+05:30

The response must have exactly this structure:

{{
    "stops": [
        {{
            "position": 1,
            "title": "Example Place",
            "category_label": "Sightseeing",
            "image_url": "https://example.com/image.jpg",
            "planned_arrival": "{date}T09:00:00+05:30",
            "planned_departure": "{date}T10:00:00+05:30",
            "travel_minutes": 20,
            "stay_minutes": 60
        }}
    ]
}}

Return only the required structured data.
"""