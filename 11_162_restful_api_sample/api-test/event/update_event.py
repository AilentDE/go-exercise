import requests
from datetime import datetime, timezone

url = "http://localhost:8080/events"
event_id = 1

body = {
    "name": "Test event",
    "description": "A test event",
    "location": "A location",
    "datetime": datetime.now(timezone.utc).isoformat().replace("+00:00", "Z")
}
r = requests.put(f"{url}/{event_id}", json=body)
# r = requests.get(url)
print(r.status_code, "-", r.json())
