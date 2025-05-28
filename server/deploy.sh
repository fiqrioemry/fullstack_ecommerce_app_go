#!/bin/bash

echo "🚀 Deploying HappyShop App Server..."

docker compose down --remove-orphans
docker compose up -d --build

echo "✅ Deployment complete!"
