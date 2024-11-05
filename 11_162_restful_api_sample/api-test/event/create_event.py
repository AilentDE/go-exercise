import requests
from datetime import datetime, timezone

url = "http://localhost:8080/events"

header = {
    "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QwMUBnbWFpbC5jb20iLCJleHAiOjE3MzA4MzMzNjUsInVzZXJJZCI6MX0.TmSv8rXClSCmycpS_3sm4lN3WcwqDH0n9M6aBaNPqrA"
}

body = {
    "name": "Test event",
    "description": "A test event",
    "location": "A location",
    "datetime": datetime.now(timezone.utc).isoformat().replace("+00:00", "Z")
}
r = requests.post(url, json=body, headers=header)
# r = requests.get(url)
print(r.status_code, "-", r.json())
