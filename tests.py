import requests
import json

import os

# test for forecast

resp = requests.get("http://127.0.0.1:8080/forecast/1")

print(resp.content)

# test for delete webhook

resp = requests.delete("http://127.0.0.1:8080/webhooks/5")

print(resp.content)
