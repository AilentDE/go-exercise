import requests

url = "http://localhost:8080/events"
event_id = 5

header = {
    "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QwMkBnbWFpbC5jb20iLCJleHAiOjE3MzA4Mjk1OTUsInVzZXJJZCI6Mn0.7XBw3zWc2hyNID6eF678jdMOH8FSoSTu7wePpbQV9rw"
}

r = requests.delete(f"{url}/{event_id}", headers=header)
# r = requests.get(url)
print(r.status_code, "-", r.json())
