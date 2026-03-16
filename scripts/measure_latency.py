import subprocess
import statistics

runs = []

for i in range(20):

    out = subprocess.check_output(
        ["go","run","cmd/buyer/main.go"]
    ).decode()

    latency = float(out.split(":")[1].replace("ms",""))

    runs.append(latency)

print("Average latency:", statistics.mean(runs))
