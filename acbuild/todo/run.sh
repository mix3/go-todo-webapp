#!/usr/bin/env bash

set -e

if [ "$DELAY" != "" ]; then
  sleep $DELAY
fi

if [ "$DB_HOST" == "" ]; then
  DB_HOST="127.0.0.1"
fi

ARGS=""
if [ "$DB_USER" != "" ]; then
  ARGS="$ARGS -u $DB_USER"
fi
if [ "$DB_PASS" != "" ]; then
  ARGS="$ARGS -p $DB_PASS"
fi
if [ "$DB_PORT" != "" ]; then
  ARGS="$ARGS -P $DB_PORT"
fi
if [ "$DB_HOST" != "" ]; then
  ARGS="$ARGS -h $DB_HOST"
fi

echo "CREATE DATABASE IF NOT EXISTS todo" | mysql $ARGS

/bin/todo
