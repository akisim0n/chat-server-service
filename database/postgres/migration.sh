#!/bin/bash
source .env

export MIGRATION_DSN="host=chat_service_db port=$PG_PORT dbname=$PG_DB_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v