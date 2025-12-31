#!/bin/bash
set -e

# Load environment variables from .env.test if present
ENV_FILE="../env/.env.test"
if [ -f "$ENV_FILE" ]; then
    export $(grep -v '^#' "$ENV_FILE" | xargs)
fi

# Fallback defaults if not set in env file
DB_URL=${DB_URL:-"postgres://postgres:postgres@localhost:5432/templatedb"}

echo "Initializing Test Database at $DB_URL..."

# Apply init.sql
# We detect if running in devcontainer where psql might need host adjustment (already handled by DB_URL usually)
# If psql is not installed locally, this might fail, so we assume it is (installed in previous step)

psql "$DB_URL" -f tests/db/init.sql

echo "Test Database Initialized Successfully."
