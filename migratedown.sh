#!/bin/sh

# This script is used to migrate down the database to a specific version
# Usage: ./migratedown.sh <version>
migrate -path data/migration -database "sqlite://data/newser.sqlite" -verbose down $1
