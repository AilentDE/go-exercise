import requests

url = "http://localhost:8080/events"
event_id = 1

r = requests.get(url + "/" + str(event_id))
print(r.status_code, "-", r.json())
