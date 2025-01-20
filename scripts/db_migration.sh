#!/bin/bash
set -e



# Debugging: Check if the variable is set
echo "PROD_DB_URL is set to: $PROD_DB_URL"

# Check if PROD_DB_URL is empty
if [ -z "$PROD_DB_URL" ]; then
    echo "Error: PROD_DB_URL is not set."
    exit 1
fi

# Proceed with migration
cd sql/schema
goose postgres $DB_URL up
