#!/bin/sh
set -e

echo "Waiting for database to be ready..."

until nc -z -v -w30 db 3306
do
  echo "Waiting for database connection..."
  sleep 5
done

echo "Database is ready, starting the app..."
/root/main
