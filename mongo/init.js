db.createCollection('item');

db.item.insertMany([
  {
    "author_id": "1",
    "title": "Hello",
    "content": "World"
  }
]);
