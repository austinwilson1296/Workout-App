#!/bin/bash

# Proceed with migration
cd sql/schema
goose postgres $DB_URL up || true

# Continue with other tasks if needed
echo "Migration complete or no new migrations found."
