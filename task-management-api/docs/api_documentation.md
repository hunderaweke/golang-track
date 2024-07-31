# Task Management API Documentation

## Base URL
```
http://localhost:7070
```

## Endpoints

### Get All Tasks
**Endpoint**: `/tasks/`  
**Method**: `GET`  
**Description**: Retrieves a list of all tasks.

**Request**:
```http
GET /tasks/
```

**Response**:
```json
{
  "tasks": {
    "1": {
      "id": "1",
      "title": "Complete Go project",
      "description": "Finish the distributed system project in Go",
      "due_date": "2023-08-07T12:34:56Z",
      "status": "pending"
    },
    "2": {
      "id": "2",
      "title": "Write blog post",
      "description": "Write a blog post about the Go project",
      "due_date": "2023-08-14T12:34:56Z",
      "status": "pending"
    },
    "3": {
      "id": "3",
      "title": "Update resume",
      "description": "Add the new project details to the resume",
      "due_date": "2023-08-03T12:34:56Z",
      "status": "in progress"
    }
  }
}
```

### Get Task by ID
**Endpoint**: `/tasks/{id}`  
**Method**: `GET`  
**Description**: Retrieves a task by its ID.

**Request**:
```http
GET /tasks/{id}
```

**Path Parameters**:
- `id` (string): The ID of the task to retrieve.

**Response**:
```json
{
  "id": "1",
  "title": "Complete Go project",
  "description": "Finish the distributed system project in Go",
  "due_date": "2023-08-07T12:34:56Z",
  "status": "pending"
}
```

### Create Task
**Endpoint**: `/tasks/`  
**Method**: `POST`  
**Description**: Creates a new task. The `ID` and `due_date` are optional. If not provided, `ID` will be auto-generated and `due_date` will default to the current date and time.

**Request**:
```http
POST /tasks/
Content-Type: application/json

{
  "title": "New Task",
  "description": "Description for the new task",
  "status": "pending"
}
```

**Response**:
```json
{
  "id": "4",
  "title": "New Task",
  "description": "Description for the new task",
  "due_date": "2023-08-20T12:34:56Z",
  "status": "pending"
}
```

### Update Task
**Endpoint**: `/tasks/{id}`  
**Method**: `PUT`  
**Description**: Updates an existing task by its ID.

**Request**:
```http
PUT /tasks/{id}
Content-Type: application/json

{
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2023-08-21T12:34:56Z",
  "status": "completed"
}
```

**Path Parameters**:
- `id` (string): The ID of the task to update.

**Response**:
```json
{
  "id": "1",
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2023-08-21T12:34:56Z",
  "status": "completed"
}
```

### Delete Task
**Endpoint**: `/tasks/{id}`  
**Method**: `DELETE`  
**Description**: Deletes a task by its ID.

**Request**:
```http
DELETE /tasks/{id}
```

**Path Parameters**:
- `id` (string): The ID of the task to delete.

**Response**:
```json
{
  "message": "deleted"
}
```

## Models

### Task
```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "string (date-time)",
  "status": "string"
}
```

## Error Handling
### Common Errors
- **404 Not Found**: Task not found
  ```json
  {
    "error": "task not found"
  }
  ```
- **400 Bad Request**: Invalid request data
  ```json
  {
    "error": "Invalid request data"
  }
  ```
