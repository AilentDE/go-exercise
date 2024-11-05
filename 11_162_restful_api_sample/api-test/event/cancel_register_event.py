import requests

url = "http://localhost:8080/events"
event_id = 1

header = {
    "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QwMkBnbWFpbC5jb20iLCJleHAiOjE3MzA4MzM0MTcsInVzZXJJZCI6Mn0.Cw8vzPellDQVkXYh0J7epEg6z0Nysx0tJ1qKXiSMDxM"
}

body = {}

r = requests.delete(f"{url}/{event_id}/register", json=body, headers=header)
# r = requests.get(url)
print(r.status_code, "-", r.json())
