
# Q_pay Prototype for fast micropayment using bitcoin 

This repository contains a prototype implementation of **Q_pay**, a fast micropayment protocol for Bitcoin that enables sub-second service confirmation using committee-based approval and Schnorr multi-signatures.

## Features

• Committee-based micropayment approval  
• Schnorr multisignature commit proof generation  
• Reduced-difficulty Q_block settlement  
• Parallel approval committees  
• Load-testing scripts for experimental evaluation  

## Repository Structure

cmd/  
  buyer/ – Buyer client  
  seller/ – Seller service node  
  committee_node/ – Approval committee member  
  aggregator/ – Signature aggregation node  

pkg/  
  crypto/ – Schnorr signature implementation  
  protocol/ – Q_pay protocol logic  
  committee/ – Committee operations  
  types/ – Transaction structures  

scripts/  
  run_experiment.sh – Reproduce experiments  
  generate_load.py – Generate approval requests  

config/  
  committee.json – Committee configuration  

## Requirements

Go 1.19+

Python 3.8+

## Running the Prototype
go run cmd/committee_node/main.go
go run cmd/seller/main.go
python scripts/generate_load.py

## Experimental Setup

The experiments in the paper were conducted using:

• 600 nodes across AWS regions  
• Committee sizes from 16–512  
• Request rates up to 3500 approvals/sec  

## Reproducing Results

1. Deploy committee nodes
2. Run load generator
3. Measure approval latency

### Start Committee Nodes
