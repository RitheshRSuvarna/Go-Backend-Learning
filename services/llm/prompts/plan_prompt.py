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

Return only the required structured data.
"""