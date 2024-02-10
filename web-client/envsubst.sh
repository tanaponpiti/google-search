#!/bin/sh
API_URL=${API_URL:-""}

# Use sed to find the line containing 'window.apiBaseUrl =' and replace the URL for using env variable in built dist package
sed -i "s|window\.apiBaseUrl = \".*\";|window.apiBaseUrl = \"$API_URL\";|g" /usr/share/nginx/html/config.js
