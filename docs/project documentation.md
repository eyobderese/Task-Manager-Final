# Project Documentation: Task Manager

## Overview

The Task Manager project is designed following the Clean Architecture principles. This architecture ensures a clear separation of concerns, making the codebase more maintainable, testable, and scalable. The project is structured into distinct layers, each with specific responsibilities.

## Folder Structure

```plaintext
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go
    └── user_usecases.go
```

## Layers Description

1. **Delivery**
   - Purpose: Handles HTTP requests and responses, providing the interface through which the user interacts with the application.
     - main.go: The entry point of the application, setting up the server and initializing routes.
     - controllers/: Contains the controller logic, which processes incoming requests and interacts with use cases.
     - controller.go: Implements request handling and response formatting.
     - routers/: Manages route definitions and middleware configuration.
     - router.go: Defines routes and associates them with the appropriate controllers.
2. **Domain**
   - Purpose: Contains the core business logic and domain models. This layer is independent of external frameworks and technologies.
     - domain.go: Defines the core entities and business rules of the application. This is where domain-specific logic resides.
3. **Infrastructure**
   - Purpose: Provides implementation details for external concerns such as authentication and communication with external services.
     - auth_middleWare.go: Implements middleware for authentication and authorization.
     - jwt_service.go: Handles JSON Web Token (JWT) operations such as generation and validation.
     - password_service.go: Provides functionality for password hashing and validation.
4. **Repositories**
   - Purpose: Contains the data access layer, interacting with databases or other storage systems.
     - task_repository.go: Implements methods for accessing and manipulating task data.
     - user_repository.go: Implements methods for accessing and manipulating user data.
5. **Usecases**
   - Purpose: Contains application-specific business logic, coordinating between the domain and repositories.
     - task_usecases.go: Defines use cases related to task management, such as creating, updating, and retrieving tasks.
     - user_usecases.go: Defines use cases related to user management, such as user registration and login.

## Getting Started

Clone the Repository:

```bash
git clone https://github.com/your-repo/task-manager.git
```

Navigate to the Project Directory:

```bash
cd task-manager
```

Install Dependencies: Ensure you have Go installed, and run:

```bash
go mod tidy
```

Run the Application:

```bash
go run Delivery/main.go
```

Run Tests: To ensure everything is working as expected:

```bash
go test ./...
```

## Testing

Unit tests are located alongside the respective components. To add new tests, create a file with a \_test.go suffix in the same directory as the component you are testing. Run tests using:

## Code Style and Conventions

- Follow Go’s standard formatting rules. Use `gofmt` to format code.
- Write clear and concise comments for exported functions and types.
- Keep business logic in the Domain and Usecases layers, and external concerns in the Infrastructure layer.
