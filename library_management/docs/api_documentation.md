# Library Management API Documentation

## Overview

The Library Management API provides endpoints to manage books and members in a library system. It allows adding new books and members, borrowing and returning books, and retrieving information about books and members.

## Base URL

`http://localhost:9090`

## Endpoints

### Add a New Book

- **URL:** `/books/`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "id": 1,
    "author": "Author Name",
    "title": "Book Title",
    "status": "Available"
  }
  ```
- **Response:**
  - **Status:** `201 Created`
  - **Body:** The created book object.

### Add a New Member

- **URL:** `/members/`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "id": 1,
    "name": "Member Name",
    "borrowed_book": []
  }
  ```
- **Response:**
  - **Status:** `201 Created`
  - **Body:** The created member object.

### Get All Books

- **URL:** `/books/`
- **Method:** `GET`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** An array of all book objects.
  ```json
  [
    {
      "id": 1,
      "author": "Author Name",
      "title": "Book Title",
      "status": "Available"
    },
    ...
  ]
  ```

### Get Book by ID

- **URL:** `/books/{id}`
- **Method:** `GET`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** The book object with the specified ID.
  ```json
  {
    "id": 1,
    "author": "Author Name",
    "title": "Book Title",
    "status": "Available"
  }
  ```

### Get All Members

- **URL:** `/members/`
- **Method:** `GET`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** An array of all member objects.
  ```json
  [
    {
      "id": 1,
      "name": "Member Name",
      "borrowed_book": []
    },
    ...
  ]
  ```

### Get Member by ID

- **URL:** `/members/{id}`
- **Method:** `GET`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** The member object with the specified ID.
  ```json
  {
    "id": 1,
    "name": "Member Name",
    "borrowed_book": []
  }
  ```

### Borrow a Book

- **URL:** `/members/{memberID}/borrow/{bookID}`
- **Method:** `POST`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** The updated member object with the borrowed book.
  ```json
  {
    "id": 1,
    "name": "Member Name",
    "borrowed_book": [
      {
        "id": 1,
        "author": "Author Name",
        "title": "Book Title",
        "status": "Borrowed"
      }
    ]
  }
  ```

### Return a Book

- **URL:** `/members/{memberID}/return/{bookID}`
- **Method:** `PUT`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** The updated member object without the returned book.
  ```json
  {
    "id": 1,
    "name": "Member Name",
    "borrowed_book": []
  }
  ```

### Get Borrowed Books by Member

- **URL:** `/members/{memberID}/books/`
- **Method:** `GET`
- **Response:**
  - **Status:** `200 OK`
  - **Body:** An array of books borrowed by the specified member.
  ```json
  [
    {
      "id": 1,
      "author": "Author Name",
      "title": "Book Title",
      "status": "Borrowed"
    },
    ...
  ]
  ```

## Models

### Book

- **Fields:**
  - `id` (int): The unique identifier of the book.
  - `author` (string): The author of the book.
  - `title` (string): The title of the book.
  - `status` (string): The status of the book (e.g., "Available", "Borrowed").

### Member

- **Fields:**
  - `id` (int): The unique identifier of the member.
  - `name` (string): The name of the member.
  - `borrowed_book` (array of Book): The list of books borrowed by the member.
