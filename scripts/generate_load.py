import requests
import json
import time
import uuid

URL = "http://localhost:8000/approve"

REQUESTS = 1000

for i in range(REQUESTS):

    tx = {
        "Tx": {
            "TxID": str(uuid.uuid4()),
            "Buyer": "Alice",
            "Seller": "Bob",
            "Amount": 0.001,
            "Message": "Q_pay micropayment"
        }
    }

    try:
        requests.post(URL, json=tx)
    except:
        pass

    time.sleep(0.001)
