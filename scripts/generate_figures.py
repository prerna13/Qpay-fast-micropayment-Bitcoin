import matplotlib.pyplot as plt
import pandas as pd

data = pd.read_csv("results/latency.csv")

plt.figure()

for committee_size in data["committee_size"].unique():

    subset = data[data["committee_size"] == committee_size]

    plt.plot(
        subset["requests"],
        subset["latency"],
        label=f"Committee {committee_size}"
    )

plt.yscale("log")

plt.xlabel("Requests per second")
plt.ylabel("Approval latency (seconds)")

plt.legend()

plt.title("Q_pay Approval Latency")

plt.savefig("figures/approval_latency.png")

print("Figure generated: figures/approval_latency.png")
