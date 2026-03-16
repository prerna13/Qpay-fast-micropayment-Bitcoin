# Q_pay Prototype Artifact

This repository contains a prototype implementation of the **Q_pay micropayment protocol** described in our paper.  
The artifact provides a runnable prototype, experimental scripts, and tools to reproduce the performance evaluation results.

Q_pay is a committee-based micropayment protocol designed to significantly reduce payment confirmation latency in Bitcoin by replacing network-wide transaction validation with lightweight committee approval.

---

## Repository Structure

qpay-artifact/

├── cmd/  
│   ├── buyer/            # Buyer client  
│   ├── seller/           # Seller service node  
│   ├── committee/        # Committee approval nodes  
│   └── aggregator/       # Commit proof aggregator  

├── pkg/  
│   ├── crypto/           # Schnorr / MuSig signature utilities  
│   ├── network/          # RPC communication layer  
│   ├── protocol/         # Q_pay approval protocol  
│   ├── bitcoin/          # Bitcoin testnet anchoring module  
│   └── types/            # Transaction data structures  

├── scripts/  
│   ├── start_committee.sh  
│   ├── run_experiment.sh  
│   ├── generate_load.py  
│   ├── measure_latency.py  
│   └── generate_figures.py  

├── config/  
│   └── committee.json  

├── results/  
│   └── latency.csv  

├── figures/  
│   └── approval_latency.png  

└── deploy/  
    └── aws_cluster.sh

---

## System Overview

The Q_pay protocol consists of four main components:

### Buyer
Initiates a micropayment transaction.

### Seller
Receives the transaction and requests approval from committee nodes.

### Committee Nodes
Validate the transaction and generate partial Schnorr signatures.

### Aggregator
Collects committee signatures and produces a commit proof.

The commit proof can optionally be anchored to the **Bitcoin testnet** using an OP_RETURN transaction.

---

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

To generate high-throughput micropayment traffic:

python scripts/generate_load.py

This simulates large numbers of payment requests to evaluate protocol scalability.

---

## Measuring Approval Latency

To reproduce the latency evaluation reported in the paper:

python scripts/measure_latency.py

This script generates experimental results:

results/latency.csv

---

## Generating Figures

Figures used in the evaluation can be reproduced automatically:

python scripts/generate_figures.py

This produces:

figures/approval_latency.png

---

## Bitcoin Testnet Integration (Optional)

To anchor commit proofs to Bitcoin testnet:

Start Bitcoin Core in testnet mode:

bitcoind -testnet -daemon

Verify the node is running:

bitcoin-cli -testnet getblockchaininfo

When enabled, Q_pay will anchor commit proofs to the testnet using OP_RETURN transactions.

---

## AWS Deployment

The repository includes scripts to deploy committee nodes on AWS.

Example:

deploy/aws_cluster.sh

This launches EC2 instances suitable for large-scale committee experiments.

---

## Reproducing Paper Experiments

The following scripts reproduce the main experiments used in the paper.

Committee latency evaluation:
python scripts/measure_latency.py

High-load stress testing:
python scripts/generate_load.py

Figure generation:
python scripts/generate_figures.py

---

## Notes

This implementation is a **research prototype** designed to demonstrate the Q_pay protocol and support reproducibility of experimental results.

It is **not intended for production deployment**.

---

## License

This project is released for academic research and reproducibility purposes.
