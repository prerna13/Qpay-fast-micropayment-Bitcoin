import requests
import uuid
import time
import csv

sizes = [16,32,64,128,256,512]

with open("results/latency.csv","w") as f:

    writer = csv.writer(f)

    writer.writerow(["committee_size","requests","latency"])

    for size in sizes:

        for rps in [500,1000,2000,3500]:

            start = time.time()

            for i in range(rps):

                tx = {
                    "TxID":str(uuid.uuid4()),
                    "Buyer":"Alice",
                    "Seller":"Bob",
                    "Amount":0.001
                }

                requests.post("http://localhost:9000/request",json=tx)

            latency = time.time() - start

            writer.writerow([size,rps,latency])
