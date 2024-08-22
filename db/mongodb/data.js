
// Use the database 'mydatabase'
use mydatabase;

// Create a collection 'users' and insert some data
db.users.insertMany([
    { "username": "user1", "email": "user1@example.com", "created_at": new Date() },
    { "username": "user2", "email": "user2@example.com", "created_at": new Date() },
    { "username": "user3", "email": "user3@example.com", "created_at": new Date() }
]);

// Create a collection 'products' and insert some data
db.products.insertMany([
    { "name": "Product1", "price": 100, "category": "Category1" },
    { "name": "Product2", "price": 200, "category": "Category2" },
    { "name": "Product3", "price": 300, "category": "Category3" }
]);
