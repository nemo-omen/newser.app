#!/bin/sh

# This script is used to force migration down the database to a specific version
# Usage: ./forcemigration.sh <version>
migrate -path data/migration -database "sqlite://data/newser.sqlite" -verbose force $1
