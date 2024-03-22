#!/bin/sh

# This script is used to migrate the database to the given version
# migrate -path data/migration -database "sqlite://data/newser.sqlite" -verbose up $1
migrate -path data/migration -database "libsql://newser-nemo-omen.turso.io?authToken=$NEWSER_TURSO_TOKEN" -verbose up $1
