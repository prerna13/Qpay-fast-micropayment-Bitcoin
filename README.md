# Q_pay Prototype Artifact

This repository contains a prototype implementation of the **Q_pay micropayment protocol** described in our paper.  
The artifact provides a runnable prototype, experimental scripts, and tools to reproduce the performance evaluation results.

Q_pay is a committee-based micropayment protocol designed to significantly reduce payment confirmation latency in Bitcoin by replacing network-wide transaction validation with lightweight committee approval.

## Prerequisites

Install the following dependencies:

- Go ≥ 1.19
- Python ≥ 3.8
- Bitcoin Core (optional, for testnet anchoring)

Install Go dependencies:

go mod tidy

---

## Running the Prototype

Start the full experiment environment:

./scripts/run_experiment.sh

This will:
1. Start committee nodes  
2. Start the seller service  
3. Send a micropayment request from the buyer  
4. Generate a committee approval proof  

Example output:

Committee node running: 8000  
Seller running: 9000  
Buyer service latency: 0.42s  
Commit proof generated  

---

## Running Load Experiments

python scripts/generate_load.py

This simulates large numbers of payment requests to evaluate protocol scalability.

## Measuring Approval Latency

python scripts/measure_latency.py
The results get stored => results/latency.csv

---

## Generating Figures
python scripts/generate_figures.py


## Bitcoin Testnet Integration (Optional)

Start Bitcoin Core in testnet mode:

bitcoind -testnet -daemon

Verify => bitcoin-cli -testnet getblockchaininfo

When enabled, Q_pay will anchor commit proofs to the testnet using OP_RETURN transactions.
The repository includes scripts to deploy committee nodes on AWS.
deploy/aws_cluster.sh

Scripts reproduce the main experiments used in the paper.

Committee latency evaluation: python scripts/measure_latency.py

High-load stress testing: python scripts/generate_load.py

Figure generation: python scripts/generate_figures.py

## License

This project is released for academic research and reproducibility purposes.
