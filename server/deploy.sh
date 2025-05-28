#!/bin/bash

echo "🚀 Deploying HappyShop App Server..."
docker-compose down

echo "🚀 Build container ...."
docker-compose up -d --build

echo "✅ Deployment complete!"
