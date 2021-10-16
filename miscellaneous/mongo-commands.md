# MongoDB

```bash

docker-compose up -d mongo
docker-compose logs -f --tail 50 mongo

docker-compose exec mongo bash


echo 'db.runCommand("ping").ok' | mongo 127.0.0.1:27017/gogrpc --quiet

docker-compose exec mongo bash -c \
  "echo 'db.runCommand(\"ping\").ok' | mongo 127.0.0.1:27017/gogrpc --quiet"


mongo --host 127.0.0.1 --port 27017
mongo --host 127.0.0.1 --port 27017 -- gogrpc


use admin;
db.auth("root", "admin");

show dbs;

db.runCommand({connectionStatus: 1});
show roles;


use gogrpc;
show collections;

db.getCollectionNames();

db.item.find({});
db.item.find({"_id": ObjectId("615e73041061f9ca09c75f6e")}).pretty();


# docker container run --rm -it --name mongo mongo:5.0 bash

```
