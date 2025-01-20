#!/bin/bash

# Proceed with migration
cd sql/schema
goose postgres $DB_URL up
