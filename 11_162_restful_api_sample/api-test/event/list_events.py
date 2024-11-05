import requests

url = "http://localhost:8080/events"

r = requests.get(url)
# r = requests.get(url)
print(r.status_code, "-", r.json())
