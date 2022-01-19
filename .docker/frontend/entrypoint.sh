#!/usr/bin/env sh

echo "API_HOST=\"$API_HOST\"" > /app/.env

nginx -g 'daemon off;'
