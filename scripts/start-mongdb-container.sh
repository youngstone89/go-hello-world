#!/bin/bash
docker run -d --name mongodb -v data:/data/db -e  MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.4.3


# mongoimport utility to load json file directly into collection
# mongoimport --username admin --password password --authenticationDatabase admin --db demo --collection recipes --file recipes.json --jsonArray

