# Project Structure

This document outlines the standard project structure and coding conventions to be followed. These guidelines ensure consistency, maintainability, and scalability across the project.

## Key Principles

1. **Separation of Concerns**: Each layer has a specific responsibility
2. **Single Responsibility**: Files and methods should have a single, well-defined purpose
3. **Controller-Usecase Mapping**: Each controller method should call exactly one use case method
4. **Modularity**: Modules should be as independent as possible
5. **Clear Interfaces**: Use well-defined interfaces between layers

## Directory Structure

```
├── cmd/                    # Main entry points to executable commands
├── config/                 # Configuration files for different environments
└── internal/               # Private application code
    ├── app/                # Application orchestration layer
    └── modules/            # Feature modules
        └── <module_name>/  # e.g., user, order, finance, etc.
    └── pkg/                # Project-specific packages
```

## Detailed Structure Rules

### 1. Command Layer (`cmd/`)

- Use a single folder tree structure with separate files for each command (e.g., `http.go`, `grpc.go`, `cli.go`)
- Keep command implementations minimal - they should only call methods from the app layer
- No build logic or graceful shutdown handling in command files - this belongs in the app layer
- Use Cobra for managing commands, enabling multiple commands in one binary

### 2. Configuration (`config/`)

- Store YAML configuration files for each environment
- Include `base.go` and module-specific configuration files (e.g., `user.go`, `order.go`)

### 3. Application Layer (`internal/app/`)

- Responsible for orchestrating the application
- Build all modules and handle dependency injection
- Implement build logic and graceful shutdown

### 4. Modules (`internal/modules/<module_name>/`)

Each module should follow this structure:

- **module.go**: Contains module build logic

- **controller/**: API controllers separated by user type

  - `userctrl/` - User-facing endpoints
  - `adminctrl/` - Admin-facing endpoints
  - `b2bctrl/` - B2B endpoints
  - `clictrl/` - CLI commands
  - `consumerctrl/` - Consumer handlers
  - Never call a single use case method from multiple controllers

- **usecase/**: Business logic organized by domain

  - Maintain one-to-one mapping between controller methods and use case methods
  - One public method per file (e.g., `create.go`, `update.go`, `admin_update.go`)
  - Extract shared logic into private methods within the use case struct
  - For shared logic between use cases, create service commands/queries
  - Implement separate use case methods for each controller
  - Each file should contain only one public method (with related private methods)

- **service/**: Optional shared business logic layer

  - `bus.go`: Container struct for all commands and queries
  - Implement each command/query in a separate file with clearly defined inputs/outputs
  - Share logic between different use cases via service commands/queries

- **public/**: Public interfaces for sharing logic with other modules

  - Contains commands and queries bus
  - Must not call components from other modules
  - Keep public interfaces clean and independent of other modules

- **repository/**: Data access implementations

  - Database repositories (e.g., `pguser/`, `pgrole/`)
  - Cache implementations
  - External API clients
  - Unit of work implementation (`pguow/`)

- **domain/**: Core domain models and interfaces
  - `entity.go`: Domain entities
  - `repo.go`: Repository interfaces
  - `conv.go`: Entity/DTO conversion methods
  - `dto.go`: Data transfer objects
  - `codes.go`: Domain error codes
  - Subdirectories for related domain concepts (e.g., `uow/`, `user/`, `role/`)
  - Organize domain concepts in their own packages

### 5. Project Packages (`internal/pkg/`)

- Clean packages specific to the current project
- General-purpose packages should be placed in a common package repository
