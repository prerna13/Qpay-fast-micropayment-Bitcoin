#!/bin/bash

REGION="us-east-1"

for i in {1..20}
do
  aws ec2 run-instances \
  --image-id ami-0c02fb55956c7d316 \
  --instance-type t3.small \
  --count 1 \
  --region $REGION
done
