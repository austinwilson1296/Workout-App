#!/bin/bash

set -e

cd ../sql/schema

goose postgres postgresql://doadmin:AVNS_9SkPpEWKkedAQAAzoTa@db-postgresql-nyc3-49886-do-user-18694038-0.g.db.ondigitalocean.com:25060/defaultdb?sslmode=require up