#!/usr/bin/env bash
set -e

echo "🚀 Starting Go API (MongoDB Atlas)"

# Load env vars
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo "❌ .env file missing"
  exit 1
fi

# Validate required vars
REQUIRED_VARS=(MONGO_URI APP_PORT)
for v in "${REQUIRED_VARS[@]}"; do
  if [ -z "${!v}" ]; then
    echo "❌ Missing env var: $v"
    exit 1
  fi
done

# Basic URI sanity check
if [[ "$MONGO_URI" != mongodb* ]]; then
  echo "❌ Invalid MONGO_URI"
  exit 1
fi

echo "✅ Env loaded"
echo "🔌 Connecting to MongoDB Atlas..."
echo "🌍 Starting server on port $APP_PORT"

go run ./cmd/server
