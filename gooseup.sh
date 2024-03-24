#!/bin/sh
ENV_FILE="./.env"
if [ -f $ENV_FILE ]; then
  echo $ENV_FILE
  set -o allexport
  source $ENV_FILE
  set +o allexport
fi
goose -dir data/goose postgres "user=$SUPABASE_USER dbname=$SUPABASE_DB password=$SUPABASE_PASSWORD host=$SUPABASE_HOST port=$SUPABASE_PORT sslmode=disable" $@
#  This script will read the  .env  file and set the environment variables for the  goose  command. 
#  Now you can run the  goose  command with the  gooseup.sh  script. 
#  $ ./gooseup.sh status
