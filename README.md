# gRPC Blog API with MongoDB

This is a simple gRPC API to expose operations with Blogs, made with Golang and persisting the data on MongoDB.

## Up and Running

```bash

docker-compose up -d mongo
docker-compose logs -f --tail 50 mongo

docker-compose up -d --build blog
docker-compose logs -f --tail 50 blog

```
