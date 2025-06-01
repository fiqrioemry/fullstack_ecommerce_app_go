#!/bin/bash

echo "🚀 Deploying HappyShop App Server..."
docker-compose -p ecommerce down -v

echo "🚀 Build container ...."
docker-compose -p ecommerce up -d --build

echo "✅ Deployment complete!"


