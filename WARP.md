# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Development Commands

### Building and Testing
- **Run all tests**: `go test ./...`
- **Run tests with verbose output**: `go test -v ./...`
- **Run tests with coverage**: `go test -coverprofile=coverage.out ./...` then `go tool cover -html=coverage.out`
- **Build project** (runs tests after linting): `make build`

### Linting and Code Quality
- **Lint with errors only**: `make lint-errors` - runs vacuum on OpenAPI spec and golangci-lint
- **Full lint with strict mode**: `make lint` - runs vacuum in hard mode and golangci-lint
- **OpenAPI spec linting**: `vacuum lint -d todo-openapi.yaml --no-banner --no-clip --ignore-file vacuum-ignore-file.yaml`
- **Go code linting**: `golangci-lint run`

### Running the Application
- **Start the nethttp server**: `make nethttp` (runs on http://localhost:8080/)
- **Direct run**: `go run nethttp/main.go`

### Testing Individual Components
- **Run model tests**: `go test ./shared/model`
- **Run repo tests**: `go test ./shared/repo`
- **Run service tests**: `go test ./shared/service`

## Architecture Overview

This is a Go web frameworks comparison project that implements a TODO application with multiple HTTP framework backends. The project follows a clean architecture pattern with clear separation of concerns.

### Core Architecture Layers

1. **Models** (`shared/model/`): Domain entities and data structures
   - User, Project, List, Todo entities with Create/Update variants
   - Built-in validation using `validator/v10`
   - Fake data generation using `gofakeit/v7`
   - ULID-based IDs using `oklog/ulid/v2`

2. **Repository Layer** (`shared/repo/`): Data access abstraction
   - In-memory implementations for all entities
   - Thread-safe operations with RWMutex
   - Generic CRUD interfaces

3. **Service Layer** (`shared/service/`): Business logic
   - Password hashing using Argon2id
   - JWT token generation and validation
   - Validation and business rule enforcement

4. **HTTP Layer** (`nethttp/`): Web framework implementations
   - Standard library `net/http` implementation
   - RESTful API with JSON responses
   - Comprehensive error handling

### Key Dependencies and Patterns

- **Authentication**: JWT tokens with Argon2id password hashing
- **Validation**: Struct-based validation with tags (`go-playground/validator/v10`)
- **IDs**: ULIDs for all entity identifiers (`oklog/ulid/v2`)
- **Testing**: Comprehensive test coverage with subtests and table-driven tests
- **Fake Data**: `gofakeit/v7` for generating test data

### API Structure

The application exposes a REST API with the following endpoints:
- `/projects` - CRUD operations for projects
- `/lists` - CRUD operations for lists 
- `/lists/{listId}/todos` - CRUD operations for todos within lists
- `/users` - User management including password updates
- `/logins` - Authentication endpoint
- `/health` - Health check endpoint

### Framework Comparison Design

The shared layer (`shared/`) contains the core business logic that can be reused across different HTTP framework implementations. Currently implements:
- `nethttp/`: Standard library implementation

The `go-don` framework dependency in `go.mod` suggests planned implementation using the Don framework, though not yet present in the codebase.

## Development Notes

### Testing Philosophy
- All models have comprehensive validation tests
- Repository layer includes concurrency testing  
- Service layer tests include error scenarios and edge cases
- Mock repositories used for service testing isolation

### Security Considerations
- Passwords are hashed using Argon2id with recommended parameters
- JWT tokens have configurable expiry (currently 1 hour)
- User authentication required for protected endpoints
- Time-constant password comparison to prevent timing attacks

### Code Generation and Fakes
All models include `Random*()` methods for generating test data. Use these for consistent test data generation rather than manual construction.

### OpenAPI Specification
The project includes a comprehensive OpenAPI 3.0.3 specification (`todo-openapi.yaml`) that documents all API endpoints. Use vacuum for OpenAPI linting with custom ignore rules in `vacuum-ignore-file.yaml`.

### Makefile Targets
The default target is `build`, which runs linting and tests. Use individual targets for specific operations:
- `make lint-errors` for development (shows errors only)
- `make lint` for CI/CD (strict mode with failure on errors)
- `make cover` for coverage reporting
