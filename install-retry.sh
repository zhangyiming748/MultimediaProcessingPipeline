#!/bin/bash

RETRIES=15

for i in $(seq 1 $RETRIES); do
    apt-get install -y "$@" && break || {
        echo "Failed, retrying... ($i)"
        sleep 5
    }
done