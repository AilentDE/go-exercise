import requests

url = "http://localhost:8080/signup"

body = {
    "email": "test02@gmail.com",
    "password": "1234abcd"
}
r = requests.post(url, json=body)
# r = requests.get(url)
print(r.status_code, "-", r.json())
