#!/bin/sh

# Create migration file
migrate create -ext sql -dir data/migration -seq $1