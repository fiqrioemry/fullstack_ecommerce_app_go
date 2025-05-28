#!/bin/bash

echo "🚀 Deploying HappyShop App Server..."
docker-compose down -v

echo "🚀 Build container ...."
docker-compose up -d --build

echo "✅ Deployment complete!"
