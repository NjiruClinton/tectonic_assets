#!/bin/bash

source .env

goose -dir ./migrations postgres "$DB_CONNECTION_STRING" up

if [ $? -eq 0 ]; then
  echo "Migration completed successfully."
else
  echo "Migration failed."
  exit 1
fi