```markdown
# Noona HQ Golang App Template

Template repository for building Noona applications using Golang.

## Tech Stack
- **Languages:** Go 1.25.5
- **Frameworks:** Echo v4, Zap
- **Key Dependencies:** MongoDB driver, JWT, envconfig, TOML, flatbson, uniuri

## Architecture / How it Works
- **Configuration Management:** Uses `envconfig` for processing configuration from environment variables.
- **Logging:** Integrated with `zap` for structured logging.
- **Server Initialization:** Sets up the Echo server with configurations and middleware.
- **Dependency Injection:** Injects configurations and logger into the server instance.
- **Service Layer:** Handles business logic and interacts with the data store.

## Key Interfaces / API
- **RESTful API:** Built with Echo, exposing endpoints for various services.
- **Middleware:** Utilizes Echo middleware for request processing, logging, and error handling.

## Dependencies
- **Noona SDK:** Integrates with `noona-sdk-go` for interacting with other Noona services.
- **MongoDB:** Uses `go.mongodb.org/mongo-driver` for database interactions.
- **JWT:** Implements authentication and authorization with `github.com/golang-jwt/jwt`.
```