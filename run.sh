#! /bin/bash
docker rm webcrawler-service
docker run -d \
  -p 8080:8080 \
  --name="webcrawler-service" \
  --env APP_PORT=8080 \
  --env GIN_MODE=release \
  webcrawler-service
