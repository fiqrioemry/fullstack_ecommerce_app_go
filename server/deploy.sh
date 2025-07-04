#!/bin/bash

echo "start building container ...."
docker-compose -p ecommerce_app down -v
docker-compose -p ecommerce_app up -d --build

echo "Deployment complete!"
