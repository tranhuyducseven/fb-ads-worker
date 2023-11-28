// /docker-entrypoint-initdb.d/init-db.js

db = db.getSiblingDB('example-db');

// Create collections, insert initial data, etc.
// For example:
db.createCollection('example_collection');
db.example_collection.insert({ key: 'value' });
