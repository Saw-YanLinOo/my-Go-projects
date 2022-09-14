# my-Go-projects
Simple project for GoLand . This is CRUD operation on books table and  presistent data on Postgres. Please make sure you install postgres database on your machine.

## lib
    
    fiber - for web connection
    gorm - for database connection

## Testing the Endpoints

## Run go server
    make server or go run cmd/main.go

## POST: Add a new Book

    --request POST 
    --url http://localhost:3000/books 
    --header 'Content-Type: application/json' 
    --data '{
    "title": "Book A",
    "author": "Kevin Vogel",
    "description": "Some cool description"
    }'

## GET: Get All Books

    --request GET --url http://localhost:3000/books

## GET: Get Book by ID
    
    --request GET --url http://localhost:3000/books/1

## PUT: Update Book by ID

    --request PUT \
    --url http://localhost:3000/books/1 \
    --header 'Content-Type: application/json' \
    --data '{
    "title": "Updated Book Name",
    "author": "Kevin Vogel",
    "description": "Updated description"
    }'

## DELETE: Delete Book by ID

    --request DELETE --url http://localhost:3000/books/1
