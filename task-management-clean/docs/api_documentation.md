### Task Management Endpoints

### API Documentation

#### Authentication and User Management

##### Register a New User
**Endpoint:** `POST /register`

**Description:** Register a new user.

**Request Body:**
```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

**Response:**
- **200 OK**
```json
{
  "id": "string",
  "name": "string",
  "email": "string"
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### User Login
**Endpoint:** `POST /login`

**Description:** Authenticate a user and return a JWT token.

**Request Body:**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response:**
- **200 OK**
```json
{
  "message": "successful login",
  "token": "string",
  "user": {
    "id": "string",
    "name": "string",
    "email": "string"
  }
}
```
- **500 Internal Server Error**
```json
{
  "error": "string"
}
```

##### Get All Users
**Endpoint:** `GET /users/`

**Description:** Retrieve a list of all users.

**Response:**
- **200 OK**
```json
{
  "users": [
    {
      "id": "string",
      "name": "string",
      "email": "string"
    }
  ]
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### Get User by ID
**Endpoint:** `GET /users/:id`

**Description:** Retrieve a user by ID.

**Response:**
- **200 OK**
```json
{
  "id": "string",
  "name": "string",
  "email": "string"
}
```
- **404 Not Found**
```json
{
  "error": "string"
}
```

##### Promote User to Admin
**Endpoint:** `PUT /users/promote`

**Description:** Promote a user to admin.

**Request Body:**
```json
{
  "id": "string",
  "email": "string"
}
```

**Response:**
- **200 OK**
```json
{
  "id": "string",
  "name": "string",
  "email": "string",
  "is_admin": true
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### Update User
**Endpoint:** `PUT /users/:id`

**Description:** Update user details.

**Request Body:**
```json
{
  "name": "string",
  "email": "string"
}
```

**Response:**
- **200 OK**
```json
{
  "id": "string",
  "name": "string",
  "email": "string"
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### Delete User
**Endpoint:** `DELETE /users/:id`

**Description:** Delete a user by ID.

**Response:**
- **200 OK**
```json
{
  "message": "deleted"
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

#### Task Management

##### Create a New Task
**Endpoint:** `POST /tasks/`

**Description:** Create a new task.
**> [!NOTE]
> Here it is not required to set the due_date and the title as a default value the due_date is the current date and the description will be empty 
**Request Body:**
```json
{
  "user_id": "string",
  "title": "string",
  "description": "string",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "string"
}
```

**Response:**
- **201 Created**
```json
{
  "id": "string",
  "user_id": "string",
  "title": "string",
  "description": "string",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "string"
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### Get All Tasks
**Endpoint:** `GET /tasks/`

**Description:** Retrieve a list of all tasks.

**Response:**
- **200 OK**
```json
[
    {
      "id": "string",
      "user_id": "string",
      "title": "string",
      "description": "string",
      "due_date": "2024-01-01T00:00:00Z",
      "status": "string"
    }
]

```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### Get Task by ID
**Endpoint:** `GET /tasks/:id`

**Description:** Retrieve a task by ID.

**Response:**
- **200 OK**
```json
{
  "id": "string",
  "user_id": "string",
  "title": "string",
  "description": "string",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "string"
}
```
- **404 Not Found**
```json
{
  "error": "string"
}
```

##### Update Task
**Endpoint:** `PUT /tasks/:id`

**Description:** Update a task by ID.

**Request Body:**
```json
{
  "title": "string",
  "description": "string",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "string"
}
```

**Response:**
- **200 OK**
```json
{
  "id": "string",
  "user_id": "string",
  "title": "string",
  "description": "string",
  "due_date": "2024-01-01T00:00:00Z",
  "status": "string"
}
```
- **400 Bad Request**
```json
{
  "error": "string"
}
```

##### Delete Task
**Endpoint:** `DELETE /tasks/:id`

**Description:** Delete a task by ID.

**Response:**
- **204 No Content**
- **400 Bad Request**
```json
{
  "error": "string"
}
```

