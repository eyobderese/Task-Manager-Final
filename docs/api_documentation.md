# API Documentation

# API Documentation for Go + Gin + MongoDB Application

---

## Table of Contents

1. [Introduction](#introduction)
2. [Getting Started](#getting-started)
3. [API Endpoints](#api-endpoints)

   - [Endpoint 1: Get](#get_task)
   - [Endpoint 2: GetById](#get_task_by_id)
   - [Endpoint 3: Create](#create_task)
   - [Endpoint 4: Update](#update_task)
   - [Endpoint 5: Delete](#delete_task)
   - [Endpoint 6: Register](#user_register)
   - [Endpoint 7: Login](#user_login)
   - [Endpoint 8: Promot](#user_promot)

4. [Data Models](#data-models)
5. [Error Handling](#error-handling)
6. [Autorization][#authorization]
7. [Testing][#testing]

---

## Introduction

This document provides an overview of the API built using Go, Gin, and MongoDB. Task Management API developed using Go, the Gin web framework, and MongoDB. It provides a set of endpoints that allow users to create, retrieve, update, and delete tasks. Each task is identified by a unique ID and includes details such as title, description, due date, and status. The API is designed to handle errors gracefully and provides meaningful error messages. It's a robust solution for managing tasks in a structured and efficient manner, offering a solid foundation for building task-oriented applications or services..

## Getting Started

To use this API, users need to have Go (latest version), MongoDB, and the Gin web framework installed. Users can get started by cloning the repository and installing any additional required libraries or tools.

### Prerequisites

- Go installed (version letest)
- MongoDB installed and running
- Gin web framework installed (`go get -u github.com/gin-gonic/gin`)
- Any additional libraries or tools required

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/eyobderese/A2SV-Backend-Learning-Path/tree/main/Tasks/Task5/task_manager_api
   cd task_manager_api
   ```
2. **Install the dependency**
   ```bash
   go mod tidy
   ```
3. **Run the Application**
   ```bash
   go run main.go
   ```

## Data Model

The main data model in this API is the Task model, which includes fields like ID, title, description, due date, and status.

The other data model is User, which include fields like ID, email, password

## Error Handling

The API provides meaningful error messages when something goes wrong, like when a user tries to retrieve a task that doesn't exist or when there's a problem connecting to the database.

## Endpoints

### <a id="get_task"></a>GET /task GET /tasks

Fetches all tasks.

Full URL: `http://localhost:8080/tasks`

#### Parameters

None

#### Request Headers

- `Authorization` (string, mandatory): Bearer token for authorization

  Example: `Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I`

#### Response

- `200 OK` on success

```json
{
  "data": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2024-08-06T15:31:14.5417452+03:00",
      "status": "Pending"
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "Second task",
      "due_date": "2024-08-07T15:31:14.5417452+03:00",
      "status": "In Progress"
    },
    {
      "id": "3",
      "title": "Task 3",
      "description": "Third task",
      "due_date": "2024-08-08T15:31:14.5427942+03:00",
      "status": "Completed"
    }
  ]
}
```

### <a id="get_task_by_id"></a> GET /tasks/{id}

Fetches a task by its ID.

Full URL: `http://localhost:8080/tasks/{id}`

#### Parameters

- `id` (integer): The ID of the task to fetch

#### Request Headers

- `Authorization` (string, mandatory): Bearer token for authorization

  Example: `Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I`

#### Response

- `200 OK` on success

#### Error Code

- `404 NOT FOUND` on fail

```json
{
  "data": {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-08-06T15:31:14.5417452+03:00",
    "status": "Pending"
  }
}
```

### <a id="create_task"></a> POST /tasks

Creates a new task.

Full URL: `http://localhost:8080/tasks`

#### Parameters

- `title` (string): The title of the task
- `description` (string): The description of the task
- `due_date` (string): The due date of the task
- `status` (string): The status of the task

#### Request Headers

- `Authorization` (string, mandatory): Bearer token for authorization

  Example: `Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I`

Request example:

```json
{
  "title": "New Task",
  "description": "New task description",
  "due_date": "2024-09-06T15:31:14.5417452+03:00",
  "status": "Pending"
}
```

#### Response

`200 OK` on sucess

```json
{
  "data": {
    "id": "4",
    "title": "New Task",
    "description": "New task description",
    "due_date": "2024-09-06T15:31:14.5417452+03:00",
    "status": "Pending"
  }
}
```

`404 Not Found` on error

```json
{
  "error": "Invalid data provided"
}
```

### <a id="update_task"></a> PUT /tasks/{id}

Updates a task by its ID.

Full URL: `http://localhost:8080/tasks/{id}`

#### Parameters

- `id` (integer): The ID of the task to update
- `title` (string, optional): The new title of the task
- `description` (string, optional): The new description of the task
- `due_date` (string, optional): The new due date of the task
- `status` (string, optional): The new status of the task

Request example:

```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2024-09-06T15:31:14.5417452+03:00",
  "status": "In Progress"
}
```

#### Request Headers

- `Authorization` (string, mandatory): Bearer token for authorization

  Example: `Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I`

#### Response

- `200 OK` on success

```json
{
  "data": {
    "id": "1",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-09-06T15:31:14.5417452+03:00",
    "status": "In Progress"
  }
}
```

`404 Not Found` on error

```json
{
  "error": "Task not found"
}
```

### <a id="delete_task"></a> DELETE /tasks/{id}

Deletes a task by its ID.

Full URL: `http://localhost:8080/tasks/{id}`

#### Parameters

- `id` (integer): The ID of the task to delete

#### Request Headers

- `Authorization` (string, mandatory): Bearer token for authorization

  Example: `Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I`

#### Response

- `204 No Content` on success

- `404 Not Found` if the task with the specified ID does not exist

```json
{
  "error": "Task not found"
}
```

### <a id="register"></a> POST /ragister

ragister new user

Full URL: `http://localhost:8080/ragister`

#### Parameters

- `email` (string, optional): The email of the user
- `password` (string, optional): The password of the user

Request example:

```json
{
  "email": "eyobderese@gmail.com",
  "password": "123456"
}
```

#### Response

- `201 Created` on success

````json
{
  "data": {
    "id": "1",
    "email": "eyobderese@gmail.com",
    "password": ""
  }
}


```json

````

`404 Not Found` on error

```json
{
  "error": "Internal server error"
}
```

### <a id="login"></a> POST /login

login user

Full URL: `http://localhost:8080/login`

#### Parameters

- `email` (string, optional): The email of the user
- `password` (string, optional): The password of the user

Request example:

```json
{
  "email": "eyobderese@gmail.com",
  "password": "1234545"
}
```

#### Response

- `200 SUCESS` on success

```json
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I"
}
```

`404 Not Found` on error
This error occurs when a requested resource is not found. It typically happens when a user tries to access a task that doesn't exist or when a specified endpoint is not available.

```json
{
  "error": "Invalid email or password"
}
```

### <a id="promot"></a> POST /promot

promot user

Full URL: `http://localhost:8080/promot/{id}`

#### Parameters

- `id` (string, mandatory): The id of the user
- `Authorization` (string, mandatory): The token of the user

Request example:

```json
{}
```

#### Request Headers

- `Authorization` (string, mandatory): Bearer token for authorization

  Example: `Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV5b2JkZXJlc2VAZ21haWwuY29tIiwiaWQiOiI2NmI0YjlkNjgxYzJhZjcwYjliYjAzOTIiLCJyb2xlIjoiYWRtaW4ifQ.lT87rkCYkAixZSStdDqrweuZbfmbjVZdvx0YzGsUo4I`

#### Response

- `200 SUCESS` on success

```json
{
  "data": {
    "id": "1",
    "email": "eyonderese@gmial.com",
    "role": "admin"
  }
}
```

`404 Not Found` on error

```json
{
  "error": "Internal server error"
}
```

# API Testing Guide

This guide provides instructions on how to run tests for our API.

## Prerequisites

- Go installed on your machine
- Access to the source code of the API

## Running the Tests

1. Open a terminal.

2. Navigate to the `test` directory in the API source code:

   ```bash
   cd path/to/api/test
   ```

   Replace `path/to/api/test` with the actual path to the `test` directory.

3. Run the tests:

   ```bash
   go test
   ```

   This command runs all tests in the `test` directory. If you want to run a specific test, you can use the `-run` flag followed by the name of the test:

   ```bash
   go test -run TestName
   ```

   Replace `TestName` with the name of the test you want to run.

## Understanding the Test Results

The `go test` command prints the results of the tests to the terminal. If a test passes, you'll see output like this:
PASS: TestName

In case of a failure, `go test` also prints more information about the failure, such as the line of code that caused the failure and a stack trace.
