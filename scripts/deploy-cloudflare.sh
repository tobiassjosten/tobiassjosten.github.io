#!/bin/bash
set -e

echo "=== Cloudflare Pages Deployment ==="
echo

# Validate required environment variables
if [ -z "$CLOUDFLARE_ACCOUNT_ID" ]; then
    echo "Error: CLOUDFLARE_ACCOUNT_ID environment variable not set"
    echo "Get your Account ID from: https://dash.cloudflare.com"
    exit 1
fi

if [ -z "$CLOUDFLARE_API_TOKEN" ]; then
    echo "Error: CLOUDFLARE_API_TOKEN environment variable not set"
    echo "Create an API token with 'Cloudflare Pages:Edit' permissions at:"
    echo "https://dash.cloudflare.com/profile/api-tokens"
    exit 1
fi

# Check that public/ directory exists
if [ ! -d "public" ]; then
    echo "Error: public/ directory not found"
    echo "Build the site first with: hugo"
    exit 1
fi

# Get project name from environment or use default
PROJECT_NAME="${PROJECT_NAME:-tobiassjosten-net}"

# Get branch from environment or use default
BRANCH="${BRANCH:-main}"

echo "Deploying to Cloudflare Pages..."
echo "Project: $PROJECT_NAME"
echo "Branch: $BRANCH"
echo

# Deploy using wrangler via npx
npx wrangler pages deploy public \
    --project-name="$PROJECT_NAME" \
    --branch="$BRANCH"

echo
echo "Deployment complete!"
