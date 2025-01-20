#!/bin/bash
set -e


if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose postgres $PROD_DB_URL up