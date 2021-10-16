#!/bin/bash

mongo --host 127.0.0.1 --port 27017 -- "$MONGO_INITDB_DATABASE" <<EOF
  var rootUser = '$MONGO_INITDB_ROOT_USERNAME';
  var rootPassword = '$MONGO_INITDB_ROOT_PASSWORD';
  var admin = db.getSiblingDB('admin');
  admin.auth(rootUser, rootPassword);

  load("/init.js")
EOF
