# API Operation Documentation

## Run the server

```sh
go run main.go
```

## Add a New Post

Use a `POST` request to add a new Post to the `/posts` endpoint. The request body should include the title, body, and author information of the Post.

```sh
curl -X POST -H "Content-Type: application/json" -d '{
  "Title": "First Post",
  "Body": "This is the body of the first post",
  "Author": {
    "FullName": "John Doe",
    "UserName": "johndoe",
    "Email": "john@example.com"
  }
}' http://localhost:8080/posts
```

Get All Posts
Use a GET request to fetch all Posts from the /posts endpoint.

```sh
curl -X GET http://localhost:8080/posts
```

Get a Specific Post
Use a GET request to fetch a Post with a specific ID from the /posts/{id} endpoint. The following example assumes the ID is 0.

```sh
curl -X GET http://localhost:8080/posts/0
```

Update a Specific Post
Use a PUT request to update a Post with a specific ID at the /posts/{id} endpoint. The request body should include the updated title, body, and author information of the Post. The following example assumes the ID is 0.

```sh
curl -X PUT -H "Content-Type: application/json" -d '{
"Title": "Updated Post",
"Body": "This is the updated body",
"Author": {
"FullName": "John Doe",
"UserName": "johndoe",
"Email": "john@example.com"
}
}' http://localhost:8080/posts/0
```

Partially Update a Specific Post
Use a PATCH request to partially update a Post with a specific ID at the /posts/{id} endpoint. The request body should only include the fields to be updated. The following example assumes the ID is 0.

```sh
curl -X PATCH -H "Content-Type: application/json" -d '{
"Body": "This is the partially updated body"
}' http://localhost:8080/posts/0
```

Delete a Specific Post
Use a DELETE request to delete a Post with a specific ID from the /posts/{id} endpoint. The following example assumes the ID is 0.

```sh
curl -X DELETE http://localhost:8080/posts/0
This document provides examples of how to use curl commands to test each API endpoint.
```
