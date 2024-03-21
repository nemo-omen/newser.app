#!/bin/sh

# This script is used to migrate the database to the given version
migrate -path data/migration -database "sqlite://data/newser.sqlite" -verbose up $1
