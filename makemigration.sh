!#/bin/sh

# command: './makemigration.sh <tablename>'
migrate create -ext sql -dir internal/schema -seq create_"$1"_table